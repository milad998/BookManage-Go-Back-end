package main

import (
    "fmt"
    "log"
    "BookManager/config"
    "BookManager/router"
    "net/http"
)

func main() {
    config.ConnectDB()
    r := router.Router()
    fmt.Println("📚 Library API running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
