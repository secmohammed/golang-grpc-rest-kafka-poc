package tests

import (
    "bytes"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "net/http"
    "net/http/httptest"
)

func MakeRequest(method, url string, body interface{}, token string, router *gin.Engine) *httptest.ResponseRecorder {

    requestBody, _ := json.Marshal(body)
    request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
    request.Header.Add("Content-Type", "application/json")
    if token != "" {
        request.Header.Add("Authorization", "Bearer "+token)
    }
    writer := httptest.NewRecorder()
    router.ServeHTTP(writer, request)
    return writer
}
