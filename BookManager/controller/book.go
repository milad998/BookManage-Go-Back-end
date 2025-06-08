package controller

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "BookManager/config"
    "BookManager/model"


    "go.mongodb.org/mongo-driver/bson"
    
)

var collection = config.GetCollection(config.DB, "book")

// إعداد مهلة الاتصال
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

// ✅ إنشاء كتاب جديد
func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var book model.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    res, err := collection.InsertOne(ctx, book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(res)
}

// ✅ استرجاع كل الكتب
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var books []model.Book

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var book model.Book
        if err := cursor.Decode(&book); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        books = append(books, book)
    }

    if err := cursor.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(books)
}
