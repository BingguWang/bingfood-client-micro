//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/biz"
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/data"
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/server"
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/service"
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.JWT, log.Logger, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
