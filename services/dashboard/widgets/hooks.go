package widgets

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (widget *Widget) PostReadProcess(_ *provider.Client, newValue any) error {
	newWidget, ok := newValue.(*Widget)
	if !ok || newWidget == nil {
		return nil
	}
	for i := range newWidget.Series {
		if i >= len(widget.Series) {
			break
		}
		newWidget.Series[i].Filters = helper.ReorderByKey(
			widget.Series[i].Filters,
			newWidget.Series[i].Filters,
			func(f Filter) string { return f.FilterBy + "|" + f.Operator + "|" + f.Value },
		)
	}
	return nil
}
