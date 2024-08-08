package models

type CreateProductResponse struct {
	Product
}

type SearchProductResponse struct {
	Paging  PagingResponse `json:"paging"`
	Results []Product      `json:"results"`
}

type PagingResponse struct {
	Total  int64 `json:"total"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
