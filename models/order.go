package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID  uint        `json:"user_id"` // คนที่ขาย
	User    Users       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Total   float32     `json:"total"`
	IsPaid  bool        `json:"paid"`
	Payment *Payment    `gorm:"constraint:OnDelete:SET NULL"`
	Items   []OrderItem `gorm:"constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint               `json:"order_id"`
	ProductID uint               `json:"product_id"`
	Product   Product            `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  uint               `json:"quantity"`
	Price     float32            `json:"price"` // ราคา ณ เวลาขาย (ป้องกันราคาถูกเปลี่ยนภายหลัง)
	Toppings  []OrderItemTopping `gorm:"constraint:OnDelete:CASCADE"`
}

type OrderItemTopping struct {
	gorm.Model
	OrderItemID uint      `json:"order_item_id"`
	OrderItem   OrderItem `gorm:"foreignKey:OrderItemID;references:ID"`
	ToppingID   uint      `json:"topping_id"`
	Topping     Topping   `gorm:"foreignKey:ToppingID;references:ID"`
}

// Response structs
type OrderResponse struct {
	ID        uint                `json:"id"`
	CreatedAt string              `json:"created_at"`
	UserID    uint                `json:"user_id"`
	User      UserBrief           `json:"user"`
	Total     float32             `json:"total"`
	IsPaid    bool                `json:"paid"`
	Payment   *PaymentBrief       `json:"payment,omitempty"`
	Items     []OrderItemResponse `json:"items"`
}

type UserBrief struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}

type PaymentBrief struct {
	Method     string  `json:"method"`
	AmountPaid float32 `json:"amount_paid"`
	Change     float32 `json:"change"`
	PaidAt     string  `json:"paid_at"`
}

type OrderItemResponse struct {
	ID        uint           `json:"id"`
	ProductID uint           `json:"product_id"`
	Product   ProductBrief   `json:"product"`
	Quantity  uint           `json:"quantity"`
	Price     float32        `json:"price"`
	Toppings  []ToppingBrief `json:"toppings"`
}

type ProductBrief struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       float32       `json:"price"`
	Category    CategoryBrief `json:"category"`
}

type CategoryBrief struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ToppingBrief struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
