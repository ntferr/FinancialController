package models

type Card struct {
	Code   string `json:"code"`
	Bank   string `json:"bank"`
	Number string `json:"number"`
	Credit bool   `json:"credit"`
}
