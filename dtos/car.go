package dtos

type Car struct {
	ID       int    `json:"id"`
	Brand    string `json:"brand"`
	Name     string `json:"name"`
	Model    string `json:"model"`
	SubModel string `json:"subModel"`
	Color    string `json:"color"`
	Price    int    `json:"price"`
}
