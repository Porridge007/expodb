package models

type Product struct {
	Id int	`json:"id"`
	Name string `json:"goods"`
	Description string `json:"description"`
}

type ProductDetail struct {
	Id int
	Name string
	Space string
	Product string
	ProductEnglish string
	Description string
	DescriptionEnglish string
	Pic1 string
	Pic2 string
	Pic3 string
}