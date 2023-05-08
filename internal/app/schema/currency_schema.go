package schema

type CurrencyReq struct {
	Name         string `json:"name" validate:"required,alpha"`
	Abbreviation string `json:"abbreviation" validate:"required,alpha"`
}

type CurrencyResp struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}
