package main

import (
    "fmt"
    "log"
    "BookManager/config"
    "BookManager/router"
    "net/http"
)

func main() {
    config.ConnectDB() // هذا يحفظ الاتصال في config.DB
    r := router.Router() // تأكد أن Router() تستخدم config.DB داخليًا
    fmt.Println("📚 Library API running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
