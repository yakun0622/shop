package models

//PayLoad jwt payload
type PayLoad struct {
	Id      int    `json:"id"`
	StoreID int    `json:"storeid"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	NBF     int    `json:"nbf"`
	EXP     int    `json:"exp"`
}
