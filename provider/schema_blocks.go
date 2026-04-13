package provider

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// hasComputedField reports whether any attribute in the given map (or its nested
// descendants) has Computed set to true. This drives the SetNestedAttribute →
// SetNestedBlock vs ListNestedBlock decision: a set whose elements never contain
// server-computed fields is safe to expose as SetNestedBlock (Terraform hashes
// elements by content so ordering is irrelevant). When at least one child is
// Computed, a ListNestedBlock is used instead to avoid unknown-hash plan failures.
/*func hasComputedField(attrs map[string]resourceschema.Attribute) bool {
	for _, a := range attrs {
		switch v := a.(type) {
		case resourceschema.StringAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.BoolAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.Int64Attribute:
			if v.Computed {
				return true
			}
		case resourceschema.Float64Attribute:
			if v.Computed {
				return true
			}
		case resourceschema.NumberAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.ListAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.SetAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.MapAttribute:
			if v.Computed {
				return true
			}
		case resourceschema.ListNestedAttribute:
			if v.Computed || hasComputedField(v.NestedObject.Attributes) {
				return true
			}
		case resourceschema.SetNestedAttribute:
			if v.Computed || hasComputedField(v.NestedObject.Attributes) {
				return true
			}
		case resourceschema.SingleNestedAttribute:
			if v.Computed || hasComputedField(v.Attributes) {
				return true
			}
		}
	}
	return false
}*/

// SplitResourceSchemaBlocks takes a flat map of resource attributes and extracts
// any non-Computed nested attributes (ListNested, SetNested, SingleNested) into the Blocks map.
// This preserves HCL block syntax compatibility with the old SDK v2 provider.
// Computed-only nested attributes remain as attributes since blocks cannot be Computed.
// SetNestedAttribute with no Computed element fields becomes SetNestedBlock (order-independent);
// those with Computed element fields become ListNestedBlock (avoids unknown-hash plan failures).
func SplitResourceSchemaBlocks(attrs map[string]resourceschema.Attribute) (map[string]resourceschema.Attribute, map[string]resourceschema.Block) {
	outAttrs := make(map[string]resourceschema.Attribute)
	outBlocks := make(map[string]resourceschema.Block)

	for key, a := range attrs {
		switch v := a.(type) {
		case resourceschema.ListNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
			outBlocks[key] = resourceschema.ListNestedBlock{
				NestedObject: resourceschema.NestedBlockObject{
					Attributes: childAttrs,
					Blocks:     childBlocks,
				},
			}
		case resourceschema.SetNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitResourceSchemaBlocks(v.NestedObject.Attributes)
			outBlocks[key] = resourceschema.ListNestedBlock{
				NestedObject: resourceschema.NestedBlockObject{
					Attributes: childAttrs,
					Blocks:     childBlocks,
				},
			}
			/*if hasComputedField(v.NestedObject.Attributes) {
				// Has server-computed element fields — use ListNestedBlock to avoid
				// unknown-hash plan failures during planning.
				outBlocks[key] = resourceschema.ListNestedBlock{
					NestedObject: resourceschema.NestedBlockObject{
						Attributes: childAttrs,
						Blocks:     childBlocks,
					},
				}
			} else {
				// No Computed element fields — use SetNestedBlock so Terraform matches
				// elements by content hash, making API-side reordering a non-issue.
				outBlocks[key] = resourceschema.SetNestedBlock{
					NestedObject: resourceschema.NestedBlockObject{
						Attributes: childAttrs,
						Blocks:     childBlocks,
					},
				}
			}*/
		case resourceschema.SingleNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitResourceSchemaBlocks(v.Attributes)
			outBlocks[key] = resourceschema.SingleNestedBlock{
				Attributes: childAttrs,
				Blocks:     childBlocks,
			}
		default:
			outAttrs[key] = a
		}
	}

	return outAttrs, outBlocks
}

// SplitDatasourceSchemaBlocks does the same as SplitResourceSchemaBlocks but for datasource schemas.
func SplitDatasourceSchemaBlocks(attrs map[string]datasourceschema.Attribute) (map[string]datasourceschema.Attribute, map[string]datasourceschema.Block) {
	outAttrs := make(map[string]datasourceschema.Attribute)
	outBlocks := make(map[string]datasourceschema.Block)

	for key, a := range attrs {
		switch v := a.(type) {
		case datasourceschema.ListNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitDatasourceSchemaBlocks(v.NestedObject.Attributes)
			outBlocks[key] = datasourceschema.ListNestedBlock{
				NestedObject: datasourceschema.NestedBlockObject{
					Attributes: childAttrs,
					Blocks:     childBlocks,
				},
			}
		case datasourceschema.SetNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitDatasourceSchemaBlocks(v.NestedObject.Attributes)
			// Use ListNestedBlock instead of SetNestedBlock (same reason as resource schema).
			outBlocks[key] = datasourceschema.ListNestedBlock{
				NestedObject: datasourceschema.NestedBlockObject{
					Attributes: childAttrs,
					Blocks:     childBlocks,
				},
			}
		case datasourceschema.SingleNestedAttribute:
			if v.Computed && !v.Optional && !v.Required {
				outAttrs[key] = a
				continue
			}
			childAttrs, childBlocks := SplitDatasourceSchemaBlocks(v.Attributes)
			outBlocks[key] = datasourceschema.SingleNestedBlock{
				Attributes: childAttrs,
				Blocks:     childBlocks,
			}
		default:
			outAttrs[key] = a
		}
	}

	return outAttrs, outBlocks
}
