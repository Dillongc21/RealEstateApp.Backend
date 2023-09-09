package model

type Address struct {
	ID		string	`json:"id"`
	Line1		string	`json:"line1"`
	Line2		string	`json:"line2"`
	StateID		string	`json:"stateId"`
	Zip		string	`json:"zip"`
}
