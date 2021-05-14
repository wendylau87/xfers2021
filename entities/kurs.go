package entities

type Kurs struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateKurs struct {
	Name      string          `json:"symbol"`
	ERate     CreateERate     `json:"e_rate"`
	TTCounter CreateTTCounter `json:"tt_counter"`
	BankNote  CreateBankNote  `json:"bank_notes"`
	ValidDate string          `json:"date"`
}

type PutKurs struct {
	Name      string          `json:"symbol"`
	ERate     CreateERate     `json:"e_rate"`
	TTCounter CreateTTCounter `json:"tt_counter"`
	BankNote  CreateBankNote  `json:"bank_notes"`
	ValidDate string          `json:"date"`
}

type KursResponse struct {
	Name      string            `json:"symbol"`
	ERate     ERateResponse     `json:"e_rate"`
	TTCounter TTCounterResponse `json:"tt_counter"`
	BankNote  BankNoteResponse  `json:"bank_notes"`
}

type KursResponses struct {
	Name       string              `json:"symbol"`
	ERates     []ERateResponse     `json:"e_rate"`
	TTCounters []TTCounterResponse `json:"tt_counter"`
	BankNotes  []BankNoteResponse  `json:"bank_notes"`
}
