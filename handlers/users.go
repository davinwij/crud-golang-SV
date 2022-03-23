package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"simple-crud-task/connection"
	"simple-crud-task/models"

	_ "github.com/gorilla/mux"
)

var (
	stockP float32
	bondP  float32
	mmP    float32
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var userInput models.USER
	json.Unmarshal(payloads, &userInput)

	countRisk(userInput.Age)

	riskPayloads := []models.RISK_PROFILE{{BondPercent: bondP, MMPercent: mmP, StockPercent: stockP}}
	fmt.Println(riskPayloads)

	connection.DB.Create(&userInput)
	connection.DB.Create(&riskPayloads)

	res := models.Result{Code: 200, Data: userInput, Message: "Success create user"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content Type", "applicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func countRisk(age int) {
	var base = 55 - age

	if base >= 30 {
		stockP = 72.5
		bondP = 21.5
		mmP = 100 - stockP + bondP
		return
	} else if base >= 20 {
		stockP = 54.5
		bondP = 25.5
		mmP = 100 - stockP + bondP
		return
	} else if base < 20 {
		stockP = 34.5
		bondP = 45.5
		mmP = 100 - stockP + bondP
		return
	}
}
