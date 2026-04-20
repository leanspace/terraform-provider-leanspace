package provider

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SplitResourceSchemaBlocks takes a flat map of resource attributes and extracts
// any non-Computed nested attributes (ListNested, SetNested, SingleNested) into the Blocks map.
// This preserves HCL block syntax compatibility with the old SDK v2 provider.
// Computed-only nested attributes remain as attributes since blocks cannot be Computed.
// SetNestedAttribute with no Computed element fields becomes ListNestedBlock (avoids unknown-hash plan failures)
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
			outBlocks[key] = resourceschema.ListNestedBlock{ // implemented as list nested block to avoid unknown hash plan failures (since set with no computed fields cannot be tracked)
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

// ResourceSchemaToDataSource converts a resource attribute map to a datasource attribute map.
// Every attribute becomes Computed:true (data sources are read-only).
// Validators, plan modifiers and defaults are dropped since they are irrelevant for data sources.
func ResourceSchemaToDataSource(attrs map[string]resourceschema.Attribute) map[string]datasourceschema.Attribute {
	out := make(map[string]datasourceschema.Attribute, len(attrs))
	for k, a := range attrs {
		out[k] = resourceAttrToDS(a)
	}
	return out
}

func resourceAttrToDS(a resourceschema.Attribute) datasourceschema.Attribute {
	switch v := a.(type) {
	case resourceschema.StringAttribute:
		return datasourceschema.StringAttribute{Computed: true, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.BoolAttribute:
		return datasourceschema.BoolAttribute{Computed: true, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.Int64Attribute:
		return datasourceschema.Int64Attribute{Computed: true, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.Float64Attribute:
		return datasourceschema.Float64Attribute{Computed: true, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.ListAttribute:
		return datasourceschema.ListAttribute{Computed: true, ElementType: v.ElementType, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.SetAttribute:
		return datasourceschema.SetAttribute{Computed: true, ElementType: v.ElementType, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.MapAttribute:
		return datasourceschema.MapAttribute{Computed: true, ElementType: v.ElementType, Description: v.Description, MarkdownDescription: v.MarkdownDescription}
	case resourceschema.ListNestedAttribute:
		return datasourceschema.ListNestedAttribute{
			Computed:            true,
			Description:         v.Description,
			MarkdownDescription: v.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: ResourceSchemaToDataSource(v.NestedObject.Attributes),
			},
		}
	case resourceschema.SetNestedAttribute:
		return datasourceschema.SetNestedAttribute{
			Computed:            true,
			Description:         v.Description,
			MarkdownDescription: v.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: ResourceSchemaToDataSource(v.NestedObject.Attributes),
			},
		}
	case resourceschema.SingleNestedAttribute:
		return datasourceschema.SingleNestedAttribute{
			Computed:            true,
			Description:         v.Description,
			MarkdownDescription: v.MarkdownDescription,
			Attributes:          ResourceSchemaToDataSource(v.Attributes),
		}
	default:
		return datasourceschema.StringAttribute{Computed: true}
	}
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
