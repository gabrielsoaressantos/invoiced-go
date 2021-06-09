package invdendpoint

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnMarshalWebhookRequest(t *testing.T) {
	s := `{
    "attempts": [{
        "status_code": 200,
        "timestamp": 1623079116
    }],
    "created_at": 1623079113,
    "event_id": 74419068,
    "id": 1755723,
    "next_attempt": null,
    "payload": {
        "id": 74419068,
        "timestamp": 1623079113,
        "type": "payment.created",
        "data": {
            "object": {
                "ach_sender_id": null,
                "amount": 384.45,
                "balance": 0,
                "created_at": 1623079091,
                "currency": "usd",
                "customer": {
                    "ach_gateway_id": null,
                    "address1": "612 IAA Dr.",
                    "address2": null,
                    "attention_to": null,
                    "autopay": true,
                    "autopay_delay_days": -1,
                    "avalara_entity_use_code": null,
                    "avalara_exemption_number": null,
                    "bill_to_parent": false,
                    "cc_gateway_id": null,
                    "chase": false,
                    "chasing_cadence": null,
                    "city": "Bloomington",
                    "consolidated": false,
                    "country": "US",
                    "created_at": 1611717943,
                    "credit_hold": false,
                    "credit_limit": null,
                    "currency": null,
                    "email": "drtatro@fairwayknolls.com",
                    "id": 2236884,
                    "language": "en",
                    "name": "Fairway Knolls Vet. Hospital",
                    "next_chase_step": null,
                    "notes": null,
                    "number": "C0055116",
                    "owner": null,
                    "parent_customer": null,
                    "payment_terms": null,
                    "phone": "309-663-1414",
                    "postal_code": "61701",
                    "state": "IL",
                    "tax_id": null,
                    "taxable": false,
                    "taxes": [],
                    "type": "company",
                    "object": "customer",
                    "statement_pdf_url": "https:\/\/covetrus.invoiced.com\/statements\/Mzn5YjKR9kVe0I9Y2IaAxuns\/pdf",
                    "sign_up_url": null,
                    "payment_source": {
                        "brand": "Visa",
                        "chargeable": true,
                        "created_at": 1613732154,
                        "exp_month": 5,
                        "exp_year": 2023,
                        "failure_reason": null,
                        "funding": "credit",
                        "gateway": "vantiv",
                        "gateway_customer": null,
                        "gateway_id": "327C230B-5F5C-47EC-A459-C351F5AC69FC",
                        "id": 572856,
                        "last4": "8363",
                        "merchant_account": 8366,
                        "receipt_email": null,
                        "object": "card"
                    },
                    "sign_up_page": null,
                    "metadata": {
                        "covetrusid": "C0055116",
                        "credit_balance": 0,
                        "D365_AHS": "321661",
                        "ElectronicDelivery": true,
                        "zero_communication": false
                    }
                },
                "date": 1623079091,
                "id": 1215844,
                "matched": null,
                "method": "credit_card",
                "notes": null,
                "reference": "1546591340",
                "source": "autopay",
                "voided": false,
                "object": "payment",
                "pdf_url": "https:\/\/covetrus.invoiced.com\/payments\/BnThMHn9htg4OK6fmOrPxMAg\/pdf",
                "charge": {
                    "amount": 384.45,
                    "amount_refunded": 0,
                    "created_at": 1623079091,
                    "currency": "usd",
                    "customer_id": 2236884,
                    "disputed": false,
                    "failure_message": null,
                    "gateway": "vantiv",
                    "gateway_id": "1546591340",
                    "id": 654683,
                    "payment_id": 1215844,
                    "receipt_email": null,
                    "refunded": false,
                    "status": "succeeded",
                    "payment_source": {
                        "brand": "Visa",
                        "chargeable": true,
                        "created_at": 1613732154,
                        "exp_month": 5,
                        "exp_year": 2023,
                        "failure_reason": null,
                        "funding": "credit",
                        "gateway": "vantiv",
                        "gateway_customer": null,
                        "gateway_id": "327C230B-5F5C-47EC-A459-C351F5AC69FC",
                        "id": 572856,
                        "last4": "8363",
                        "merchant_account": 8366,
                        "receipt_email": null,
                        "object": "card"
                    },
                    "refunds": []
                },
                "applied_to": [{
                    "id": 14905981,
                    "type": "invoice",
                    "amount": 384.45,
                    "invoice": 13053883
                }]
            }
        }
    },
    "url": "https:\/\/prod-45.eastus.logic.azure.com:443\/workflows\/71f92ed2eca64bb38d5154e521767032\/triggers\/request\/paths\/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Frequest%2Frun&sv=1.0&sig=Oq6ybMQ2pIj7wMkQ0hOwLiGwos63jHdREdNtUKSJoh8"
}`

	so := new(WebhookAttempt)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}
    b, _ := so.Payload.MarshalJSON()
	fmt.Println(string(b))
    fmt.Println(so.Attempts[0].StatusCode)
}

