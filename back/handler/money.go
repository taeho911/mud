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

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
