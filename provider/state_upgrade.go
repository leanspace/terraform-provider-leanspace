package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// UnwrapSingleNestedStateUpgrader returns a StateUpgrader that migrates SDK v2 state
// to Framework state by unwrapping any "list-of-1" arrays ([{...}]) that were the
// SDK v2 workaround for single nested objects.  The v1 schema drives the
// transformation — every field that is now a SingleNestedAttribute (Computed-only)
// or a SingleNestedBlock (Optional/Required, converted by SplitResourceSchemaBlocks)
// will have its value unwrapped from [{...}] to {...}.  List/set nested fields are
// recursed into so that deeply nested single-objects are also fixed.
func UnwrapSingleNestedStateUpgrader(schema map[string]resourceschema.Attribute) resource.StateUpgrader {
	return resource.StateUpgrader{
		StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
			// 1. Parse the raw v0 state JSON into a generic map.
			var raw map[string]any
			if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
				resp.Diagnostics.AddError("State upgrade: failed to parse v0 state JSON", err.Error())
				return
			}

			// 2. Recursively unwrap list-of-1 arrays for SingleNested fields.
			attrs, blocks := SplitResourceSchemaBlocks(schema)
			unwrapSingleNestedInMap(raw, attrs, blocks)

			// 2b. Remove keys that no longer exist in the v1 schema.
			//     DynamicValue.Unmarshal is strict and will error on unknown keys.
			stripUnknownKeys(raw, attrs, blocks)

			// 3. Re-marshal the transformed state.
			transformed, err := json.Marshal(raw)
			if err != nil {
				resp.Diagnostics.AddError("State upgrade: failed to re-marshal v1 state JSON", err.Error())
				return
			}

			// 4. Decode the transformed JSON into a tftypes.Value using the v1 schema type.
			//    tfprotov6.DynamicValue understands Terraform's JSON wire format, which is
			//    the same format used by req.RawState.JSON.
			fullSchema := resourceschema.Schema{Attributes: attrs, Blocks: blocks}
			dv := tfprotov6.DynamicValue{JSON: transformed}
			val, err := dv.Unmarshal(fullSchema.Type().TerraformType(ctx))
			if err != nil {
				resp.Diagnostics.AddError("State upgrade: failed to decode transformed JSON into schema type", err.Error())
				return
			}

			// 5. Set the upgraded state directly via the raw tftypes.Value.
			resp.State.Raw = val
		},
	}
}

// UnwrapSingleNestedUpgraderMap is a convenience wrapper that returns the standard
// v0→v1 upgrader map for use in DataSourceType.StateUpgraders.  Callers do not
// need to import the resource package directly.
func UnwrapSingleNestedUpgraderMap(schema map[string]resourceschema.Attribute) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: UnwrapSingleNestedStateUpgrader(schema),
	}
}

// unwrapSingleNestedInMap recursively walks a JSON-decoded state map and:
//  1. Converts any list-of-1 values ([{...}]) into plain objects ({...}) for
//     fields that are SingleNestedAttribute or SingleNestedBlock in the v1 schema.
//  2. Converts any list-of-1 scalar values ([x]) into plain scalars (x) for
//     Float64Attribute and Int64Attribute fields — these were the SDK v2 workaround
//     for optional numeric values (e.g. "lower_limit": [5] → "lower_limit": 5).
func unwrapSingleNestedInMap(
	state map[string]any,
	attrs map[string]resourceschema.Attribute,
	blocks map[string]resourceschema.Block,
) {
	for key, attr := range attrs {
		switch v := attr.(type) {
		case resourceschema.SingleNestedAttribute:
			unwrapField(state, key, v.Attributes, nil)
		case resourceschema.ListNestedAttribute:
			// Computed-only list nested attributes stay as attributes (not blocks).
			// Recurse into each element so nested single-objects are unwrapped.
			if arr, ok := state[key].([]any); ok {
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
						unwrapSingleNestedInMap(obj, childAttrs, childBlocks)
					}
				}
			}
		case resourceschema.SetNestedAttribute:
			// Same as ListNestedAttribute — recurse into each element.
			if arr, ok := state[key].([]any); ok {
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
						unwrapSingleNestedInMap(obj, childAttrs, childBlocks)
					}
				}
			}
		case resourceschema.Float64Attribute:
			unwrapScalarList(state, key)
		case resourceschema.Int64Attribute:
			unwrapScalarList(state, key)
		case resourceschema.BoolAttribute:
			unwrapScalarList(state, key)
		}
	}
	for key, block := range blocks {
		switch v := block.(type) {
		case resourceschema.SingleNestedBlock:
			unwrapField(state, key, v.Attributes, v.Blocks)
		case resourceschema.ListNestedBlock:
			if arr, ok := state[key].([]any); ok {
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						unwrapSingleNestedInMap(obj, v.NestedObject.Attributes, v.NestedObject.Blocks)
					}
				}
			}
		}
	}
}

