package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql" // import manual mysql
	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}

}

func main() {

	// validate := validator.New()
	// db := app.NewDB()
	// // Controller
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)
	// // http Router impl
	// router := app.NewRouter(categoryController)
	// // Http Server Impl
	// server := http.Server{
	// 	Addr:    "localhost:3001",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	server := ServerInject()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
