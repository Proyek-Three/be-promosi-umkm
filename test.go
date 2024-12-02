package _proyek3

import (
	"fmt"
	"testing"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"github.com/Proyek-Three/be-promosi-umkm/module"

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
		Address:    "Jl. Sudirman No. 1 Jakarta Pusat", 
	}

	productdata := model.Product{
		ProductName:  productName,
		Description: description,
		Image:    image,
		Price: price,
		CategoryName:  product_category,
		StoreName: store,
		Address: store,
	}

	insertedID, err := module.InsertProduct(module.MongoConn, "product", productdata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
