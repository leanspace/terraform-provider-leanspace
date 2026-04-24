package widgets

import (
	"sort"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var WidgetDataType = provider.DataSourceType[Widget, *Widget]{
	ResourceIdentifier: "leanspace_widgets",
	Path:               "dashboard-repository/widgets",
	Schema:             widgetSchema,
	FilterSchema:       dataSourceFilterSchema,
	TFModelFactory:     func() any { return &WidgetTF{} },
	SchemaVersion:      1,
	StateUpgraders: map[int64]provider.StateUpgrader{
		0: provider.UnwrapSingleNestedStateUpgraderWithTransforms(widgetSchema, sortFiltersInState),
	},
}

// sortFiltersInState sorts the filters within each series element of the raw
// state map canonically by (filter_by, operator, value).  This is applied
// during the v0→v1 state upgrade so that the stored order matches the order
// produced by ToTF() on every subsequent read, preventing "inconsistent result
// after apply" errors.
func sortFiltersInState(raw map[string]any) {
	seriesList, ok := raw["series"].([]any)
	if !ok {
		return
	}
	for _, seriesElem := range seriesList {
		seriesObj, ok := seriesElem.(map[string]any)
		if !ok {
			continue
		}
		filters, ok := seriesObj["filters"].([]any)
		if !ok || len(filters) < 2 {
			continue
		}
		sort.SliceStable(filters, func(a, b int) bool {
			fa, _ := filters[a].(map[string]any)
			fb, _ := filters[b].(map[string]any)
			fba, _ := fa["filter_by"].(string)
			fbb, _ := fb["filter_by"].(string)
			if fba != fbb {
				return fba < fbb
			}
			opa, _ := fa["operator"].(string)
			opb, _ := fb["operator"].(string)
			if opa != opb {
				return opa < opb
			}
			va, _ := fa["value"].(string)
			vb, _ := fb["value"].(string)
			return va < vb
		})
	}
}
