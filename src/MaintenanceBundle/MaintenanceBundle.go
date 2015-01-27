package MaintenanceBundle

import (
    "time"
)

type MaintenanceResponse struct {
    Uptime int64
}

type Maintenance struct {
    StartTime int64
}

func Init() Maintenance {
    return Maintenance{time.Now().Unix()}
}

func (m *Maintenance) GetResponse() MaintenanceResponse {
    interval := int64(time.Now().Unix() - m.StartTime)
    return MaintenanceResponse{interval}
}