// unwrapScalarList converts "[x]" → x and "[]" → null for optional scalar
// (Float64/Int64) fields that were stored as list-of-1 in SDK v2.
func unwrapScalarList(state map[string]any, key string) {
	raw, exists := state[key]
	if !exists || raw == nil {
		return
	}
	if arr, ok := raw.([]any); ok {
		if len(arr) == 1 {
			state[key] = arr[0]
		} else {
			state[key] = nil
		}
	}
}

// list-of-1, it is replaced with the inner object and recursed into.  If it is
// already a plain object it is just recursed into (idempotent).
func unwrapField(state map[string]any,
	key string,
	childAttrs map[string]resourceschema.Attribute,
	childBlocks map[string]resourceschema.Block,
) {
	raw, exists := state[key]
	if !exists || raw == nil {
		return
	}
	if arr, ok := raw.([]any); ok {
		if len(arr) == 1 {
			if obj, ok := arr[0].(map[string]any); ok {
				unwrapSingleNestedInMap(obj, childAttrs, childBlocks)
				state[key] = obj
			}
		} else {
			// Empty list — treat as absent/null.
			state[key] = nil
		}
		return
	}
	// Already a plain object — recurse for nested single fields.
	if obj, ok := raw.(map[string]any); ok {
		unwrapSingleNestedInMap(obj, childAttrs, childBlocks)
	}
}

// stripUnknownKeys removes any keys from the state map that are not present in
// the v1 schema (attrs + blocks).  This prevents DynamicValue.Unmarshal from
// failing when the old SDK v2 state contained fields that were subsequently
// removed from the schema entirely (e.g. the "constraints" field on resources).
// It recurses into list/set nested blocks so that removed child fields are also
// cleaned up.
func stripUnknownKeys(
	state map[string]any,
	attrs map[string]resourceschema.Attribute,
	blocks map[string]resourceschema.Block,
) {
	for key := range state {
		_, inAttrs := attrs[key]
		_, inBlocks := blocks[key]
		if !inAttrs && !inBlocks {
			delete(state, key)
			continue
		}
		// Recurse into nested objects/blocks.
		switch v := attrs[key].(type) {
		case resourceschema.SingleNestedAttribute:
			if obj, ok := state[key].(map[string]any); ok {
				childAttrs, childBlocks := SplitResourceSchemaBlocks(v.Attributes)
				stripUnknownKeys(obj, childAttrs, childBlocks)
			}
		case resourceschema.ListNestedAttribute:
			if arr, ok := state[key].([]any); ok {
				childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						stripUnknownKeys(obj, childAttrs, childBlocks)
					}
				}
			}
		case resourceschema.SetNestedAttribute:
			if arr, ok := state[key].([]any); ok {
				childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						stripUnknownKeys(obj, childAttrs, childBlocks)
					}
				}
			}
		}
		switch v := blocks[key].(type) {
		case resourceschema.SingleNestedBlock:
			if obj, ok := state[key].(map[string]any); ok {
				stripUnknownKeys(obj, v.Attributes, v.Blocks)
			}
		case resourceschema.ListNestedBlock:
			if arr, ok := state[key].([]any); ok {
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						stripUnknownKeys(obj, v.NestedObject.Attributes, v.NestedObject.Blocks)
					}
				}
			}
		}
	}
}
