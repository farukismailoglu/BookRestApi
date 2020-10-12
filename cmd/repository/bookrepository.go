package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func connectToCollection() (*mongo.Client, *mongo.Collection) {
	clientOptions := options.Client().ApplyURI("mongodb://username:password@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("bookdb").Collection("book")
	return client, collection
}

func closeMongoDb(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection  to MongoDb closed.")
}

func InsertSingleBook(book Book) {
	client, collection := connectToCollection()
	defer closeMongoDb(client)
	insertResult, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document", insertResult.InsertedID)
}

func FindBook(isbn string) Book {
	client, collection := connectToCollection()
	defer closeMongoDb(client)
	var result Book
	filter := bson.D{{"isbn", isbn}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Found a single document %+v\n", result)
	return result
}

func UpdateBook(isbn string, book Book) {
	client, collection := connectToCollection()
	defer closeMongoDb(client)
	filter := bson.D{{"isbn", isbn}}
	updateBook := bson.M{"title": book.Title, "isbn": book.Isbn, "author": bson.M{"firstname": &book.Author.Firstname, "lastname": &book.Author.Lastname}}

	_ = collection.FindOneAndReplace(context.TODO(), filter, updateBook, nil)
}
func DeleteBook(isbn string) {
	client, collection := connectToCollection()
	defer closeMongoDb(client)
	filter := bson.D{{"isbn", isbn}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	log.Println("Deleted " + isbn)
}

func GetAllBook() []*Book {
	client, collection := connectToCollection()
	defer closeMongoDb(client)
	findOptions := options.Find()
	findOptions.SetLimit(2)
	var result []*Book
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		log.Println(err)
	}
	for cur.Next(context.TODO()) {
		var elem Book
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
		}

		result = append(result, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	return result
}
