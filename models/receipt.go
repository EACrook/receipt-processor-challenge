package models

type StandardReceipt struct {
	ID string `json:"id"`
	Retailer string `json:"retailer" binding:"required,retailerFormat"`
	PurchaseDate string `json:"purchaseDate" binding:"required,dateFormat"`
	PurchaseTime string `json:"purchaseTime" binding:"required,timeFormat"`
	Items []ItemData `json:"items" binding:"required,dive"`
	Total string `json:"total" binding:"required,priceFormat"`
}

type ItemData struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price string `json:"price" binding:"required,priceFormat"`
}