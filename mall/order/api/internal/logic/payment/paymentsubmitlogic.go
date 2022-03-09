package payment

import (
	"context"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) PaymentSubmitLogic {
	return PaymentSubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentSubmitLogic) PaymentSubmit(req types.PayRequest) (resp *types.PayResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
