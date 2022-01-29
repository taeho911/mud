package handler

import (
	"net/http"
	"strconv"
	"taeho/mud/agent"
	"taeho/mud/errors"
	"taeho/mud/model"
)

// ----------------------------------------------------------
// Handler Functions
// ----------------------------------------------------------
func MoneyGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodGet {
		writeError(w, errors.INVALID_METHOD, "get method only", http.StatusBadRequest)
		return
	}
	query := r.URL.Query()
	yearStr := query.Get("year")
	monthStr := query.Get("month")
	countStr := query.Get("count")
	var moneyList []model.Money
	var err error
	if len(yearStr) > 0 && len(monthStr) > 0 && len(countStr) > 0 {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			writeError(w, errors.INVALID_QUERY, "year is not number format", http.StatusBadRequest)
			return
		}
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			writeError(w, errors.INVALID_QUERY, "month is not number format", http.StatusBadRequest)
			return
		}
		count, err := strconv.Atoi(countStr)
		if err != nil {
			writeError(w, errors.INVALID_QUERY, "count is not number format", http.StatusBadRequest)
			return
		}
		moneyList, err = agent.MoneyFindByMonth(ctx, ctx.Value(usernameKey).(string), year, month, count)
		if err != nil {
			writeError(w, errors.FAILED_GET, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		moneyList, err = agent.MoneyFindByUsername(ctx, ctx.Value(usernameKey).(string))
		if err != nil {
			writeError(w, errors.FAILED_GET, err.Error(), http.StatusBadRequest)
			return
		}
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
	if _, err := agent.MoneyDeleteByID(ctx, money.ID, ctx.Value(usernameKey).(string)); err != nil {
		writeError(w, errors.DELETE_MONEY_FAILED, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func MoneyAutoInputGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodGet {
		writeError(w, errors.INVALID_METHOD, "get method only", http.StatusBadRequest)
		return
	}
	moneyAutoInputList, err := agent.MoneyAutoInputFindByUsername(ctx, ctx.Value(usernameKey).(string))
	if err != nil {
		writeError(w, errors.FAILED_GET, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, moneyAutoInputList, http.StatusOK)
}

func MoneyAutoInputPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPost {
		writeError(w, errors.INVALID_METHOD, "post method only", http.StatusBadRequest)
		return
	}
	var moneyAutoInput model.MoneyAutoInput
	if err := parseReqBody(r.Body, &moneyAutoInput); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	moneyAutoInput.Username = ctx.Value(usernameKey).(string)
	if err := agent.MoneyAutoInputInsertOne(ctx, &moneyAutoInput); err != nil {
		writeError(w, errors.FAILED_POST, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, moneyAutoInput, http.StatusOK)
}

func MoneyAutoInputPutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodPut {
		writeError(w, errors.INVALID_METHOD, "put method only", http.StatusBadRequest)
		return
	}
	var moneyAutoInput model.MoneyAutoInput
	if err := parseReqBody(r.Body, &moneyAutoInput); err != nil {
		writeError(w, errors.INVALID_REQUEST_BODY, err.Error(), http.StatusBadRequest)
		return
	}
	if moneyAutoInput.Username != ctx.Value(usernameKey).(string) {
		writeError(w, errors.INVALID_REQUEST_BODY, "invalid username", http.StatusBadRequest)
		return
	}
	if _, err := agent.MoneyAutoInputUpdateOne(ctx, &moneyAutoInput); err != nil {
		writeError(w, errors.FAILED_PUT, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, moneyAutoInput, http.StatusOK)
}

func MoneyAutoInputDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if r.Method != http.MethodDelete {
		writeError(w, errors.INVALID_METHOD, "delete method only", http.StatusBadRequest)
		return
	}
	if _, err := agent.MoneyAutoInputDeleteByID(ctx, r.URL.Query().Get("_id"), ctx.Value(usernameKey).(string)); err != nil {
		writeError(w, errors.DELETE_MONEY_FAILED, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ----------------------------------------------------------
// Extra Functions
// ----------------------------------------------------------
