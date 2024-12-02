package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users represents the table Users
type Users struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Role        string             `json:"role,omitempty" bson:"role,omitempty"`
	Username    string             `json:"username,omitempty" bson:"username,omitempty" gorm:"unique;not null"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	PhoneNumber string             `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email,omitempty"`
}

// Product represents the table Produk
type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName  string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Name         string             `json:"nama_toko,omitempty" bson:"nama_toko,omitempty"`
	Description  string             `json:"deskripsi,omitempty" bson:"deskripsi,omitempty"`
	Image        string             `json:"gambar,omitempty" bson:"gambar,omitempty"`
	Price        int                `json:"harga,omitempty" bson:"harga,omitempty"`
	CategoryName Category           `bson:"category_name,omitempty" json:"category_name,omitempty"`
	StoreName    Store              `bson:"store_name,omitempty" json:"store_name,omitempty"`
}

// Store represents the table Store
type Store struct {
	ID        int    `json:"_id,omitempty" bson:"id,omitempty"`
	StoreName string `json:"store_name,omitempty" bson:"store_name,omitempty"`
	Address   string `json:"address,omitempty" bson:"address,omitempty"`
}

// Category represents the table Category
type Category struct {
	ID           int    `json:"id,omitempty" bson:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	CategoryName string `json:"category_name,omitempty" bson:"category_name,omitempty"`
}
