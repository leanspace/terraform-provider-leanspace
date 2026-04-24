package processors

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct Processor

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Processor struct {
	general_objects.AuditModel
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Version     string  `json:"version"`
	Type        string  `json:"type"`
	FilePath    string  `json:"filePath"`
	FileSha     string  `json:"fileSha"`
}

type ProcessorUrl struct {
	Url     string `json:"url"`
	Expires string `json:"expires"`
}

type AttachedRoute struct {
	Content []AttachedRouteContent `json:"content"`
}

type AttachedRouteContent struct {
	ID string `json:"id"`
}
