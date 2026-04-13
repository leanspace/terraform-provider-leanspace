package provider

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

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
