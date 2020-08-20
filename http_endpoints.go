package products

import (
	"encoding/json"
	"net/http"
)

type HttpEndpointsFactory interface {
	ListProductsEndpoint() func(w http.ResponseWriter, r *http.Request)
	CreateProductEndpoint() func(w http.ResponseWriter, r *http.Request)
	//GetProductByIdEndpoint(idParam string) func(w http.ResponseWriter,r *http.Request)
	//UpdateProductEndpoint(ipParam string) func(w http.ResponseWriter,r *http.Request)
	//DeleteProductEndpoint(idParam string) func(w http.ResponseWriter,r *http.Request)
}

type httpEndpointsFactory struct {
	productService ProductService
}

type customError struct {
	Message string `json:"message"`
}

func NewHttpEndpoints(productService ProductService) HttpEndpointsFactory {
	return &httpEndpointsFactory{productService: productService}
}

func (httpFac *httpEndpointsFactory) ListProductsEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		listProductReq := &ListProductCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(listProductReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := listProductReq.Exec(httpFac.productService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}
func (httpFac *httpEndpointsFactory) CreateProductEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		createProductReq := &CreateProductCommand{}
		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(createProductReq)
			if err != nil {
				respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
				return
			}
		}
		data, err := createProductReq.Exec(httpFac.productService)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, &customError{err.Error()})
			return
		}
		respondJSON(w, http.StatusCreated, data)
	}
}

//GetProductByIdEndpoint(idParam string) func(w http.ResponseWriter,r *http.Request)
//UpdateProductEndpoint(ipParam string) func(w http.ResponseWriter,r *http.Request)
//DeleteProductEndpoint(idParam string) func(w http.ResponseWriter,r *http.Request)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
