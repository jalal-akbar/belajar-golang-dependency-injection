package app

import (
	"github.com/jalal-akbar/belajar-golang-dependency-injection/controller"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {

	// http Router impl
	router := httprouter.New()

	router.GET("/api/categories/", categoryController.FindAll)
	router.GET(`/api/categories/:categoryId`, categoryController.FindById)
	router.POST("/api/categories/", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// Error Handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
