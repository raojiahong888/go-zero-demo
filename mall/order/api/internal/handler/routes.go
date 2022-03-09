// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	history "go-zero-demo/mall/order/api/internal/handler/history"
	payment "go-zero-demo/mall/order/api/internal/handler/payment"
	"go-zero-demo/mall/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/order/create",
					Handler: orderCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/order/detail/:orderSn",
					Handler: orderDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/order/get/:id",
					Handler: getOrderHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/payment/call-back",
					Handler: payment.PaymentCallBackHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/payment/submit",
					Handler: payment.PaymentSubmitHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/order/history/list",
					Handler: history.OrderHistoryListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api"),
	)
}