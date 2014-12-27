package Configuration

import (
    "encoding/json"
    "os"
    "log"
)

var Core struct {
    Port    string      `json:"port"`
    BindAddress    string      `json:"bindAddress"`
}

func ReadConfiguration() {
    file, err := os.Open("./config.json")
    if err != nil {
        log.Fatal("Can't open configuration file (config.json): ", err)
    }

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&Core)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
}
