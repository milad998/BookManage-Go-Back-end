package controller

import (
    "context"
    "encoding/json"
    "BookManager/config"
    "BookManager/model"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = config.GetCollection(config.DB,"book")
var ctx, _= context.WithTimeout(context.Background(),10*time.Secound())

func createBook(w http.ResponseWriter,r *http.Request){
  var book  model.Book
  json.NewDecoder(r.Body).Decode(&book)
  res,_:=collection.InsertOne(ctx,book)
  json.Encoder(w).Encode(res)
}

func getBooks(w http.ResponseWriter, r *http.Request){
  var books []model.Book
  courser,_:= collection.Find(ctx,bson.M{})
  defer courser.Close()
  for courser.Next(ctx){
    var book model.Book
    courser.Decode(&book)
    books= append(books,book)
    
  }
  json.Encoder(w).Encode(books)
  
}
