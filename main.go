package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type User struct {
	Name  string
	Email string
}

func main() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	ctx := context.Background()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB-CONN")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	collections := client.Database("sample_mflix").Collection("users")

	//user := User{
	//	Name:  "Cao Van Hoang",
	//	Email: "caovanhoang@gmail.com",
	//}

	//result, err := collections.InsertOne(ctx, user)
	//if err != nil {
	//	log.Println(err)
	//}
	//id := result.InsertedID
	//fmt.Println(id)

	filter := bson.D{{"name", "Cao Van Hoang"}}
	result, err := collections.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result.DeletedCount)
	//var limit int64
	//limit = 5
	//cur, err := collections.Find(context.Background(), bson.D{}, &options.FindOptions{Limit: &limit})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cur.Close(ctx)
	//for cur.Next(ctx) {
	//	var result bson.D
	//	err := cur.Decode(&result)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	doc, err := bson.Marshal(result)
	//	var test User
	//	err = bson.Unmarshal(doc, &test)
	//	fmt.Println(test.Name)
	//	// do something with result....
	//}
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}

}
