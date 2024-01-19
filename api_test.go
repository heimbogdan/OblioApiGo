package oblioapigo

import (
	"encoding/json"
	"testing"
)

func TestCompaniesResponse_GetData(t *testing.T) {
	jsonData := `{"status": 200,"statusMessage":"Success","data":[{"cif":"RO12345678","company":"Test Company","userTypeAccess":"admin"}]}`
	companyResponse := &CompanyResponse{}
	err := json.Unmarshal([]byte(jsonData), &companyResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := companyResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	company := (*resp)[0]
	expected := `Test Company`
	if company.Company != expected {
		t.Fatal(`company was: `, company.Company, `, expected: `, expected)
	}
	t.Log(company)
}

func TestClientsResponse_GetData(t *testing.T) {
	jsonData := `{"status":200,"statusMessage":"Success","data":[{"cif":"RO12345678","name":"Test Company","rc":"J13/887/2017","code":"","address":"","state":"Constanta","city":"Constanta","country":"","iban":"","bank":"","email":"","phone":"","contact":"","vatPayer":true}]}`
	clientsResponse := &ClientsResponse{}
	err := json.Unmarshal([]byte(jsonData), &clientsResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := clientsResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	client := (*resp)[0]
	cif := `RO12345678`
	if client.Cif != cif {
		t.Fatal(`company was: `, client.Cif, `, expected: `, cif)
	}
	name := `Test Company`
	if client.Name != name {
		t.Fatal(`company was: `, client.Name, `, expected: `, name)
	}
	t.Log(client)
}
func TestProductsResponse_GetData(t *testing.T) {
	jsonData := `{"status":200,"statusMessage":"Success","data":[{"name":"Montare","code":"","measuringUnit":"buc","productType":"Serviciu","price":"119.00","currency":"RON","vatName":"Normala","vatPercentage":19,"vatIncluded":true},{"name":"Birou","code":"","description":"","measuringUnit":"buc","productType":"Marfa","stock":[{"workStation":"Sediu","management":"Magazin","quantity":2,"price":"200.00","currency":"RON","vatName":"Normala","vatPercentage":19,"vatIncluded":false}]}]}`
	productsResponse := &ProductsResponse{}
	err := json.Unmarshal([]byte(jsonData), &productsResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := productsResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	product1 := (*resp)[0]
	name := `Montare`
	if product1.Name != name {
		t.Fatal(`name was: `, product1.Name, `, expected: `, name)
	}
	if product1.Code != "" {
		t.Fatal(`code is not empty`)
	}
	if product1.Description != "" {
		t.Fatal(`description is not empty`)
	}
	t.Log(*resp)
}

func TestSeriesResponse_GetData(t *testing.T) {
	jsonData := `{"status":200,"statusMessage":"Success","data":[{"type":"Factura","name":"FCT","start":"0001","next":"0051","default":true},{"type":"Proforma","name":"PR","start":"0001","next":"0008","default":true}]}`
	seriesResponse := &SeriesResponse{}
	err := json.Unmarshal([]byte(jsonData), &seriesResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := seriesResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	series := (*resp)[0]
	name := `FCT`
	if series.Name != name {
		t.Fatal(`name was: `, series.Name, `, expected: `, name)
	}
	expectedType := `Factura`
	if series.Type != expectedType {
		t.Fatal(`type was: `, series.Type, `, expected: `, expectedType)
	}
	t.Log(*resp)
}

func TestLanguagesResponse_GetData(t *testing.T) {
	jsonData := `{"status":200,"statusMessage":"Success","data":[{"code":"EN","name":"Engleza"},{"code":"FR","name":"Franceza"}]}`
	languagesResponse := &LanguagesResponse{}
	err := json.Unmarshal([]byte(jsonData), &languagesResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := languagesResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	series := (*resp)[0]
	name := `Engleza`
	if series.Name != name {
		t.Fatal(`name was: `, series.Name, `, expected: `, name)
	}
	code := `EN`
	if series.Code != code {
		t.Fatal(`code was: `, series.Code, `, expected: `, code)
	}
	t.Log(*resp)
}

func TestProformaSummaryResponse_GetData(t *testing.T) {
	jsonData := `{"status":200,"statusMessage":"Success","data":{"seriesName":"PR","number":"0008","link":"https://www.oblio.eu/utils/show_file/?ic=91467&id=261747&it=b8590b71a702ab3788f0dda5d5c77d23"}}`
	proformaSummaryResponse := &ProformaSummaryResponse{}
	err := json.Unmarshal([]byte(jsonData), &proformaSummaryResponse.ApiResponse)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := proformaSummaryResponse.GetData()
	if err != nil {
		t.Error(err)
	}
	seriesName := `PR`
	if resp.SeriesName != seriesName {
		t.Fatal(`seriesName was: `, resp.SeriesName, `, expected: `, seriesName)
	}
	number := `0008`
	if resp.Number != number {
		t.Fatal(`number was: `, resp.Number, `, expected: `, number)
	}
	t.Log(*resp)
}
