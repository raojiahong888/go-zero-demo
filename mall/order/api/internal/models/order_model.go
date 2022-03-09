package models

import (
	"fmt"
	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"
)

type IOrder interface {
	GetList(params *types.OrderHistoryRequest) (types.OrderHistoryRes, error)
	CreateOrder(params *types.OrderCreateReq) (res *types.OrderCreateReply, err error)
}

type orderModel struct {
	svcCtx *svc.ServiceContext
}

func NewOrderModel(svcCtx *svc.ServiceContext) IOrder {
	return &orderModel{
		svcCtx: svcCtx,
	}
}

func (m *orderModel) GetList(params *types.OrderHistoryRequest) (res types.OrderHistoryRes, err error) {
	orderModel := new(types.Order)
	query := m.svcCtx.Db.Model(orderModel)
	if params.OrderSn != "" {
		query = query.Where("order_sn=?", params.OrderSn)
	}
	query.Count(&res.Total)
	query.Scan(&res.List)
	fmt.Println(res)
	return
}

func (m *orderModel) CreateOrder(params *types.OrderCreateReq) (res *types.OrderCreateReply, err error) {
	orderModel := &types.Order{
		OrderSn:     params.OrderSn,
		PayStatus:   1,
		TotalAmount: params.TotalAmount,
		PayAmount:   params.PayAmount,
	}
	err = m.svcCtx.Db.Create(orderModel).Error
	if err != nil {
		return
	}
	res = &types.OrderCreateReply{
		Id:        orderModel.Id,
		OrderSn:   orderModel.OrderSn,
		PayAmount: orderModel.PayAmount,
	}
	return
}