package module

import (
	"context"
	"errors"
	"fmt"

	// "time"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// func MongoConnect(dbname string) (db *mongo.Database) {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
// 	if err != nil {
// 		fmt.Printf("MongoConnect: %v\n", err)
// 	}
// 	return client.Database(dbname)
// }

// func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
// 	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
// 	if err != nil {
// 		fmt.Printf("InsertOneDoc: %v\n", err)
// 	}
// 	return insertResult.InsertedID
// }

// INSERT PRODUCT
func InsertStore(db *mongo.Database, col string, store model.Store) (insertedID primitive.ObjectID, err error) {
	// Membuat dokumen BSON untuk disimpan di MongoDB
	categorydata := bson.M{
		"store_name": store.StoreName,
		"address":    store.Address,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), categorydata)
	if err != nil { //Jika terjadi kesalahan saat menyisipkan dokumen, maka akan mengembalikan pesan kesalahan
		fmt.Printf("InsertStore: %v\n", err) //mencetak pesan kesalahan ke console
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID) //Mengambil ID dari dokumen yang baru saja disisipkan dan mengubahnya ke tipe primitive.ObjectID.
	return insertedID, nil                              //mengembalikan insertedID dan nil sebagai nilai err jika tidak ada kesalahan.
}

// ALL
func GetAllStore(db *mongo.Database, col string) (data []model.Store) {
	categorydata := db.Collection(col)
	filter := bson.M{}
	cursor, err := categorydata.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// ID
func GetStoreFromID(_id primitive.ObjectID, db *mongo.Database, col string) (category model.Store, errs error) {
	categorydata := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := categorydata.FindOne(context.TODO(), filter).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return category, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return category, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}
	return category, nil
}

// UPDATE
func UpdateStore(db *mongo.Database, col string, id primitive.ObjectID, StoreName string, Address string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"store_name": StoreName,
			"address":    Address,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateStore: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteStoreByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	productdata := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := productdata.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error()) //mengembalikan pesan kesalahan jika terjadi kesalahan saat menghapus data
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
