package models


type ProductListResp struct {
	Code int	`json:"error_code"`
	Data 	[]ProductResp
}

type ProductDetailResp struct {
	Code int	`json:"erorr_code"`
	Data ProductDetail
}

type ExpoListResp struct {
	Code int	`json:"error_code"`
	Data []Expo
}

type ExpoDetailResp struct {
	Code int	`json:"error_code"`
	Data ExpoDetail
}