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
// SDK v2 workaround for single nested objects. The v1 schema drives the
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

			// 2. In a single tree walk: unwrap list-of-1 SDK v2 workarounds and
			//    strip keys absent from the v1 schema (DynamicValue.Unmarshal is strict).
			attrs, blocks := SplitResourceSchemaBlocks(schema)
			processStateMap(raw, attrs, blocks)

			// 3–5. Marshal, decode into tftypes.Value, and set the upgraded state.
			marshalAndSetUpgradedState(ctx, raw, attrs, blocks, resp)
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

// StateUpgrader is a re-export of resource.StateUpgrader so callers in service
// packages can build upgrader maps without importing the resource package.
type StateUpgrader = resource.StateUpgrader

// UnwrapSingleNestedStateUpgraderWithTransforms is like UnwrapSingleNestedStateUpgrader
// but additionally applies zero or more raw-map transform functions after the standard
// unwrap/strip pass.  Each transform receives the JSON-decoded state map and may
// mutate it in place (e.g. to canonicalise list ordering).
func UnwrapSingleNestedStateUpgraderWithTransforms(schema map[string]resourceschema.Attribute, transforms ...func(map[string]any)) resource.StateUpgrader {
	return resource.StateUpgrader{
		StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
			var raw map[string]any
			if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
				resp.Diagnostics.AddError("State upgrade: failed to parse v0 state JSON", err.Error())
				return
			}

			attrs, blocks := SplitResourceSchemaBlocks(schema)
			processStateMap(raw, attrs, blocks)

			for _, fn := range transforms {
				fn(raw)
			}

			// 3–5. Marshal, decode into tftypes.Value, and set the upgraded state.
			marshalAndSetUpgradedState(ctx, raw, attrs, blocks, resp)
		},
	}
}

// marshalAndSetUpgradedState re-marshals the transformed state map, decodes it
// into a tftypes.Value using the v1 schema type, and sets it on the response.
func marshalAndSetUpgradedState(
	ctx context.Context,
	raw map[string]any,
	attrs map[string]resourceschema.Attribute,
	blocks map[string]resourceschema.Block,
	resp *resource.UpgradeStateResponse,
) {
	transformed, err := json.Marshal(raw)
	if err != nil {
		resp.Diagnostics.AddError("State upgrade: failed to re-marshal v1 state JSON", err.Error())
		return
	}

	fullSchema := resourceschema.Schema{Attributes: attrs, Blocks: blocks}
	dv := tfprotov6.DynamicValue{JSON: transformed}
	val, err := dv.Unmarshal(fullSchema.Type().TerraformType(ctx))
	if err != nil {
		resp.Diagnostics.AddError("State upgrade: failed to decode transformed JSON into schema type", err.Error())
		return
	}

	resp.State.Raw = val
}

// processStateMap performs both operations in a single tree walk:
//  1. Strips state keys absent from the v1 schema (attrs + blocks) so that
//     DynamicValue.Unmarshal does not fail on unknown fields.
//  2. Unwraps SDK v2 list-of-1 workarounds:
//     - ([{...}]) → {...} for SingleNestedAttribute / SingleNestedBlock fields.
//     - ([x]) → x and ([]) → null for Float64Attribute, Int64Attribute, BoolAttribute.
//
// The strip pass runs first at each level, then the unwrap+recurse pass descends
// into nested structures, so every nested object gets both operations applied.
func processStateMap(
	state map[string]any,
	attrs map[string]resourceschema.Attribute,
	blocks map[string]resourceschema.Block,
) {
	// Strip keys that are absent from the v1 schema at this level.
	for key := range state {
		_, inAttrs := attrs[key]
		_, inBlocks := blocks[key]
		if !inAttrs && !inBlocks {
			delete(state, key)
		}
	}
	// Unwrap SDK v2 workarounds and recurse into nested structures.
	for key, attr := range attrs {
		switch v := attr.(type) {
		case resourceschema.SingleNestedAttribute:
			processField(state, key, v.Attributes, nil)
		case resourceschema.ListNestedAttribute:
			if arr, ok := state[key].([]any); ok {
				childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						processStateMap(obj, childAttrs, childBlocks)
					}
				}
			}
		case resourceschema.SetNestedAttribute:
			if arr, ok := state[key].([]any); ok {
				childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						processStateMap(obj, childAttrs, childBlocks)
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
			processField(state, key, v.Attributes, v.Blocks)
		case resourceschema.ListNestedBlock:
			if arr, ok := state[key].([]any); ok {
				for _, elem := range arr {
					if obj, ok := elem.(map[string]any); ok {
						processStateMap(obj, v.NestedObject.Attributes, v.NestedObject.Blocks)
					}
				}
			}
		}
	}
}

// unwrapScalarList converts "[x]" → x and "[]" → null for optional scalar
// (Float64/Int64/Bool) fields that were stored as list-of-1 in SDK v2.
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

// processField handles a single nested field in the state map. If the current
// value is a list-of-1 ([{...}]), it is replaced with the inner object and
// recursed into via processStateMap. If it is already a plain object it is just
// recursed into (idempotent).
func processField(
	state map[string]any,
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
				processStateMap(obj, childAttrs, childBlocks)
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
		processStateMap(obj, childAttrs, childBlocks)
	}
}
