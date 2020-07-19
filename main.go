package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
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

type MongoHotDog struct {
	HotDogType  string   `json:"HotDogType"`
	Condiments  []string `json:"Condiments"`
	Calories    int      `json:"Calories"`
	Name        string   `json:"Name"`
	FoodID      int      `json:"FoodID"`
	UserID      int      `json:"UserID"` //User WHOMST this hotDog belongs to
	DateCreated string   `json:"DateCreated"`
	DateUpdated string   `json:"DateUpdated"`
}

type MongoHotDogs struct {
	Hotdogs []MongoHotDog `json:"Hotdogs"`
}

type MongoHamburger struct {
	BurgerType  string   `json:"BurgerType"`
	Condiments  []string `json:"Condiments"`
	Calories    int      `json:"Calories"`
	Name        string   `json:"Name"`
	FoodID      int      `json:"FoodID"`
	UserID      int      `json:"UserID"` //User WHOMST this hotDog belongs to
	DateCreated string   `json:"DateCreated"`
	DateUpdated string   `json:"DateUpdated"`
}

type MongoHamburgers struct {
	Hamburgers []MongoHamburger `json:"Hamburgers"`
}

/* Mongo No-SQL Variable Declarations */
type AUser struct { //Using this for Mongo
	UserName    string         `json:"UserName"`
	Password    string         `json:"Password"` //This was formally a []byte but we are changing our code to fit the database better
	First       string         `json:"First"`
	Last        string         `json:"Last"`
	Role        string         `json:"Role"`
	UserID      int            `json:"UserID"`
	DateCreated string         `json:"DateCreated"`
	DateUpdated string         `json:"DateUpdated"`
	Hotdogs     MongoHotDogs   `json:"Hotdogs"`
	Hamburgers  MongoHamburger `json:"Hamburgers"`
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
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
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
	//insertData(client) //Insert 1 data

	//insertMultipleData(client) //Insert ALOT of data

	//dataUpdateAll(client) //Update Document

	//retrieveData(client) //Get Document

	//retrieveMultipleData(client) //Get Multiple Document

	//deleteDocs(client)//Delete Multiple Documents
	//Test for Other Projects
	//testFindUser(client)
	//testDeleteCertainItems(client)
	//testUpdateCertainItems(client)
	//testBigDelete(client)
	testBigFind(client)

	fmt.Println()
	fmt.Println()

	var anything []interface{}
	var stringList []string
	var anArray []int
	stringList = append(stringList, "first word")
	stringList = append(stringList, "Second word")
	stringList = append(stringList, "Third Word")
	anArray = append(anArray, 1, 2, 3)
	for _, val := range stringList {
		anything = append(anything, val)
	}
	for _, val := range anArray {
		anything = append(anything, val)
	}

	updateInterfaces := []interface{}{}
	updateInterfaces = append(updateInterfaces, stringList, anArray)

	fmt.Println(anything[4])
	fmt.Println(updateInterfaces)
	fmt.Println(updateInterfaces[0])
	anInterface := updateInterfaces[1]
	fmt.Println(updateInterfaces[1])
	fmt.Printf("Here is anInterface: %v\n", anInterface)
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

	filter := bson.D{{"flavor", "test flavor the one"},
		{"alias", "Tester Cone"}} //Here's our filter to look for

	update := bson.D{ //Here is our data to update
		{"$set", bson.D{
			{"flavor", "raw sewage"},
		}},
		{"$set", bson.D{
			{"calories", 10000005},
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
	/* WARNING...THIS ALSO MIGHT WORK BETTER WITH A BSON M!!! */
	//Here's how to find a single document and apply it to a struct.
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	var result icecream
	filter := bson.D{{"flavor", "raw sewage"}} //Here's our filter to look for

	err := ic_collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		gotIt := strings.Contains(err.Error(), "no documents in results")
		if gotIt == true {
			fmt.Printf("No documents found, but we good.\n")
		} else {
			fmt.Printf("Error searching for a document: %v\n", err.Error())
			log.Fatal(err)
		}
	} else {
		fmt.Printf("We found a document: %v\n", result)
	}

	fmt.Printf("Found a single document: %+v\n", result)
}

func retrieveMultipleData(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	filter := bson.D{{"calories", 800}}                                      //Here's our filter to look for
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

func testFindUser(client *mongo.Client) {
	idString := "57324802"
	theID, _ := strconv.Atoi(idString)
	user_collection := client.Database("superdbtest1").Collection("users")
	filterUserID := bson.D{{"userid", theID}}
	var testAUser AUser
	theErr := user_collection.FindOne(context.TODO(), filterUserID).Decode(&testAUser)
	if theErr != nil {
		if strings.Contains(theErr.Error(), "no documents in result") {
			fmt.Printf("It's all good, this document wasn't found for User and our ID is clean.\n")
		} else {
			fmt.Printf("DEBUG: We have another error for finding a unique UserID: \n%v\n", theErr)
		}
	}
	fmt.Printf("Found the testUser: %v\n", testAUser)
}

func testDeleteCertainItems(client *mongo.Client) {
	ic_collection := client.Database("superdbtest1").Collection("icecreams") //Here's our collection
	// list of deletes
	deletes := []bson.M{
		{"flavor": "weed"},
		{"flavor": "cocaine"},
		{"flavor": "test flavor"},
	}
	deletes = append(deletes, bson.M{"flavor": bson.M{
		"$eq": "test flavor the two",
	}})

	// create the slice of write models
	var writes []mongo.WriteModel

	for _, del := range deletes {
		model := mongo.NewDeleteManyModel().SetFilter(del)
		writes = append(writes, model)
	}

	// run bulk write
	res, err := ic_collection.BulkWrite(context.TODO(), writes)
	if err != nil {
		log.Fatal(err)
	}
	//Print Results
	fmt.Printf(
		"insert: %d, updated: %d, deleted: %d",
		res.InsertedCount,
		res.ModifiedCount,
		res.DeletedCount,
	)
}

func testUpdateCertainItems(client *mongo.Client) {
	theTimeNow := time.Now()
	postedHotDogs := MongoHotDogs{
		Hotdogs: []MongoHotDog{},
	}
	newHotDog := MongoHotDog{
		HotDogType:  "testDoggo1",
		Condiments:  []string{"one condiment", "two condiment"},
		Calories:    777,
		Name:        "The test dog",
		FoodID:      77777777,
		UserID:      17721576,
		DateCreated: theTimeNow.Format("2006-01-02 15:04:05"),
		DateUpdated: theTimeNow.Format("2006-01-02 15:04:05"),
	}
	postedHotDogs.Hotdogs = append(postedHotDogs.Hotdogs, newHotDog)

	userCollection := client.Database("superdbtest1").Collection("users")
	var userGot AUser
	afilter := bson.M{
		"userid": bson.M{
			"$eq": postedHotDogs.Hotdogs[0].UserID, // check if bool field has value of 'false'
		},
	}
	theFilter := bson.M{"userid": postedHotDogs.Hotdogs[0].UserID}
	err := userCollection.FindOne(context.TODO(), theFilter).Decode(&userGot)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			fmt.Printf("No biggie, we just didn't find user documents for this ID: %v\n", postedHotDogs.Hotdogs[0].UserID)
		} else {
			fmt.Printf("We had trouble finding User with this ID: %v\n", postedHotDogs.Hotdogs[0].UserID)
			log.Fatal(err)
		}
	} else {
		fmt.Printf("We got the User from Mongo DB: \n%v\n\n", userGot)
	}
	var theHotDogs MongoHotDogs
	theHotDogs = userGot.Hotdogs
	theHotDogs.Hotdogs = append(theHotDogs.Hotdogs, postedHotDogs.Hotdogs...)
	//Update the User with new User
	theTimeNow = time.Now()
	update1 := bson.D{
		{"$set", bson.D{{"hotdogs", theHotDogs}}},
	}
	update2 := bson.D{
		{"$set", bson.D{{"dateupdated", theTimeNow.Format("2006-01-02 15:04:05")}}},
	}
	update3 := bson.D{
		{"$set", bson.D{{"username", "sillybit"}}},
	}
	updateInterfaces := []interface{}{}
	filterInterfaces := []interface{}{}
	filterInterfaces = append(filterInterfaces, theFilter, theFilter)
	updateInterfaces = append(updateInterfaces, update1, update2, update3)
	updateResult, err := userCollection.UpdateOne(context.TODO(),
		afilter,
		updateInterfaces)
	if err != nil {
		fmt.Printf("We had trouble updating User with this ID: %v\n", postedHotDogs.Hotdogs[0].UserID)
		log.Fatal(err)
	} else {
		fmt.Printf("Updated this UserID's Hotdogs: %v\n", updateResult.ModifiedCount)
	}
}

func testBigDelete(client *mongo.Client) {
	var theIDS []int
	theIDS = append(theIDS, 50177074, 12228170)
	var bsonArrayFilters []bson.M
	deleteInterfaces := []interface{}{}
	/*
		for p := 0; p < len(theIDS); p++ {
			filter := bson.D{{"userid", theIDS[p]}}
			deleteInterfaces = append(deleteInterfaces, filter)
		}
	*/

	filter := bson.M{"userid": 12228170}
	bsonArrayFilters = append(bsonArrayFilters, filter)
	deleteInterfaces = append(deleteInterfaces, bsonArrayFilters)
	myDeleteInterface := deleteInterfaces[0]
	fmt.Printf("Here is our myDeleteInterface: %v\n", myDeleteInterface)
	hotdog_collection := client.Database("superdbtest1").Collection("hotdogs") //Here's our collection
	//Here's how to delete MULTIPLE documents
	for z := 0; z < len(theIDS); z++ {
		deleteResult, err := hotdog_collection.DeleteMany(context.TODO(), myDeleteInterface)
		if err != nil {
			fmt.Printf("There was an error deleting hotdogs: %v\n", err)
			log.Fatal(err)
		} else {
			fmt.Printf("Deleted %v documents in the hotdog collection\n", deleteResult.DeletedCount)
		}
	}
}

func testBigFind(client *mongo.Client) {
	//Check hotdog collection
	theID := 65732442
	hotdogCollection := client.Database("superdbtest1").Collection("hotdogs") //Here's our collection
	var testHotdog MongoHotDog
	theFilter := bson.M{
		"$or": []interface{}{
			bson.M{"userid": theID},
			bson.M{"foodid": theID},
		},
	}
	theErr := hotdogCollection.FindOne(context.TODO(), theFilter).Decode(&testHotdog)
	if theErr != nil {
		if strings.Contains(theErr.Error(), "no documents in result") {
			fmt.Printf("It's all good, this document wasn't found for User and our ID is clean.\n")
		} else {
			fmt.Printf("DEBUG: We have another error for finding a unique UserID: \n%v\n", theErr)
			canExit := false
			fmt.Println(canExit)
			log.Fatal(theErr)
		}
	} else {
		fmt.Printf("We found the resulting hotdog: %v\n", testHotdog)
	}
}
