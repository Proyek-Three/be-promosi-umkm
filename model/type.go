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
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Category    Category           `bson:"category,omitempty" json:"category,omitempty"`
	Store       Store              `bson:"store,omitempty" json:"store,omitempty"`
}

type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CategoryName string             `bson:"category_name,omitempty" json:"category_name,omitempty"`
}

type Store struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StoreName string             `json:"store_name,omitempty" bson:"store_name,omitempty"`
	Address   string             `json:"address,omitempty" bson:"address,omitempty"`
}




// Model untuk User
type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username,omitempty" json:"username"`
	Email    string `bson:"email,omitempty" json:"email"`
	Password string `bson:"password,omitempty" json:"password"`
}

// Model untuk permintaan registrasi
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}