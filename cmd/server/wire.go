//go:build wireinject
// +build wireinject

package server

import (
	"github.com/cuiyuanxin/airuisi-admin/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ServerSet = wire.NewSet(router.NewRouter)

//var RepositorySet = wire.NewSet(
//	repository.NewDb,
//	repository.NewRepository,
//	repository.NewUserRepository,
//)
//
//var ServiceSet = wire.NewSet(
//	service.NewService,
//	service.NewUserService,
//)
//
//var HandlerSet = wire.NewSet(
//	handler.NewHandler,
//	handler.NewUserHandler,
//)

func newApp(*zap.Logger) (*gin.Engine, error) {
	panic(wire.Build(
		ServerSet,
		//RepositorySet,
		//ServiceSet,
		//HandlerSet,
	))
}
