package controller

import (
	"P1_API/data/request"
	"P1_API/data/response"
	"P1_API/helper"
	"P1_API/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	BookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &BookCreateRequest)

	controller.BookService.Create(requests.Context(), BookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	controller.BookService.Update(requests.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	result := controller.BookService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
