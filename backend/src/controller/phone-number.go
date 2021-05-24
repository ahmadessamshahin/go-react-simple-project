package controller

import (
	"backend/src/db"
	h "backend/src/handler"
	m "backend/src/model"
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var (
	database      = db.GetDBInstance()
	codeToCountry = map[string]string{
		"237": "Cameroon",
		"251": "Ethiopia",
		"212": "Morocco",
		"258": "Mozambique",
		"256": "Uganda",
	}
	codeToExpression = map[string]string{
		"237": "\\(237\\)\\ ?[2368]\\d{7,8}$",
		"251": "\\(251\\)\\ ?[1-59]\\d{8}$",
		"212": "\\(212\\)\\ ?[5-9]\\d{8}$",
		"258": "\\(258\\)\\ ?[28]\\d{7,8}$",
		"256": "\\(256\\)\\ ?\\d{9}$",
	}
)

type PhoneNumberPostBody struct {
	Number string `json:"number"`
}

type PhoneNumbersRes struct {
	Total int             `json:"total"`
	Data  []m.PhoneNumber `json:"data"`
}

func GetPhoneNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	p := m.PhoneNumber{ID: id}

	if err := p.GetPhoneNumber(database.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			h.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	h.RespondWithJSON(w, http.StatusOK, p)
}

func GetPhoneNumbers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	// base case for filtering
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	data, errData := m.GetPhoneNumbers(database.DB, start, count)
	total, errCount := m.GetPhoneNumbersCounts(database.DB)
	if errData != nil && errCount != nil {
		h.RespondWithError(w, http.StatusInternalServerError, errData.Error())
		return
	}
	resContent := PhoneNumbersRes{
		Data:  data,
		Total: total,
	}

	h.RespondWithJSON(w, http.StatusOK, resContent)
}

func CreatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var b PhoneNumberPostBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	code := getCode(b.Number)
	if code == "" {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid Number Format")
	}

	if country, ok := codeToCountry[code]; ok {
		p := &m.PhoneNumber{
			Country: country,
			Code:    code,
			Number:  b.Number,
		}

		if match, _ := regexp.MatchString(codeToExpression[code], b.Number); match {
			p.State = "Valid"
			if err := p.CreatePhoneNumber(database.DB); err != nil {
				h.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			h.RespondWithJSON(w, http.StatusCreated, p)
			return
		}
		p.State = "InValid"
		if err := p.CreatePhoneNumber(database.DB); err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		h.RespondWithJSON(w, http.StatusCreated, p)
		return
	}
	h.RespondWithError(w, http.StatusBadRequest, "This Code currently Not Supported")
}

func getCode(number string) (val string) {
	var (
		idx1 = strings.Index(number, "(")
		idx2 = strings.Index(number, ")")
	)
	if idx1 != -1 && idx2 != -1 {
		val = number[idx1+1 : idx2]
	}
	return
}

func DeletePhoneNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := m.PhoneNumber{ID: id}
	if err := p.DeletePhoneNumber(database.DB); err != nil {
		h.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
