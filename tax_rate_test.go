package invdapi

import (
	"reflect"
	"testing"
	"time"

	"github.com/gabrielsoaressantos/invoiced-go/invdendpoint"
	"github.com/gabrielsoaressantos/invoiced-go/invdmockserver"
)

func TestTaxRate_Create(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.TaxRate)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "nomenclature"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewTaxRate()

	requestEntity := entity.NewTaxRate()

	requestEntity.Id = "example"
	requestEntity.Name = "nomenclature"

	createdEntity, err := entity.Create(requestEntity)
	if err != nil {
		t.Fatal("Error Creating entity", err)
	}

	if !reflect.DeepEqual(createdEntity.TaxRate, mockResponse) {
		t.Fatal("entity was not created", createdEntity.TaxRate, mockResponse)
	}
}

func TestTaxRate_Save(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.TaxRate)
	mockResponse.Id = "example"
	mockResponse.CreatedAt = time.Now().UnixNano()
	mockResponse.Name = "new-name"

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)

	entityToUpdate := conn.NewTaxRate()

	entityToUpdate.Name = "new-name"

	err = entityToUpdate.Save()

	if err != nil {
		t.Fatal("Error updating entity", err)
	}

	if !reflect.DeepEqual(mockResponse, entityToUpdate.TaxRate) {
		t.Fatal("Error: entity not updated correctly")
	}
}

func TestTaxRate_Delete(t *testing.T) {
	key := "api key"

	mockResponse := ""
	mockResponseId := "example"

	server, err := invdmockserver.New(204, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}

	defer server.Close()

	conn := mockConnection(key, server)

	entity := conn.NewTaxRate()

	entity.Id = mockResponseId

	err = entity.Delete()

	if err != nil {
		t.Fatal("Error occurred deleting entity")
	}
}

func TestTaxRate_Retrieve(t *testing.T) {
	key := "test api key"

	mockResponse := new(invdendpoint.TaxRate)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewTaxRate()

	retrievedPayment, err := entity.Retrieve("example")
	if err != nil {
		t.Fatal("Error retrieving entity", err)
	}

	if !reflect.DeepEqual(retrievedPayment.TaxRate, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}

func TestTaxRate_ListAll(t *testing.T) {
	key := "test api key"

	var mockListResponse [1]invdendpoint.TaxRate

	mockResponse := new(invdendpoint.TaxRate)
	mockResponse.Id = "example"
	mockResponse.Name = "nomenclature"

	mockResponse.CreatedAt = time.Now().UnixNano()

	mockListResponse[0] = *mockResponse

	server, err := invdmockserver.New(200, mockListResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	conn := mockConnection(key, server)
	entity := conn.NewTaxRate()

	filter := invdendpoint.NewFilter()
	sorter := invdendpoint.NewSort()

	result, err := entity.ListAll(filter, sorter)
	if err != nil {
		t.Fatal("Error listing entity", err)
	}

	if !reflect.DeepEqual(result[0].TaxRate, mockResponse) {
		t.Fatal("Error messages do not match up")
	}
}
