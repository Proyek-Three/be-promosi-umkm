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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// INSERT PRODUCT
func InsertProduct(db *mongo.Database, col string, product model.Product) (insertedID primitive.ObjectID, err error) {
	// Logging untuk memastikan data diterima dengan benar
	fmt.Printf("Inserting product: %+v\n", product)

	// Menyusun dokumen BSON untuk produk
	productData := bson.M{
		"product_name": product.ProductName,
		"description":  product.Description,
		"image":        product.Image,
		"price":        product.Price,
		"category": bson.M{
			"_id":           product.Category.ID,
			"category_name": product.Category.CategoryName,
		},
		"store": bson.M{
			"_id":        product.Store.ID,
			"store_name": product.Store.StoreName,
			"address":    product.Store.Address,
		},
	}

	// Melakukan insert ke MongoDB
	collection := db.Collection(col)
	result, err := collection.InsertOne(context.TODO(), productData)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Mengembalikan ID yang baru di-insert
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return insertedID, nil
}





// ALL
func GetAllProduct(db *mongo.Database, col string) (data []model.Product) {
	productdata := db.Collection(col)
	filter := bson.M{}
	cursor, err := productdata.Find(context.TODO(), filter)
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
func GetProductFromID(_id primitive.ObjectID, db *mongo.Database, col string) (product model.Product, errs error) {
	productdata := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := productdata.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return product, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return product, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}
	return product, nil
}

// UPDATE
func UpdateProduct(db *mongo.Database, col string, id primitive.ObjectID, ProductName string, Description string, Image string, Price float64, CategoryName model.Category, StoreName model.Store, Address model.Store) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"product_name":  ProductName,
			"description":   Description,
			"image":         Image,
			"price":         Price,
			"category_name": CategoryName,
			"store_name":    StoreName,
			"address":       Address,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateProduct: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteProductByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
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
