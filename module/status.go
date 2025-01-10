package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// INSERT STATUS
func InsertStatus(db *mongo.Database, col string, status model.Status) (insertedID primitive.ObjectID, err error) {
	// Membuat dokumen BSON untuk disimpan di MongoDB
	statusData := bson.M{
		"status": status.Status,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), statusData)
	if err != nil {
		fmt.Printf("InsertStatus: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// GET ALL
func GetAllStatus(db *mongo.Database, col string) (data []model.Status) {
	statusData := db.Collection(col)
	filter := bson.M{}
	cursor, err := statusData.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllStatus :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// GET BY ID
func GetStatusFromID(_id primitive.ObjectID, db *mongo.Database, col string) (status model.Status, errs error) {
	statusData := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := statusData.FindOne(context.TODO(), filter).Decode(&status)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return status, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return status, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}
	return status, nil
}

// UPDATE
func UpdateStatus(db *mongo.Database, col string, id primitive.ObjectID, newStatus string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"status": newStatus,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateStatus: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

// DELETE
func DeleteStatusByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	statusData := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := statusData.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
