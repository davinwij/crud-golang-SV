package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"simple-crud-task/connection"
	"simple-crud-task/models"

	"github.com/gorilla/mux"
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
	var userRisk models.RISK_PROFILE
	json.Unmarshal(payloads, &userInput)
	countRisk(userInput.Age)
	connection.DB.Create(&userInput)

	risk := models.RISK_PROFILE{mmP, bondP, stockP, userInput.ID}
	userRisk = risk

	connection.DB.Create(&userRisk)

	res := models.Result{Code: 200, Data: &userInput, Message: "Success create user"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content Type", "applicaiton/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	take := r.URL.Query().Get("take")
	limit, err := strconv.Atoi(take)
	offset, err := strconv.Atoi(page)

	offsets := (limit * offset) - limit

	if len(page) == 0 || len(take) == 0 {
		limit = 10
		offsets = 0
	}

	userDetail := []models.USER{}

	connection.DB.
		Limit(limit).
		Offset(offsets).
		Find(&userDetail)

	res := models.Result{Code: 200, Data: &userDetail, Message: "Success get users, page = " + page + " take = " + take}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user models.USER
	var riskProfile models.RISK_PROFILE

	connection.DB.First(&user, userId)
	connection.DB.First(&riskProfile, "user_id = ?", userId)

	userDetail := models.USER_DETAIL{user.ID, user.Name, user.Age, riskProfile.MMPercent, riskProfile.BondPercent, riskProfile.StockPercent}

	res := models.Result{Code: 200, Data: &userDetail, Message: "Success get user by id: " + userId}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func countRisk(age int) {
	var base = 55 - age

	if base >= 30 {
		stockP = 72.5
		bondP = 21.5
		mmP = 100 - (stockP + bondP)
		return
	} else if base >= 20 {
		stockP = 54.5
		bondP = 25.5
		mmP = 100 - (stockP + bondP)
		return
	} else if base < 20 {
		stockP = 34.5
		bondP = 45.5
		mmP = 100 - (stockP + bondP)
		return
	}
}
