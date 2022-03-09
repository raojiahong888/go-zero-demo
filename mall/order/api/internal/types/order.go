package types

import "time"

type Order struct {
	Id             uint64     `json:"id"`
	OrderSn        string     `json:"orderSn" gorm:"column:order_sn;default:''"`
	PayStatus      uint64     `json:"payStatus" gorm:"column:pay_status;default:1"`
	TotalAmount    uint64     `json:"totalAmount" gorm:"column:total_amount;default:0"`
	DiscountAmount uint64     `json:"discountAmount" gorm:"column:discount_amount;default:0"`
	PayAmount      uint64     `json:"payAmount" gorm:"column:pay_amount;default:0"`
	Receipt        string     `json:"receipt" gorm:"column:receipt;default:''"`
	PayMethod      string     `json:"payMethod" gorm:"column:pay_method;default:''"`
	PayDate        string     `json:"payDate" gorm:"column:pay_date"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

func (*Order) TableName() string {
	return `order`
}
