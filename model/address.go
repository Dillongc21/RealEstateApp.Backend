package model

type Address struct {
	ID		string	`json:"id"`
	Line1	string	`json:"line1"`
	Line2   string	`json:"line2"`
    City    string  `json:"city"`
	Zip		string	`json:"zip"`
	State	string	`json:"state"`
    Country string  `json:"country"`
}
