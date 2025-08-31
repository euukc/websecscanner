package models


type Result struct {
	Check string `json:"check"`
	Details string `json:"details"`
	Status string `json:"status"`
}