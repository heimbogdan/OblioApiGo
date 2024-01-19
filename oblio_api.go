package oblioapigo

import (
	"encoding/json"
	"errors"
)

type AccessToken struct {
	Token       string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	RequestTime int64  `json:"request_time"`
	Type        string `json:"token_type"`
	Scope       string `json:"scope,omitempty"`
}

type ApiResponse struct {
	Status        int         `json:"status"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data"`
}

func getDataAs[T any](data interface{}, t T) (*T, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("data cannot be marshaled")
	}
	err = json.Unmarshal(jsonString, &t)
	if err != nil {
		return nil, errors.New(`data cannot be unmarshalled"`)
	}
	return &t, nil
}

type Company struct {
	Cif            string `json:"cif"`
	Company        string `json:"company"`
	UserTypeAccess string `json:"userTypeAccess"`
}

type CompanyResponse struct {
	ApiResponse
}

func (c *CompanyResponse) GetData() (*[]Company, error) {
	return getDataAs(c.Data, make([]Company, 0))
}

type Client struct {
	NameCodeType
	Cif      string `json:"cif,omitempty"`
	Rc       string `json:"rc,omitempty"`
	Address  string `json:"address,omitempty"`
	State    string `json:"state,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
	Iban     string `json:"iban,omitempty"`
	Bank     string `json:"bank,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Contact  string `json:"contact,omitempty"`
	VatPayer bool   `json:"vatPayer"`
	Save     bool   `json:"save,omitempty"`
}

type ClientsResponse struct {
	ApiResponse
}

func (c *ClientsResponse) GetData() (*[]Client, error) {
	return getDataAs(c.Data, make([]Client, 0))
}

type VatType struct {
	VatName       string `json:"vatName"`
	VatPercentage int    `json:"vatPercentage"`
	VatIncluded   bool   `json:"vatIncluded"`
}

type PriceQuantityType struct {
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

type Product struct {
	CurrencyType
	VatType
	PriceQuantityType
	NameCodeType
	NameTranslation          string  `json:"nameTranslation,omitempty"`
	Description              string  `json:"description,omitempty"`
	MeasuringUnit            string  `json:"measuringUnit"`
	MeasuringUnitTranslation string  `json:"measuringUnitTranslation,omitempty"`
	ProductType              string  `json:"productType"`
	Stock                    []Stock `json:"stock,omitempty"`
	Management               string  `json:"management,omitempty"`
	Discount                 int     `json:"discount,omitempty"`
	DiscountType             string  `json:"discountType,omitempty"`
	Save                     bool    `json:"save,omitempty"`
}

type Stock struct {
	VatType
	PriceQuantityType
	WorkStation string `json:"workStation"`
	Management  string `json:"management"`
	Currency    string `json:"currency"`
}

type ProductsResponse struct {
	ApiResponse
}

func (c *ProductsResponse) GetData() (*[]Product, error) {
	return getDataAs(c.Data, make([]Product, 0))
}

type Series struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Start   string `json:"start"`
	Next    string `json:"next"`
	Default bool   `json:"default"`
}

type SeriesResponse struct {
	ApiResponse
}

func (c *SeriesResponse) GetData() (*[]Series, error) {
	return getDataAs(c.Data, make([]Series, 0))
}

type NameCodeType struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Language struct {
	NameCodeType
}

type LanguagesResponse struct {
	ApiResponse
}

func (c *LanguagesResponse) GetData() (*[]Language, error) {
	return getDataAs(c.Data, make([]Language, 0))
}

type Management struct {
	Management     string `json:"management"`
	WorkStation    string `json:"workStation"`
	ManagementType string `json:"managementType"`
}

type ManagementResponse struct {
	ApiResponse
}

func (c *ManagementResponse) GetData() (*[]Management, error) {
	return getDataAs(c.Data, make([]Management, 0))
}

type CurrencyType struct {
	Currency     string  `json:"currency"`
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
}

type Proforma struct {
	CurrencyType
	Cif                string    `json:"cif"`
	Client             Client    `json:"client"`
	IssueDate          string    `json:"issueDate"`
	DueDate            string    `json:"dueDate"`
	SeriesName         string    `json:"seriesName"`
	Language           string    `json:"language"`
	Precision          int       `json:"precision"`
	Products           []Product `json:"products"`
	IssuerName         string    `json:"issuerName"`
	IssuerId           int64     `json:"issuerId"`
	NoticeNumber       string    `json:"noticeNumber"`
	InternalNote       string    `json:"internalNote"`
	DeputyName         string    `json:"deputyName"`
	DeputyIdentityCard string    `json:"deputyIdentityCard"`
	DeputyAuto         string    `json:"deputyAuto"`
	SelesAgent         string    `json:"selesAgent"`
	Mentions           string    `json:"mentions"`
	WorkStation        string    `json:"workStation"`
	SendEmail          bool      `json:"sendEmail"`
}

type SeriesNameNumber struct {
	SeriesName string `json:"seriesName"`
	Number     string `json:"number"`
}

type ProformaSummary struct {
	SeriesNameNumber
	Link string `json:"link"`
}

type ProformaSummaryResponse struct {
	ApiResponse
}

func (c *ProformaSummaryResponse) GetData() (*ProformaSummary, error) {
	return getDataAs(c.Data, ProformaSummary{})
}

type Notice struct {
	Proforma
	UseStock bool `json:"useStock,omitempty"`
}

type Invoice struct {
	Notice
	DeliveryDate      string            `json:"deliveryDate,omitempty"`
	CollectDate       string            `json:"collectDate,omitempty"`
	ReferenceDocument ReferenceDocument `json:"referenceDocument,omitempty"`
	Collect           Collect           `json:"collect,omitempty"`
}

type ReferenceDocument struct {
	SeriesNameNumber
	Type string `json:"type,omitempty"`
}

type Collect struct {
	Type           string  `json:"type,omitempty"`
	SeriesName     string  `json:"seriesName,omitempty"`
	DocumentNumber string  `json:"documentNumber,omitempty"`
	Value          float64 `json:"value"`
	IssueDate      string  `json:"issueDate,omitempty"`
	Mentions       string  `json:"mentions,omitempty"`
}
