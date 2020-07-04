package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func main() {
	// Rest of the code will go here

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error connecting to database: %v\n", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error connecting to database: %v\n", err)
	}

	fmt.Println("Connected to MongoDB!")
	//Here is an example DB,(and a personel Database to try with it, collection2)
	collection := client.Database("test").Collection("trainers")

	collection2 := client.Database("newGoLangDB").Collection("citizens") //Example 2

	//Here are some example 'Trainer' Structs that will go into our DB,(and some examples)
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	/*Example Test Structs
	 */
	address1 := address{
		Zipcode:    63129,
		Streetname: "Big City Street",
		Province:   "Provinceville State",
		Country:    "United States istan",
	}
	address2 := address{
		Zipcode:    64556,
		Streetname: "Ugly Town Ville",
		Province:   "The Bible Belt State",
		Country:    "Middle Eastern Country",
	}
	address3 := address{
		Zipcode:    33412,
		Streetname: "Ugly Street",
		Province:   "Newer Rebuilt New York",
		Country:    "A Country",
	}

	citizen1 := citizen{
		Firstname:      "Jimmy",
		Lastname:       "Bobby",
		Ethnicity:      "Black",
		Skincolor:      "Greay",
		Age:            56,
		SS:             456321456,
		Origincountry:  "Pakistanistan",
		Sex:            'M',
		Gender:         "Bi-directional",
		Citizennumber:  70647837586331888,
		Employedhere:   false,
		Employeenum:    52108560147860712,
		Currentaddress: address3,
	}

	citizen2 := citizen{
		Firstname:      "Carl",
		Lastname:       "Jr.",
		Ethnicity:      "White",
		Skincolor:      "Middle Eastern",
		Age:            22,
		SS:             456589456,
		Origincountry:  "Countryville",
		Sex:            'F',
		Gender:         "Ugly",
		Citizennumber:  44534444,
		Employedhere:   false,
		Employeenum:    223223,
		Currentaddress: address2,
	}
	citizen3 := citizen{
		Firstname:      "Beamus",
		Lastname:       "theThird.",
		Ethnicity:      "A color",
		Skincolor:      "Red",
		Age:            88,
		SS:             996589456,
		Origincountry:  "Origin Country",
		Sex:            'f',
		Gender:         "Cute boy",
		Citizennumber:  333,
		Employedhere:   true,
		Employeenum:    21,
		Currentaddress: address1,
	}

	//Here's how to insert a single document into our DB
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
		println(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//Here's our inserted first Citizen...
	insertResult2, err2 := collection2.InsertOne(context.TODO(), citizen1)
	if err2 != nil {
		log.Fatal(err2)
		println(err2)
	}
	fmt.Println("Inserted a single document: ", insertResult2.InsertedID)
	//Here's how to insert MULTIPLE Documents into a DB
	trainers := []interface{}{misty, brock}
	citizens := []interface{}{citizen2, citizen3}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	insertManyResult3, err3 := collection2.InsertMany(context.TODO(), citizens)

	if err3 != nil {
		log.Fatal(err3)
		fmt.Println(err3)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	fmt.Println("Inserted multiple documents: ", insertManyResult3.InsertedIDs)

	//Here's how to update a document with a filter 'BSON' Json object
	filter := bson.D{{"name", "Ash"}}
	filter2 := bson.D{{"firstname", "Carl"}} //Our new BSON object.

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	update2 := bson.D{
		{"$set", bson.D{
			{"firstname", "BigG Thanos"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	//Our new UpdateResult
	updateResult2, err2 := collection2.UpdateOne(context.TODO(), filter2, update2)

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult2.MatchedCount, updateResult2.ModifiedCount)

	//Here's how to find a single document and apply it to a struct.
	var result Trainer
	var result2 citizen //Here's our example Citizen to find.

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	err2 = collection.FindOne(context.TODO(), filter2).Decode(&result2)
	if err2 != nil {
		log.Fatal(err2)
		fmt.Println(err2)
	}

	fmt.Printf("Found a single document: %+v\n", result)
	fmt.Printf("I also Found a single document: %+v\n", result2)

	//Here's how to find and assign multiple Documents using a cursor
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	findOptions2 := options.Find() //These are our citizen finds
	findOptions2.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Trainer
	var results2 []*citizen
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	cur2, err4 := collection2.Find(context.TODO(), bson.D{{}}, findOptions2) //Finding our citizen
	if err4 != nil {
		log.Fatal(err4)
		fmt.Println(err4)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur2.Next(context.TODO()) { //Loop over our second Citizens cursor
		var elem2 citizen
		err := cur.Decode(&elem2)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		results2 = append(results2, &elem2)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	cur2.Close(context.TODO()) //Close our citizen cursor

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results2)

	//Here's how to delete MULTIPLE documents
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//This closes connection to the Mongo DB
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
