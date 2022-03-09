package logic

import (
	"context"
	"go-zero-demo/mall/order/api/internal/models"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCreateLogic {
	return OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req types.OrderCreateReq) (resp *types.OrderCreateReply, err error) {
	orderModel := models.NewOrderModel(l.svcCtx)
	resp, err = orderModel.CreateOrder(&req)
	return
}
