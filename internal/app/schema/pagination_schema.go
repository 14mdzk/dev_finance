package schema

type PaginationReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
