package metrics

import (
	metricstypes "github.com/leanspace/terraform-provider-leanspace/services/metrics/metrics/types"
)

// Metric is a type alias for the SDK-free struct in the types subpackage
// (see its doc comment), so external tools (e.g. xtceimporter) can import
// the plain JSON shape for compile-time drift checking without pulling in
// terraform-plugin-sdk, the same way general_objects.DefinitionAttribute
// works.
type Metric[T any] = metricstypes.Metric[T]
