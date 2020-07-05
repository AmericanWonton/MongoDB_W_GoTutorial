package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

/*Below are two example types I'm going to use in
a seperate Citizen Program.
*/
type address struct {
	Zipcode    int    `json:"Zipcode"`
	Streetname string `json:"Streetname"`
	Province   string `json:"Province"`
	Country    string `json:"Country"`
}
type citizen struct {
	Firstname      string  `json:"Firstname"`
	Lastname       string  `json:"Lastname"`
	Ethnicity      string  `json:"Ethnicity"`
	Skincolor      string  `json:"Skincolor"`
	Age            int     `json:"Age"`
	SS             int     `json:"socialSecurity"`
	Origincountry  string  `json:"Origincountry"`
	Sex            byte    `json:"Sex"`
	Gender         string  `json:"Gender"`
	Citizennumber  int     `json:"Citizennumber"`
	Employedhere   bool    `json:"EmployedHere"`
	Employeenum    int     `json:"Employeenum"`
	Currentaddress address `json:"Currentaddress"`
}

type icecream struct {
	Flavor   string  `json:"Flavor"`
	Calories float64 `json:"Calories"`
	Name     string  `json:"Name"`
	Alias    string  `json:"Alias"`
}

func main() {
	//Setup Mongo connection to Atlas Cluster
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://joek:superduperPWord@superdbcluster.kswud.mongodb.net/superdbtest1?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) //Disconnect in 10 seconds if you can't connect
	//Double check to see if we've connected to the database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	//List all available databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	//Insert 1 data
	insertData(client)
	//Insert ALOT of data
	insertMultipleData(client)
	//Update Document
	dataUpdateAll(client)
	//Get Document
	retrieveData(client)
	//Get Multiple Document
	retrieveMultipleData(client)
	//Delete Multiple Documents
	deleteDocs(client)
}

func insertData(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams")      //Here's our collection
	testInsertion := icecream{"test flavor", 800, "The Test Cone", "Tester Cone"} //Test Document insertion
	//Here's how to insert a single document into our DB
	insertResult, err := ic_collection.InsertOne(context.TODO(), testInsertion)
	if err != nil {
		log.Fatal(err)
		println(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func insertMultipleData(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams")               //Here's our collection
	testInsertion := icecream{"test flavor the one", 800, "The Test Cone", "Tester Cone"}  //Test Document insertion 1
	testInsertion2 := icecream{"test flavor the two", 800, "The Test Cone", "Tester Cone"} //Test Document insertion 1
	theIceCream := []interface{}{testInsertion, testInsertion2}                            //The Interface of Data
	//Insert Our Data
	insertManyResult, err := ic_collection.InsertMany(context.TODO(), theIceCream)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs) //Data insert results
}

func dataUpdateAll(client *mongo.Client) {
	//Here's how to update a document with a filter 'BSON' Json object
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection

	filter := bson.D{{"flavor", "test flavor the one"}} //Here's our filter to look for

	update := bson.D{ //Here is our data to update
		{"$set", bson.D{
			{"flavor", "raw sewage"},
		}},
	}

	updateResult, err := ic_collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	//Our new UpdateResult
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func retrieveData(client *mongo.Client) {
	//Here's how to find a single document and apply it to a struct.
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	var result icecream
	filter := bson.D{{"flavor", "raw sewage"}} //Here's our filter to look for

	err := ic_collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}

func retrieveMultipleData(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	filter := bson.D{{"flavor", "raw sewage"}}                               //Here's our filter to look for
	//Here's how to find and assign multiple Documents using a cursor
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []icecream
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := ic_collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem icecream
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents: %+v\n", results)
}

func deleteDocs(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	filter := bson.D{{"flavor", "raw sewage"}}                               //Here's our filter to look for
	//Here's how to delete MULTIPLE documents
	deleteResult, err := ic_collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the icecream collection\n", deleteResult.DeletedCount)
}
