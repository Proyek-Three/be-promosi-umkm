package _proyek3

import (
	"context"
	"testing"
	"github.com/Proyek-Three/be-promosi-umkm/model"
	"github.com/Proyek-Three/be-promosi-umkm/module"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertUser(t *testing.T) {
	// Setup MongoDB client and context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dzulkiflifaiz11:SAKTIMlucu12345@webservice.dqol9t4.mongodb.net/"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database("test_db")
	col := "users"

	// Create a valid store for testing
	storeID := primitive.NewObjectID()
	store := model.Store{
		ID:        storeID,
		StoreName: "Mie Gacoan",
		Address:   "Jln Surya Sumantri",
	}

	storeCollection := db.Collection("stores")
	_, err = storeCollection.InsertOne(context.TODO(), store)
	if err != nil {
		t.Fatalf("Failed to insert store: %v", err)
	}

	// Create a DataUser instance with the store ID
	user := model.DataUsers{
		Username:  "julparker",
		Password:  "securepassword",
		PhoneNumber: "1234567890",
		Email:     "johndoe@example.com",
		Store:     store,
	}

	// Call InsertUser function
	insertedID, err := module.InsertUser(db, col, user, storeID)
	assert.NoError(t, err)
	assert.NotEqual(t, insertedID, primitive.NilObjectID)

	// Verify the user is inserted into the collection
	var insertedUser model.DataUsers
	err = db.Collection(col).FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedUser)
	assert.NoError(t, err)
	assert.Equal(t, "julparker", insertedUser.Username)
	assert.Equal(t, storeID, insertedUser.Store.ID)
}

func TestGetAllUsers(t *testing.T) {
	// Setup MongoDB client and context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dzulkiflifaiz11:SAKTIMlucu12345@webservice.dqol9t4.mongodb.net/"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database("test_db")
	col := "users"

	// Insert a user to test GetAllUsers
	storeID := primitive.NewObjectID()
	store := model.Store{
		ID:        storeID,
		StoreName: "Mie Gacoan",
		Address:   "Jln Surya Sumantri",
	}

	storeCollection := db.Collection("stores")
	_, err = storeCollection.InsertOne(context.TODO(), store)
	if err != nil {
		t.Fatalf("Failed to insert store: %v", err)
	}

	user := model.DataUsers{
		Username:  "julparker",
		Password:  "securepassword",
		PhoneNumber: "1234567890",
		Email:     "johndoe@example.com",
		Store:     store,
	}

	_, err = module.InsertUser(db, col, user, storeID)
	assert.NoError(t, err)

	// Call GetAllUsers function
	users, err := module.GetAllUsers(db, col)
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestGetUserByID(t *testing.T) {
	// Setup MongoDB client and context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dzulkiflifaiz11:SAKTIMlucu12345@webservice.dqol9t4.mongodb.net/"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database("test_db")
	col := "users"

	// Create a store and insert a user
	storeID := primitive.NewObjectID()
	store := model.Store{
		ID:        storeID,
		StoreName: "Mie Gacoan",
		Address:   "Jln Surya Sumantri",
	}

	storeCollection := db.Collection("stores")
	_, err = storeCollection.InsertOne(context.TODO(), store)
	if err != nil {
		t.Fatalf("Failed to insert store: %v", err)
	}

	user := model.DataUsers{
		Username:  "julparker",
		Password:  "securepassword",
		PhoneNumber: "1234567890",
		Email:     "johndoe@example.com",
		Store:     store,
	}

	insertedID, err := module.InsertUser(db, col, user, storeID)
	assert.NoError(t, err)

	// Call GetUserByID function
	fetchedUser, err := module.GetUserByID(insertedID, db, col)
	assert.NoError(t, err)
	assert.Equal(t, "julparker", fetchedUser.Username)
}

func TestUpdateUser(t *testing.T) {
	// Setup MongoDB client and context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dzulkiflifaiz11:SAKTIMlucu12345@webservice.dqol9t4.mongodb.net/"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database("test_db")
	col := "users"

	// Create a store and insert a user
	storeID := primitive.NewObjectID()
	store := model.Store{
		ID:        storeID,
		StoreName: "Mie Gacoan",
		Address:   "Jln Surya Sumantri",
	}

	storeCollection := db.Collection("stores")
	_, err = storeCollection.InsertOne(context.TODO(), store)
	if err != nil {
		t.Fatalf("Failed to insert store: %v", err)
	}

	user := model.DataUsers{
		Username:  "julparker",
		Password:  "securepassword",
		PhoneNumber: "1234567890",
		Email:     "johndoe@example.com",
		Store:     store,
	}

	insertedID, err := module.InsertUser(db, col, user, storeID)
	assert.NoError(t, err)

	// Update user details
	updatedUser := model.DataUsers{
		Username:  "julparker_updated",
		Password:  "newpassword",
		PhoneNumber: "0987654321",
		Email:     "newemail@example.com",
		Store:     store,
	}

	err = module.UpdateUser(db, col, insertedID, updatedUser, storeID)
	assert.NoError(t, err)

	// Verify the user is updated
	fetchedUser, err := module.GetUserByID(insertedID, db, col)
	assert.NoError(t, err)
	assert.Equal(t, "julparker_updated", fetchedUser.Username)
}

func TestDeleteUserByID(t *testing.T) {
	// Setup MongoDB client and context
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://dzulkiflifaiz11:SAKTIMlucu12345@webservice.dqol9t4.mongodb.net/"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())
	db := client.Database("test_db")
	col := "users"

	// Create a store and insert a user
	storeID := primitive.NewObjectID()
	store := model.Store{
		ID:        storeID,
		StoreName: "Mie Gacoan",
		Address:   "Jln Surya Sumantri",
	}

	storeCollection := db.Collection("stores")
	_, err = storeCollection.InsertOne(context.TODO(), store)
	if err != nil {
		t.Fatalf("Failed to insert store: %v", err)
	}

	user := model.DataUsers{
		Username:  "julparker",
		Password:  "securepassword",
		PhoneNumber: "1234567890",
		Email:     "johndoe@example.com",
		Store:     store,
	}

	insertedID, err := module.InsertUser(db, col, user, storeID)
	assert.NoError(t, err)

	// Call DeleteUserByID function
	err = module.DeleteUserByID(insertedID, db, col)
	assert.NoError(t, err)

	// Verify the user is deleted
	_, err = module.GetUserByID(insertedID, db, col)
	assert.Error(t, err)
}

