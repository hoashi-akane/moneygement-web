package model

type SavingDetail struct {
	// ID ID
	ID int `json:"id"`
	// SavingId 貯金ID
	SavingId int `json:"savingId"`
	// SavingAmount 貯金金額
	SavingAmount int `json:"savingAmount"`
	// SavingDate 貯金日
	SavingDate string `json:"savingDate"`
	// Note メモ
	Note string `json:"note"`
}
