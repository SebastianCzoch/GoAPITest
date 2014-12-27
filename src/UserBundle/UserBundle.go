package UserBundle

import (
    "net/http"
    "encoding/json"
)

type AuthorizeResponse struct {
  Success       bool
  TokenLong     string
  TokenShort    string
}

func Authorize(rw http.ResponseWriter, request *http.Request) {
    response := AuthorizeResponse{true,"asdasd","asdasdsad"}

    js, err := json.Marshal(response)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.Write(js)
}
