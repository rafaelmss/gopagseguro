package controllers

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"fmt"
	"../tools"
)

type PagSeguroTransactionRequest struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

type PagSeguroTansactionResponse struct {
	Key       string `json:"key"`
	Completed bool   `json:"completed"`
}

var token = "TOKEN"
var email = "EMAIL"

func TransactionPagSeguro(w http.ResponseWriter, r *http.Request) {

	var param PagSeguroTransactionRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := r.Body.Close(); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.Unmarshal(body, &param); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			return
		}
	}

	req := tools.NewPaymentRequest(token, email, "REFID", "http://localhost:8080/pagseguro/redirection", "http://localhost:8080/pagseguro/notification")
	req.AddItem("ID", "DESCRIPTION", 23.56, 1)
	req.SetBuyer("Nome do Comprador", "comprador@email.com").SetCPF("00000000000")
	req.SetShipping(tools.ShippingOther, 10.0).SetAddress("SP", "São Paulo", "00000000", "Bairro", "Rua Teste", "1040", "Apt 111")
	result := req.Submit()

	response := PagSeguroTansactionResponse{}
	if !result.Success {
		response.Key = ""
		response.Completed = false
	} else {
		response.Key = result.CheckoutResponse.Code
		response.Completed = result.Success
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NotificationPagSeguro(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NOTIFICATION")
}

func RedirectionPagSeguro(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REDIRECTION")
}