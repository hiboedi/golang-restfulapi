package controller

import (
	"encoding/json"
	"net/http"
	"restful-api/helper"
	"restful-api/model/web"
	"restful-api/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controler *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	// panggil service
	categoryResponse := controler.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controler *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	// baca data json
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	// konversi string ke int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	// panggil service
	categoryResponse := controler.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// berithau bahwa ini json di header
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	// ubah ke json
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controler *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// konversi string ke int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// panggil service
	controler.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	// berithau bahwa ini json di header
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	// ubah ke json
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controler *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// konversi string ke int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	// panggil service
	categoryResponse := controler.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	// berithau bahwa ini json di header
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	// ubah ke json
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controler *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// panggil service
	categoryResponses := controler.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}

	// berithau bahwa ini json di header
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	// ubah ke json
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
