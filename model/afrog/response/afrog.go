package response

// PocListResponse Poc列表返回
type PocListResponse struct {
	Index int `json:"index"`
	Id string `json:"id"`
	Name string `json:"name"`
	Severity string `json:"severity"`
	Description string `json:"description"`
}