package dto

type WelletDtoInput struct {
	WelletID     string `json:"wellet_id"`
	UserID       string `json:"user_id"`
	Balance      int    `json:"balance"`
	DebitAmount  int    `json:"debit_amount"`
	CreditAmount int    `json:"credit_amount"`
	Debit        int    `json:"debit"`
	Credit       int    `json:"credit"`
	Tarife       int    `json:"tarife"`
	Iva          int    `json:"iva"`
	Cancel       bool   `json:"cancel"`
}

type WelletDtoOutput struct {
	WelletID     string `json:"wellet_id"`
	UserID       string `json:"user_id"`
	Balance      int    `json:"balance"`
	DebitAmount  int    `json:"debit_amount"`
	CreditAmount int    `json:"credit_amount"`
	Debit        int    `json:"debit"`
	Credit       int    `json:"credit"`
	Tarife       int    `json:"tarife"`
	Iva          int    `json:"iva"`
	Cancel       bool   `json:"cancel"`
}
