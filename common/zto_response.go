package common

type ZTOResponse struct {
	OrderID  string `json:"orderId"`
	BillCode string `json:"billCode"`
	Update   bool   `json:"update,omitempty"`
	SiteCode string `json:"siteCode,omitempty"`
	SiteName string `json:"siteName,omitempty"`
}
