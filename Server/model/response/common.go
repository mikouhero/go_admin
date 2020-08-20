package response

type PageResult struct {
	List     interface{} `json:"list"`
	Totle    int         `json:"totle"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
