package reeleezee_test

import (
	"log"
	"os"
	"testing"

	null "gopkg.in/guregu/null.v3"

	invoicexpress "github.com/tim-online/go-invoicexpress"
)

func TestInvoicesListAll(t *testing.T) {
	accountName := os.Getenv("INVOICEXPRESS_ACCOUNTNAME")
	token := os.Getenv("INVOICEXPRESS_TOKEN")
	api := invoicexpress.NewAPI(nil, accountName, token)
	api.SetDebug(true)

	req := api.Invoices.NewListAllRequest()
	// req.QueryParams().DateFrom = invoicexpress.NewDate(2011, 1, 12)
	req.QueryParams().Text = "Kees"
	req.QueryParams().Types = append(req.QueryParams().Types, invoicexpress.Invoice)
	req.QueryParams().Status = append(req.QueryParams().Status, invoicexpress.Draft)
	req.QueryParams().NonArchived = null.NewBool(false, true)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	log.Println(resp)

	// req := api.Invoices.NewCreateRequest()
	// req.Params().SetDocumentType(invoicexpress.Invoice)
	// req.RequestBody().Invoice.Client.Name = "Kees"
	// req.RequestBody().Invoice = invoicexpress.NewInvoice{
	// 	Date:                 date.Today(),
	// 	DueDate:              date.Today(),
	// 	Reference:            "",
	// 	Observations:         "",
	// 	Retention:            "",
	// 	TaxExemption:         invoicexpress.M01,
	// 	SequenceID:           "",
	// 	ManualSequenceNumber: "",
	// 	Client: invoicexpress.NewClient{
	// 		Name:         "Kees Zorge",
	// 		Code:         "",
	// 		Email:        "info@omniboost.io",
	// 		Address:      "Axelsestraat 4",
	// 		City:         "Zaamslag",
	// 		PostalCode:   "4543CJ",
	// 		Country:      "Niad",
	// 		FiscalID:     "",
	// 		Website:      "https://www.omniboost.io",
	// 		Phone:        "0031115851851",
	// 		Fax:          "",
	// 		Observations: "",
	// 		PreferredContact: invoicexpress.Contact{
	// 			Name:  "Kees Zorge",
	// 			Email: "info@omniboost.io",
	// 			Phone: "0031115851851",
	// 		},
	// 	},
	// 	Items: []invoicexpress.NewItem{
	// 		invoicexpress.NewItem{
	// 			Name:        "Test",
	// 			Description: "Test Description",
	// 			UnitPrice:   10.0,
	// 			Quantity:    1.0,
	// 			Unit:        "service",
	// 			Discount:    0.0,
	// 			Tax: invoicexpress.Tax{
	// 				Name: "TEST",
	// 			},
	// 		},
	// 	},
	// 	MbReference:    "",
	// 	OwnerInvoiceID: "",
	// }

	// req.RequestBody().Invoice.Client = invoicexpress.NewClient{
	// 	Name:         "Kees Zorge",
	// 	Code:         "1",
	// 	Email:        "info@omniboost.io",
	// 	Address:      "Axelsestraat 4",
	// 	City:         "Zaamslag",
	// 	PostalCode:   "4543CJ",
	// 	Country:      "Portugal",
	// 	FiscalID:     "",
	// 	Website:      "https://www.omniboost.io",
	// 	Phone:        "0031115851851",
	// 	Fax:          "",
	// 	Observations: "",
	// 	PreferredContact: invoicexpress.Contact{
	// 		Name:  "Kees Zorge",
	// 		Email: "info@omniboost.io",
	// 		Phone: "0031115851851",
	// 	},
	// }
	// resp, err := req.Do()
	// if err != nil {
	// 	t.Error(err)
	// }
	// log.Println(resp)
}
