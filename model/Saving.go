package model

type Saving struct {
	// ID ID
	ID int `json:"id"`
	// UserId ユーザID
	UserId int `json:"userId"`
	// TargetAmount 目標貯金額
	TargetAmount int `json:"targetAmount"`
}