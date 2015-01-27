package Router

import (
    "net/http"
    "strings"
    "fmt"
)

func HandleRequest(rw http.ResponseWriter, request *http.Request) {
    fmt.Println(request.URL)
    h := strings.Split(request.URL.Path, "/")
    fmt.Println(h[1])
    fmt.Println(h[2])
   // fmt.Println(h[3])
}
