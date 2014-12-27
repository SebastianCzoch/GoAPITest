package Router

import (
    "net/http"
    "fmt"
)

func HandleRequest(rw http.ResponseWriter, request *http.Request) {
    var bundle, action string;
     
    fmt.Println(request.URL)
}
