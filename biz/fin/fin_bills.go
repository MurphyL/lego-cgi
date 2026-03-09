package fin

import "time"

// Bill 账单信息
type Bill struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BillCode    string    `gorm:"size:30;uniqueIndex" json:"bill_code"`
	ContractID  uint      `json:"contract_id"`
	TenantID    uint      `json:"tenant_id"`
	PropertyID  uint      `json:"property_id"`
	BillType    uint8     `json:"bill_type"`
	Amount      float64   `json:"amount"`
	DueDate     time.Time `json:"due_date"`
	Status      uint8     `json:"status"`
	PaidAmount  float64   `json:"paid_amount"`
	Description string    `gorm:"size:255" json:"description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

func (*Bill) TableName() string {
	return "hrs_bill"
}
