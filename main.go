package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Request struct {
	Income float64 `json:"income"`
}

type Response struct {
	Income float64 `json:"income"`
	Tax    float64 `json:"tax"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func calculateIncomeTax(income float64) float64 {
	if income <= 11000 {
		return 0
	} else if income > 11000 && income <= 25000 {
		return ((income - 11000) * 5110) / 14000
	} else if income > 25000 && income <= 60000 {
		return 5110 + (((income - 25000) * 15125) / 35000)
	} else {
		return 20235 + (income-60000)*0.5
	}
}

type IncomeTaxHandler struct{}

func (t *IncomeTaxHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var encoder = json.NewEncoder(res)
	var requestJson = &Request{}
	var header = res.Header()
	header["Content-Type"] = []string{"application/json"}
	header["Access-Control-Allow-Origin"] = []string{"*"}

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

	// Convert to 2 decimals precision and parse as float, so it's also
	// encoded as float in JSON
	tax, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", calculateIncomeTax(income)), 64)

	encoder.Encode(&Response{Tax: tax, Income: income})
}

func main() {
	http.Handle("/income-tax", &IncomeTaxHandler{})
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "See https://github.com/CHH/at-tax for usage.")
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
