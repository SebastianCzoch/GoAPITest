package main

import (
    "net/http"
    // "time"
    "./vendors/Configuration"
    "./vendors/Router"
    "fmt"
)

func main() {
    Configuration.ReadConfiguration()
    //MaintenanceBundle.SetUptime(time.Now().Unix())

    http.HandleFunc("/", Router.HandleRequest)

    fmt.Println("Server started")
    http.ListenAndServe(Configuration.Core.BindAddress+":"+Configuration.Core.Port, nil)
    fmt.Println("Server stoped")
}

func error404(rw http.ResponseWriter, request *http.Request) {
    rw.WriteHeader(http.StatusNotFound)
}
