package history

import (
	"context"
	"go-zero-demo/mall/order/api/internal/models"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderHistoryListLogic {
	return OrderHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderHistoryListLogic) OrderHistoryList(req types.OrderHistoryRequest) (resp types.OrderHistoryRes, err error) {
	//resp = &types.OrderHistoryRes{
	//	Total: 10,
	//}
	//orderDetail := types.OrderDetail{
	//	OrderSn:   req.OrderSn,
	//	PayStatus: 2,
	//	PayAmount: 10,
	//	Receipt:   "KP202202231108",
	//	PayMethod: "Ewallet",
	//}
	//list := make([]types.OrderDetail,0,resp.Total)
	//list = append(list, orderDetail)
	//resp.List = list
	orderModel := models.NewOrderModel(l.svcCtx)
	resp, err = orderModel.GetList(&req)
	return
}
