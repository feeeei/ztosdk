package common

type ZTOUpdateResponse struct {
	OrderCode string `json:"orderCode"`
	FieldName string `json:"fieldName"`
	Result    bool   `json:"result"`
	Reason    string `json:"reason"`
}
