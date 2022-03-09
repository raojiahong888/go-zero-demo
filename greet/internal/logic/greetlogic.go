package logic

import (
	"context"
	"fmt"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = req.Name
	fmt.Println(l.svcCtx.Config.Name)
	fmt.Println(l.svcCtx.Config.Host)
	fmt.Println(l.svcCtx.Config.Port)
	logx.Info(l.svcCtx.Config)
	return
}
