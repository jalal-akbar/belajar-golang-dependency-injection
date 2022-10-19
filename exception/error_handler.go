package exception

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/helper"
	"github.com/jalal-akbar/belajar-golang-dependency-injection/model/web"
)

// Error Handler
func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// Not Found
	if notFoundError(writer, request, err) {
		return
	}
	// validation error
	if validationError(writer, request, err) {
		return
	}
	// Internal Server Error
	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	// conversion err interface {} to NotFoundError
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application-json")
		writer.WriteHeader(http.StatusNotFound)
		// Web Response
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	// conversion err interface{} to validator Error
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		// Web Response
		wenResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteToResponseBody(writer, wenResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	// Web Response
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
