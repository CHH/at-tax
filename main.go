package main

import (
	"encoding/json"
	"net/http"
	"os"
	"log"
)

type Request struct {
	Income float64 `json:"income"`
}

type Response struct {
	Tax float64 `json:"tax"`
}

type ErrorResponse struct {
	Reason string
}

func calculateIncomeTax(income float64) float64 {
	if income <= 11000 {
		return 0
	} else if income > 11000 && income <= 25000 {
		return ((income - 11000) * 5110) / 14000
	} else if income > 25000 && income <= 60000 {
		return 5110 + (((income - 25000) * 15125) / 35000)
	} else {
		return 20235 + (income - 60000) * 0.5
	}
}

type IncomeTaxHandler struct {}

func (t *IncomeTaxHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var encoder = json.NewEncoder(res)
	var requestJson = &Request{}
	var header = res.Header()
	header["Content-Type"] = []string{"application/json"}

	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		encoder.Encode(&ErrorResponse{Reason: "Wrong method"})
		return
	}

	if err := decoder.Decode(requestJson); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		encoder.Encode(&ErrorResponse{Reason: "Could not parse JSON: " + err.Error()})
		return
	}

	income := requestJson.Income
	encoder.Encode(&Response{Tax: calculateIncomeTax(income)})
}

func main() {
	http.Handle("/income-tax", &IncomeTaxHandler{})

	if err := http.ListenAndServe(":" + os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
