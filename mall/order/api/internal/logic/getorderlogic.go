package logic

import (
	"context"
	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
	return GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req types.OrderReq) (resp *types.OrderReply, err error) {
	return
	//user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
	//	Id: "1",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//if user.Id != "1" {
	//	return nil, errors.New("user doest exists")
	//}
	//
	//return &types.OrderReply{
	//	Id:   req.Id,
	//	Name: user.Name,
	//}, nil
}
