package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/web"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

// HTTP Router Impl
func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Category Create Request
	categoryCreateRequest := web.CategoryCreateRequest{}
	// Read From Request Body
	// decoder := json.NewDecoder(request.Body)
	// err := decoder.Decode(&categoryCreateRequest)
	// helper.PanicIfError(err)
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	// Controller Response
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	// Web Response
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	// Write To Request Body
	// encode := json.NewEncoder(writer)
	// err = encode.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Category Update Request
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	// Read From Request Body
	decode := json.NewDecoder(request.Body)
	decode.Decode(&categoryUpdateRequest)
	// Set Id
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	// Controller Response
	categoryRespomse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	// Web Response
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespomse,
	}
	// Write To Request Body
	// encode := json.NewEncoder(writer)
	// err = encode.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)

}
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Set Id
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	// Controller
	controller.CategoryService.Delete(request.Context(), id)
	// Web Response
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	// Write To Rrequest Body
	// writer.Header().Add("Content-Type", "application/json")
	// encode := json.NewEncoder(writer)
	// err = encode.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Set Id
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	// Controller
	categoryRespponse := controller.CategoryService.FindById(request.Context(), id)
	// Web Response
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespponse,
	}
	// Write To Request Body
	// encode := json.NewEncoder(writer)
	// err = encode.Encode(webResponse)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Category Responses
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	// Web Rresponses
	webResponses := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	// Write To Rrequest Body
	// encoder := json.NewEncoder(writer)
	// err := encoder.Encode(webResponses)
	// helper.PanicIfError(err)
	helper.WriteToResponseBody(writer, webResponses)
}
