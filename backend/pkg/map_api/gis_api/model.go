package gisapi

type CompanyList struct {
	Result struct {
		Items []struct {
			AddressComment string `json:"address_comment,omitempty"`
			AddressName    string `json:"address_name"`
			ID             string `json:"id"`
			Name           string `json:"name"`
			Type           string `json:"type"`
		} `json:"items"`
		Total int `json:"total"`
	} `json:"result"`
}
