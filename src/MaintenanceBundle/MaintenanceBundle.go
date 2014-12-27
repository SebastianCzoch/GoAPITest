package MaintenanceBundle

import (
    "net/http"
    "encoding/json"
    "time"
)

var Uptime int64

func SetUptime(uptime int64) {
    Uptime = uptime
}

func ServerStatus(rw http.ResponseWriter, request *http.Request) {
    type ServerStatusResponse struct {
        Uptime      int64
    }

    interval := int64(time.Now().Unix() - Uptime)
    response := ServerStatusResponse{interval}

    js, err := json.Marshal(response)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.Write(js)
}
