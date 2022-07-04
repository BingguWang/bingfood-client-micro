// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/biz"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/conf"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/data"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/server"
    "github.com/go-kratos/bingfood-client-micro/app/order/service/internal/service"
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, jwt *conf.JWT, logger log.Logger, registry *conf.Registry) (*kratos.App, func(), error) {
    client, err := data.NewRedis(confData)
    if err != nil {
        return nil, nil, err
    }
    db := data.NewDB(confData)
    dataData, cleanup, err := data.NewData(confData, client, logger, db)
    if err != nil {
        return nil, nil, err
    }
    orderRepo := data.NewOrderRepo(dataData, logger)
    discovery := biz.NewDiscovery(registry)
    cartServiceClient := biz.NewCartServiceClient(discovery)
    ordercase := biz.NewOrdercase(orderRepo, cartServiceClient, logger)
    orderServiceImpl := service.NewOrderService(ordercase)
    httpServer := server.NewHTTPServer(confServer, jwt, orderServiceImpl, logger)
    grpcServer := server.NewGRPCServer(confServer, orderServiceImpl, logger)
    registrar := server.NewRegistrar(registry)
    app := newApp(logger, httpServer, grpcServer, registrar)
    return app, func() {
        cleanup()
    }, nil
}
