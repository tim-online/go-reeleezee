package reeleezee

import "github.com/aodin/date"

type ManaualJournal struct {
	id                 string              `json:"id"`
	DocumentType       int                 `json:"DcoumentType"`
	Reference          string              `json:"Reference"`
	Description        string              `json:"Description"`
	TotalPayableAmount float64             `json:"TotalPayableAmount"`
	DocumentLineItems  []ManualJournalLine `json:"DocumentLineItems"`
}

type ManualJournalLine struct {
	Account      Account `json:"Account"`
	CreditAmount float64 `json:"CreditAmount"`
	DebitAmount  float64 `json:"DebitAmount"`
}

type Account struct {
	id string `json:"id"`
}
