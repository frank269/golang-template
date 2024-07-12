package dtos

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationResponse struct {
	Records      interface{} `json:"records"`
	CurrentPage  int         `json:"currentPage"`
	PageSize     int         `json:"pageSize"`
	TotalRecords int64       `json:"totalRecords"`
}
