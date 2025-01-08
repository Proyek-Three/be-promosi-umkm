package module

import (
	"context"
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
	// Logging untuk debugging
	fmt.Printf("Inserting product: %+v\n", product)

	// Validasi ID kategori
	if product.Category.ID.IsZero() {
		return primitive.NilObjectID, fmt.Errorf("invalid category ID: cannot be empty")
	}

	// Validasi ID toko
	if product.Store.ID.IsZero() {
		return primitive.NilObjectID, fmt.Errorf("invalid store ID: cannot be empty")
	}

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

	// Menyisipkan dokumen ke MongoDB
	collection := db.Collection(col)
	result, err := collection.InsertOne(context.TODO(), productData)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert product: %w", err)
	}

	// Mendapatkan ID yang disisipkan
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to parse inserted ID")
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

func UpdateProduct(
	db *mongo.Database, 
	col string, 
	id primitive.ObjectID, 
	ProductName string, 
	Description string, 
	Image string, 
	Price float64, 
	Category model.Category, 
	StoreID primitive.ObjectID,
) error {
	// Cari store berdasarkan ID
	var store model.Store
	err := db.Collection("stores").FindOne(context.TODO(), bson.M{"_id": StoreID}).Decode(&store)
	if err != nil {
		return fmt.Errorf("store ID not found: %w", err)
	}

	// Cari kategori berdasarkan ID (opsional, jika ingin konsisten)
	var category model.Category
	if !Category.ID.IsZero() {
		err = db.Collection("categories").FindOne(context.TODO(), bson.M{"_id": Category.ID}).Decode(&category)
		if err != nil {
			return fmt.Errorf("category ID not found: %w", err)
		}
		Category.CategoryName = category.CategoryName
	}

	// Filter dokumen berdasarkan ID
	filter := bson.M{"_id": id}

	// Update dokumen
	update := bson.M{
		"$set": bson.M{
			"product_name": ProductName,
			"description":  Description,
			"image":        Image,
			"price":        Price,
			"category": bson.M{
				"_id":           Category.ID,
				"category_name": Category.CategoryName,
			},
			"store": bson.M{
				"_id":        store.ID,
				"store_name": store.StoreName,
				"address":    store.Address,
			},
		},
	}

	// Eksekusi update
	result, err := db.Collection(col).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no data has been changed with the specified ID")
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
