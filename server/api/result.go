package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
}

func SetResult(statusCode int, result interface{}, errorInfo error, w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")

	w.Header().Set("name", "my-site")
	apiResult := ApiResult{}
	if statusCode == http.StatusOK {
		apiResult.Success = true

	} else {
		apiResult.Success = false
	}
	apiResult.Result = result
	if errorInfo != nil {
		apiResult.Error = errorInfo.Error()

	}
	jsonResponse, err := json.Marshal(apiResult)
	if err != nil {
		apiResult.Success = false
		apiResult.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		fmt.Fprintf(w, "%s", string(jsonResponse))

	}
	w.Header().Add("date", "nowwwwwwwwww")
}
