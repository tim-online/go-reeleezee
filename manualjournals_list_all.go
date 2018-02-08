package reeleezee

import (
	"net/http"
	"net/url"

	null "gopkg.in/guregu/null.v3"
)

func (s *InvoicesService) NewListAllRequest() InvoicesListAllRequest {
	return InvoicesListAllRequest{
		api:    s.api,
		method: http.MethodGet,
		urlParams: InvoicesListAllURLParams{
			DocumentType: Invoice,
		},
		queryParams: InvoicesListAllQueryParams{
			Text:   "",
			Types:  []InvoiceDocumentType{},
			Status: []InvoiceStatus{},
		},
	}
}

type InvoicesListAllRequest struct {
	api *API
	// queryParams InvoicesListAllQueryParams
	// pathParams  InvoicesListAllPathParams
	method      string
	urlParams   InvoicesListAllURLParams
	queryParams InvoicesListAllQueryParams
}

func (r *InvoicesListAllRequest) Method() string {
	return r.method
}

func (r *InvoicesListAllRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesListAllRequest) URL() url.URL {
	// path := fmt.Sprintf("%s.json", r.params.documentType.Path)
	// path := r.params.documentType.Path
	path := "invoices.json"
	return r.api.GetEndpointURL(path)
}

func (r *InvoicesListAllRequest) Do() (InvoicesListAllResponseBody, error) {
	// Create http request
	req, err := r.api.NewRequest(nil, r.Method(), r.URL(), nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.queryParams, req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.api.Do(req, responseBody)
	return *responseBody, err
}

func (r *InvoicesListAllRequest) NewResponseBody() *InvoicesListAllResponseBody {
	return &InvoicesListAllResponseBody{}
}

func (r *InvoicesListAllRequest) QueryParams() *InvoicesListAllQueryParams {
	return &r.queryParams
}

type InvoicesListAllURLParams struct {
	Text         string
	DocumentType InvoiceDocumentType
}

// func (p *InvoicesListAllQueryParams) SetDocumentType(documentType InvoiceDocumentType) {
// 	p.types = []InvoiceDocumentType{documentType}
// }

// func (p *InvoicesListAllQueryParams) AddDocumentType(documentType InvoiceDocumentType) {
// 	p.types = append(p.types, documentType)
// }

type InvoicesListAllQueryParams struct {
	Text                 string                `schema:"text,omitempty"`
	Types                []InvoiceDocumentType `schema:"type[],omitempty"`
	Status               []InvoiceStatus       `schema:"status[],omitempty"`
	DateFrom             Date                  `schema:"date[from],omitempty"`
	DateTo               Date                  `schema:"date[to],omitempty"`
	DueDateFrom          Date                  `schema:"due_date[from],omitempty"`
	DueDateTo            Date                  `schema:"due_date[to],omitempty"`
	TotalBeforeTaxesFrom null.Float            `schema:"total_before_taxes[from],omitempty"`
	TotalBeforeTaxesTo   null.Float            `schema:"total_before_taxes[to],omitempty"`
	NonArchived          null.Bool             `schema:"non_archived,omitempty"`
	Page                 int                   `schema:"page,omitempty"`
	PerPage              int                   `schema:"per_page,omitempty"`
}

func (p InvoicesListAllQueryParams) ToURLValues() url.Values {
	params := url.Values{}
	encoder := newSchemaEncoder()
	encoder.Encode(p, params)

	if params.Get("total_before_taxes[from]") == "0.000000" {
		params.Del("total_before_taxes[from]")
	}
	if params.Get("total_before_taxes[to]") == "0.000000" {
		params.Del("total_before_taxes[to]")
	}

	return params
}

type InvoicesListAllResponseBody struct{}
