// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/biz"
	"github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/conf"
	"github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/data"
	"github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/server"
	"github.com/BingguWang/bingfood-client-micro/app/nsqSrv/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, jwt *conf.JWT, logger log.Logger, nsq *conf.NSQ, registry *conf.Registry) (*kratos.App, func(), error) {
	producer, err := data.NewNsqProducer(nsq)
	if err != nil {
		return nil, nil, err
	}
	discovery := biz.NewDiscovery(registry)
	orderServiceClient := biz.NewOrderServiceClient(discovery, jwt)
	nsqCase := biz.NewNsqCase(producer, orderServiceClient, logger)
	nsqServiceImpl := service.NewNsqService(nsqCase)
	grpcServer := server.NewGRPCServer(confServer, jwt, nsqServiceImpl, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
	}, nil
}
