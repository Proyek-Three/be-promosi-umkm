package module

import (
	"context"
	"fmt"
	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func InsertUser(db *mongo.Database, col string, user model.DataUsers, storeID primitive.ObjectID) (insertedID primitive.ObjectID, err error) {
	if user.Username == "" || user.Password == "" {
		return primitive.NilObjectID, fmt.Errorf("username and password cannot be empty")
	}

	// Periksa apakah store_id valid
	storeCollection := db.Collection("stores")
	var store model.Store
	err = storeCollection.FindOne(context.TODO(), bson.M{"_id": storeID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, fmt.Errorf("store with ID %s not found", storeID.Hex())
		}
		return primitive.NilObjectID, fmt.Errorf("failed to fetch store: %w", err)
	}

	// Tambahkan informasi store ke user
	user.ID = primitive.NewObjectID()
	user.Store = store

	collection := db.Collection(col)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert user: %w", err)
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to parse inserted ID")
	}

	return insertedID, nil
}


func GetAllUsers(db *mongo.Database, col string) (data []model.DataUsers, err error) {
	collection := db.Collection(col)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %w", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding users: %w", err)
	}

	return data, nil
}


func GetUserByID(_id primitive.ObjectID, db *mongo.Database, col string) (user model.DataUsers, err error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("no user found with ID: %s", _id.Hex())
		}
		return user, fmt.Errorf("error retrieving user with ID %s: %w", _id.Hex(), err)
	}
	return user, nil
}


func UpdateUser(db *mongo.Database, col string, userID primitive.ObjectID, updatedUser model.DataUsers, storeID primitive.ObjectID) error {
	if updatedUser.Username == "" || updatedUser.Password == "" {
		return fmt.Errorf("username and password cannot be empty")
	}

	// Periksa apakah store_id valid
	storeCollection := db.Collection("stores")
	var store model.Store
	err := storeCollection.FindOne(context.TODO(), bson.M{"_id": storeID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("store with ID %s not found", storeID.Hex())
		}
		return fmt.Errorf("failed to fetch store: %w", err)
	}

	// Siapkan data untuk update
	updateData := bson.M{
		"$set": bson.M{
			"username":        updatedUser.Username,
			"password":        updatedUser.Password,
			"phone_number":    updatedUser.PhoneNumber,
			"email":           updatedUser.Email,
			"store":           store,
		},
	}

	collection := db.Collection(col)
	filter := bson.M{"_id": userID}
	result, err := collection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no user found with ID: %s", userID.Hex())
	}

	return nil
}


func DeleteUserByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting user with ID %s: %w", _id.Hex(), err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no user found with ID: %s", _id.Hex())
	}

	return nil
}
