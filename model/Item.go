package model

type Item struct {
	ItemID      int    `json:"item_id" gorm:"primaryKey"`
	ItemCode    int    `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     string `json:"order_id"`
}
