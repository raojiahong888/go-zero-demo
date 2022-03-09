package payment

import (
	"context"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentCallBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentCallBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) PaymentCallBackLogic {
	return PaymentCallBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentCallBackLogic) PaymentCallBack(req types.PayCallBackRequest) (resp *types.PayCallBackResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
