package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

)


// Product represents the table Produk
type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName  string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Image        string             `json:"image,omitempty" bson:"image,omitempty"`
	Price        float64            `json:"price,omitempty" bson:"price,omitempty"`
	CategoryName Category           `bson:"category_name,omitempty" json:"category_name,omitempty"`
	StoreName    Store              `bson:"store_name,omitempty" json:"store_name,omitempty"`
	Address      Store              `bson:"address,omitempty" json:"address,omitempty"`
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

// Model untuk User
type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

// Model untuk permintaan registrasi
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}