func TestUnMarshalWebhookRequestEmpty(t *testing.T) {
	s := `{
    "attempts": [],
    "created_at": 1623079113,
    "event_id": 74419068,
    "id": 1755723,
    "next_attempt": null,
    "payload": {
        "id": 74419068,
        "timestamp": 1623079113,
        "type": "payment.created",
        "data": {
            "object": {
                "ach_sender_id": null,
                "amount": 384.45,
                "balance": 0,
                "created_at": 1623079091,
                "currency": "usd",
                "customer": {
                    "ach_gateway_id": null,
                    "address1": "612 IAA Dr.",
                    "address2": null,
                    "attention_to": null,
                    "autopay": true,
                    "autopay_delay_days": -1,
                    "avalara_entity_use_code": null,
                    "avalara_exemption_number": null,
                    "bill_to_parent": false,
                    "cc_gateway_id": null,
                    "chase": false,
                    "chasing_cadence": null,
                    "city": "Bloomington",
                    "consolidated": false,
                    "country": "US",
                    "created_at": 1611717943,
                    "credit_hold": false,
                    "credit_limit": null,
                    "currency": null,
                    "email": "drtatro@fairwayknolls.com",
                    "id": 2236884,
                    "language": "en",
                    "name": "Fairway Knolls Vet. Hospital",
                    "next_chase_step": null,
                    "notes": null,
                    "number": "C0055116",
                    "owner": null,
                    "parent_customer": null,
                    "payment_terms": null,
                    "phone": "309-663-1414",
                    "postal_code": "61701",
                    "state": "IL",
                    "tax_id": null,
                    "taxable": false,
                    "taxes": [],
                    "type": "company",
                    "object": "customer",
                    "statement_pdf_url": "https:\/\/covetrus.invoiced.com\/statements\/Mzn5YjKR9kVe0I9Y2IaAxuns\/pdf",
                    "sign_up_url": null,
                    "payment_source": {
                        "brand": "Visa",
                        "chargeable": true,
                        "created_at": 1613732154,
                        "exp_month": 5,
                        "exp_year": 2023,
                        "failure_reason": null,
                        "funding": "credit",
                        "gateway": "vantiv",
                        "gateway_customer": null,
                        "gateway_id": "327C230B-5F5C-47EC-A459-C351F5AC69FC",
                        "id": 572856,
                        "last4": "8363",
                        "merchant_account": 8366,
                        "receipt_email": null,
                        "object": "card"
                    },
                    "sign_up_page": null,
                    "metadata": {
                        "covetrusid": "C0055116",
                        "credit_balance": 0,
                        "D365_AHS": "321661",
                        "ElectronicDelivery": true,
                        "zero_communication": false
                    }
                },
                "date": 1623079091,
                "id": 1215844,
                "matched": null,
                "method": "credit_card",
                "notes": null,
                "reference": "1546591340",
                "source": "autopay",
                "voided": false,
                "object": "payment",
                "pdf_url": "https:\/\/covetrus.invoiced.com\/payments\/BnThMHn9htg4OK6fmOrPxMAg\/pdf",
                "charge": {
                    "amount": 384.45,
                    "amount_refunded": 0,
                    "created_at": 1623079091,
                    "currency": "usd",
                    "customer_id": 2236884,
                    "disputed": false,
                    "failure_message": null,
                    "gateway": "vantiv",
                    "gateway_id": "1546591340",
                    "id": 654683,
                    "payment_id": 1215844,
                    "receipt_email": null,
                    "refunded": false,
                    "status": "succeeded",
                    "payment_source": {
                        "brand": "Visa",
                        "chargeable": true,
                        "created_at": 1613732154,
                        "exp_month": 5,
                        "exp_year": 2023,
                        "failure_reason": null,
                        "funding": "credit",
                        "gateway": "vantiv",
                        "gateway_customer": null,
                        "gateway_id": "327C230B-5F5C-47EC-A459-C351F5AC69FC",
                        "id": 572856,
                        "last4": "8363",
                        "merchant_account": 8366,
                        "receipt_email": null,
                        "object": "card"
                    },
                    "refunds": []
                },
                "applied_to": [{
                    "id": 14905981,
                    "type": "invoice",
                    "amount": 384.45,
                    "invoice": 13053883
                }]
            }
        }
    },
    "url": "https:\/\/prod-45.eastus.logic.azure.com:443\/workflows\/71f92ed2eca64bb38d5154e521767032\/triggers\/request\/paths\/invoke?api-version=2016-10-01&sp=%2Ftriggers%2Frequest%2Frun&sv=1.0&sig=Oq6ybMQ2pIj7wMkQ0hOwLiGwos63jHdREdNtUKSJoh8"
}`

	so := new(WebhookAttempt)

	err := json.Unmarshal([]byte(s), so)
	if err != nil {
		t.Fatal(err)
	}

	if so.Attempts == nil || len(so.Attempts) == 0 {
		fmt.Println("NIL")
	} else {
		fmt.Println(len(so.Attempts) == 0)
	}
}