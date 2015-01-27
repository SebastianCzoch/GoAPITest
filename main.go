package main

import (
    "net/http"
    "./vendors/Configuration"
    "./vendors/Router"
    "./src/MaintenanceBundle"
    "encoding/json"
    "fmt"
)

var M MaintenanceBundle.Maintenance

func main() {
    Configuration.ReadConfiguration()
    M = MaintenanceBundle.Init()
    http.HandleFunc("/status", Status)
    http.HandleFunc("/", Router.HandleRequest)

    fmt.Println("Server started")
    http.ListenAndServe(Configuration.Core.BindAddress+":"+Configuration.Core.Port, nil)
    fmt.Println("Server stoped")
}

func error404(rw http.ResponseWriter, request *http.Request) {
    rw.WriteHeader(http.StatusNotFound)
}

func Status(rw http.ResponseWriter, request *http.Request) {
    response := M.GetResponse()

    js, err := json.Marshal(response)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.Write(js)
}
