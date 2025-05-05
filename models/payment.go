package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID    uint      `json:"order_id"`
	Order      *Order    `gorm:"foreignKey:OrderID"`
	Method     string    `json:"method"`      // cash, qr, card, transfer
	AmountPaid float32   `json:"amount_paid"` // จำนวนเงินที่ลูกค้าจ่าย
	Change     float32   `json:"change"`      // เงินทอน
	PaidAt     time.Time `json:"paid_at"`     // เวลาที่จ่ายจริง
}
