// Code generated by goctl. DO NOT EDIT.
package types

type OrderReq struct {
	Id string `path:"id"`
}

type OrderReply struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type OrderDetailReq struct {
	OrderSn string `path:"orderSn"`
}

type OrderDetailReply struct {
	OrderSn   string `json:"orderSn"`
	PayStatus uint64 `json:"payStatus"`
	PayAmount uint64 `json:"payAmount"`
	Receipt   string `json:"receipt"`
	PayMethod string `json:"payMethod"`
}

type OrderCreateReq struct {
	OrderSn     string `json:"orderSn"`
	TotalAmount uint64 `json:"totalAmount"`
	PayAmount   uint64 `json:"payAmount"`
}

type OrderCreateReply struct {
	Id        uint64 `json:"id"`
	OrderSn   string `json:"orderSn"`
	PayAmount uint64 `json:"payAmount"`
}

type PayRequest struct {
	OrderSn string `json:"orderSn"`
}

type PayResponse struct {
	OrderSn   string `json:"orderSn"`
	Receipt   string `json:"receipt"`
	PayStatus uint64 `json:"payStatus"`
}

type PayCallBackRequest struct {
	OrderSn   string `json:"orderSn"`
	PayStatus uint64 `json:"payStatus"`
	PayAmount uint64 `json:"payAmount"`
	Receipt   string `json:"receipt"`
	PayMethod string `json:"payMethod"`
}

type PayCallBackResponse struct {
	OrderSn   string `json:"orderSn"`
	PayStatus uint64 `json:"payStatus"`
	PayAmount uint64 `json:"payAmount"`
	Receipt   string `json:"receipt"`
	PayMethod string `json:"payMethod"`
	PayDate   string `json:"payDate"`
}

type OrderHistoryRequest struct {
	OrderSn string `json:"orderSn,optional"`
}

type OrderHistoryRes struct {
	List  []OrderDetail `json:"list"`
	Total uint64        `json:"total"`
}

type OrderDetail struct {
	OrderSn   string `json:"orderSn"`
	PayStatus uint64 `json:"payStatus"`
	PayAmount uint64 `json:"payAmount"`
	Receipt   string `json:"receipt"`
	PayMethod string `json:"payMethod"`
}
