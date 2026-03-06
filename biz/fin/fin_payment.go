package fin

import "time"

// Payment 支付记录
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	PaymentCode   string    `gorm:"size:30;uniqueIndex" json:"payment_code"`
	BillID        uint      `json:"bill_id"`
	TenantID      uint      `json:"tenant_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod uint8     `json:"payment_method"`
	PaymentTime   time.Time `json:"payment_time"`
	TransactionID string    `gorm:"size:100" json:"transaction_id"`
	Status        uint8     `json:"status"`
	CreateTime    time.Time `json:"create_time"`
}

func (*Payment) TableName() string {
	return "hrs_payment"
}
