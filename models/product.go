package models

type Product struct {
	Id int	`json:"id"`
	Name string `json:"goods"`
	Space string `json:"space"`
	Description string `json:"description"`
	Pic1 string	`json:"pic1"`
	Pic2 string `json:"pic2"`
	Pic3 string `json:"pic3"`
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

type ProductResp struct {
	Id int	`json:"id"`
	Name string `json:"goods"`
	Space string `json:"space"`
	Description string `json:"description"`
	Pics []string	`json:"pics"`
}