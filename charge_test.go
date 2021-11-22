package invdapi

import (
	"testing"
	"time"

	"github.com/gabrielsoaressantos/invoiced-go/invdendpoint"
	"github.com/gabrielsoaressantos/invoiced-go/invdmockserver"
)

func TestCreateCharge(t *testing.T) {
	key := "test api key"

	mockChargeResponseID := int64(1523)
	mockChargeResponse := new(invdendpoint.Payment)
	mockChargeResponse.Id = mockChargeResponseID
	mockChargeResponse.Customer = 234112
	mockChargeResponse.Reference = "234"

	mockChargeResponse.CreatedAt = time.Now().UnixNano()

	server, err := invdmockserver.New(200, mockChargeResponse, "json", true)
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	conn := mockConnection(key, server)

	chargeRequest := new(invdendpoint.ChargeRequest)

	charge, err := conn.NewCharge().Create(chargeRequest)
	if err != nil {
		t.Fatal("Error Creating charge", err)
	}

	if charge.Id != int64(1523) {
		t.Fatal("Error Messages Do Not Match Up")
	}
}
