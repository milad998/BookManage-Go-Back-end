package config

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "time"
)

var DB *mongo.Client

func ConnectDB() *mongo.Client {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://miladnasir2023:5F1gMW7nj0ywRbAy@cluster0.typokzf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client
    return DB
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
    return client.Database("librarydb").Collection(collectionName)
}
