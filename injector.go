//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/app"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/controller"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/middleware"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/repository"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/service"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func ServerInject() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
