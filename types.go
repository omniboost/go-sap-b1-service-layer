package sap

type BusinessPartner struct {
	OdataContext string `json:"@odata.context,omitempty"`
	CardCode     string
	CardName     string
	FederalTaxID string
	Phone1       string
	Phone2       string
	EmailAddress string
	Address      string
	ZipCode      string
	City         string
	Country      string
	CardType     string `json:"CardType,omitempty"`
}

type Invoice struct {
	DocDate       string        `json:"DocDate"`
	DocDueDate    string        `json:"DocDueDate"`
	CardCode      string        `json:"CardCode"`
	Comments      string        `json:"Comments"`
	DocumentLines DocumentLines `json:"DocumentLines"`
}

type DocumentLines []DocumentLine

type DocumentLine struct {
	ItemCode        string `json:"ItemCode"`
	ItemDescription string `json:"ItemDescription,omitempty"`
	UnitPrice       string `json:"UnitPrice"`
	Quantity        string `json:"Quantity"`
	DiscountPercent string `json:"DiscountPercent"`
	VatRate         string `json:"VatRate"`
	FreeText        string `json:"FreeText"`
	PriceAfterVAT   string `json:"PriceAfterVAT"`
	ItemsGroupCode  string `json:"ItemsGroupCode"`
	Ccoste          string `json:"Ccoste"`
	Currency        string `json:"Currency"`
}
