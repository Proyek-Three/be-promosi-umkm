package _proyek3

import (
	"fmt"
	"testing"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"github.com/Proyek-Three/be-promosi-umkm/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// INSERT MENU
func TestInsertProduct(t *testing.T) {
	productName := "Dimsum"
	description := "makanan dengan rasa spesial"
	image := "image.jpg"
	price := 10000.0

	var product_category = model.Category{
		CategoryName: "Makanan",
	}

	var store = model.Store{
		StoreName: "Food Store",
		Address:   "Jl. Sudirman No. 1 Jakarta Pusat",
	}

	productdata := model.Product{
		ProductName:  productName,
		Description:  description,
		Image:        image,
		Price:        price,
		CategoryName: product_category,
		StoreName:    store,
		Address:      store,
	}

	insertedID, err := module.InsertProduct(module.MongoConn, "product", productdata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

// BY ID
func TestGetProductFromID(t *testing.T) {
	id := "667e27a6cccefc9e0156f40d"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	productdata, err := module.GetProductFromID(objectID, module.MongoConn, "product")
	if err != nil {
		t.Fatalf("error calling GetMenuFromID: %v", err)
	}
	fmt.Println(productdata)
}

// ALL
func TestGetAll(t *testing.T) {
	data := module.GetAllProduct(module.MongoConn, "product")
	fmt.Println(data)
}

// DELETE
func TestDeleteProductByID(t *testing.T) {
	id := "667e5f1c0da481424d4fae0b" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteProductByID(objectID, module.MongoConn, "product")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetProductFromID(objectID, module.MongoConn, "product")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
