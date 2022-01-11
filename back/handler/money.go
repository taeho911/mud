package handler

import (
	"net/http"
	"taeho/mud/agent"
	"taeho/mud/errors"
	"taeho/mud/model"
)

// ----------------------------------------------------------
// Handler Functions
// ----------------------------------------------------------
func MoneyGetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodGet {
		writeError(w, errors.INVALID_METHOD, "get method only", http.StatusBadRequest)
		return
	}
	moneyList, err := agent.MoneyFindByUsername(ctx, ctx.Value(usernameKey).(string))
	if err != nil {
		writeError(w, errors.FAILED_GET, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, moneyList, http.StatusOK)
}

func MoneyPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost {
		writeError(w, errors.INVALID_METHOD, "post method only", http.StatusBadRequest)
		return
	}
	var money model.Money
	if err := parseReqBody(r.Body, &money); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	money.Username = ctx.Value(usernameKey).(string)
	if err := agent.MoneyInsertOne(ctx, &money); err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, money, http.StatusOK)
}

func MoneyPutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPut {
		writeError(w, errors.INVALID_METHOD, "put method only", http.StatusBadRequest)
		return
	}
	var money model.Money
	if err := parseReqBody(r.Body, &money); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	if money.Username != ctx.Value(usernameKey).(string) {
		writeError(w, errors.INVALID_REQUEST_BODY, "invalid username", http.StatusBadRequest)
		return
	}
	if _, err := agent.MoneyUpdateOne(ctx, &money); err != nil {
		writeError(w, errors.FAILED_PUT, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, money, http.StatusOK)
}

func MoneyDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodDelete {
		writeError(w, errors.INVALID_METHOD, "delete method only", http.StatusBadRequest)
		return
	}
	var money model.Money
	if err := parseReqBody(r.Body, &money); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := agent.MoneyDeleteByID(ctx, money.ID); err != nil {
		writeError(w, errors.DELETE_MONEY_FAILED, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
