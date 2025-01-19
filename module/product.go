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

	// Validasi ID status
	if product.Status.ID.IsZero() {
		return primitive.NilObjectID, fmt.Errorf("invalid status ID: cannot be empty")
	}

	// Validasi ID user
	if product.User.ID.IsZero() {
		return primitive.NilObjectID, fmt.Errorf("invalid status ID: cannot be empty")
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
		"status": bson.M{
			"_id":    product.Status.ID,
			"status": product.Status.Status,
		},
		"user": bson.M{
			"_id":      product.User.ID,          // ID pengguna
			"name":     product.User.Name,        // Nama pengguna
			"role":     product.User.Role,        // Peran pengguna
			"username": product.User.Username,    // Nama pengguna
			"email":    product.User.Email,       // Email pengguna
			"phone":    product.User.PhoneNumber, // Nomor telepon pengguna
			"store": bson.M{
				"id":         product.User.Store.ID,        // ID toko
				"store_name": product.User.Store.StoreName, // Nama toko
				"address":    product.User.Store.Address,   // Alamat toko
			},
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

// GetProductsByUserID mengambil produk berdasarkan user_id
func GetProductsByUserID(db *mongo.Database, col string, userID primitive.ObjectID) []model.Product {
	productCollection := db.Collection(col)

	// Filter berdasarkan user_id
	filter := bson.M{"user_id._id": userID}

	// Query ke MongoDB
	cursor, err := productCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	// Decode hasil query ke slice produk
	var products []model.Product
	err = cursor.All(context.TODO(), &products)
	if err != nil {
		fmt.Println("Error decoding products:", err)
		return nil
	}

	return products
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

func UpdateProduct(db *mongo.Database, col string, productID primitive.ObjectID, updatedProduct model.Product) error {
	// Logging untuk debugging
	fmt.Printf("Updating product ID: %s with data: %+v\n", productID.Hex(), updatedProduct)

	// Validasi ID kategori
	if updatedProduct.Category.ID.IsZero() {
		return fmt.Errorf("invalid category ID: cannot be empty")
	}

	// Validasi ID toko
	if updatedProduct.Store.ID.IsZero() {
		return fmt.Errorf("invalid store ID: cannot be empty")
	}

	// Validasi ID status
	if updatedProduct.Status.ID.IsZero() {
		return fmt.Errorf("invalid status ID: cannot be empty")
	}

	// Menyusun dokumen BSON untuk pembaruan
	updateData := bson.M{
		"$set": bson.M{
			"product_name": updatedProduct.ProductName,
			"description":  updatedProduct.Description,
			"image":        updatedProduct.Image,
			"price":        updatedProduct.Price,
			"category": bson.M{
				"_id":           updatedProduct.Category.ID,
				"category_name": updatedProduct.Category.CategoryName,
			},
			"store": bson.M{
				"_id":        updatedProduct.Store.ID,
				"store_name": updatedProduct.Store.StoreName,
				"address":    updatedProduct.Store.Address,
			},
			"status": bson.M{
				"_id":    updatedProduct.Status.ID,
				"status": updatedProduct.Status.Status,
			},
		},
	}

	// Mendapatkan koleksi dan memperbarui dokumen
	collection := db.Collection(col)
	filter := bson.M{"_id": productID}
	result, err := collection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	// Memastikan ada dokumen yang diperbarui
	if result.MatchedCount == 0 {
		return fmt.Errorf("no product found with ID: %s", productID.Hex())
	}

	fmt.Printf("Successfully updated product ID: %s\n", productID.Hex())
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
