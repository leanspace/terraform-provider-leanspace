package general_objects

type Sort struct {
	Direction    string `json:"direction" terra:"direction"`
	Property     string `json:"property" terra:"property"`
	IgnoreCase   bool   `json:"ignoreCase" terra:"ignore_case"`
	NullHandling string `json:"nullHandling" terra:"null_handling"`
	Ascending    bool   `json:"ascending" terra:"ascending"`
	Descending   bool   `json:"descending" terra:"descending"`
}

type Pageable struct {
	Sort       []Sort `json:"sort" terra:"sort"`
	Offset     int    `json:"offset" terra:"offset"`
	PageNumber int    `json:"pageNumber" terra:"page_number"`
	PageSize   int    `json:"pageSize" terra:"page_size"`
	Paged      bool   `json:"paged" terra:"paged"`
	Unpaged    bool   `json:"unpaged" terra:"unpaged"`
}

type Tag struct {
	Key   string `json:"key" terra:"key"`
	Value string `json:"value,omitempty" terra:"value"`
}

type PaginatedList[T any] struct {
	Content          []T      `json:"content" terra:"content"`
	TotalElements    int      `json:"totalElements" terra:"total_elements"`
	TotalPages       int      `json:"totalPages" terra:"total_pages"`
	NumberOfElements int      `json:"numberOfElements" terra:"number_of_elements"`
	Number           int      `json:"number" terra:"number"`
	Size             int      `json:"size" terra:"size"`
	Sort             []Sort   `json:"sort" terra:"sort"`
	First            bool     `json:"first" terra:"first"`
	Last             bool     `json:"last" terra:"last"`
	Empty            bool     `json:"empty" terra:"empty"`
	Pageable         Pageable `json:"pageable" terra:"pageable"`
}
