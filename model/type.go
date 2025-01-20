package model

import (
	"time"

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
	Store       Store              `json:"store,omitempty" bson:"store,omitempty"`
}

// Product represents the table Produk
type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User         Users              `bson:"user,omitempty" json:"user,omitempty"`
	ProductName  string             `bson:"product_name,omitempty" json:"product_name,omitempty"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
	Image        string             `bson:"image,omitempty" json:"image,omitempty"`
	Price        float64            `bson:"price,omitempty" json:"price,omitempty"`
	Category     Category           `bson:"category,omitempty" json:"category,omitempty"`
	Status       Status             `bson:"status,omitempty" json:"status,omitempty"`
	StoreName    string             `bson:"store_name,omitempty" json:"store_name,omitempty"`
	StoreAddress string             `bson:"store_address,omitempty" json:"store_address,omitempty"`
}

type Status struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Status string             `bson:"status,omitempty" json:"status,omitempty"`
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
	Role     string `bson:"role,omitempty" json:"role"`
}

// Model untuk permintaan registrasi
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string             `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Token struct {
	ID        string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Token     string    `bson:"token,omitempty" json:"token,omitempty"`
	AdminID   string    `bson:"admin_id,omitempty" json:"admin_id,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type DataUsers struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username    string             `bson:"username,omitempty" json:"username,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	//Store       Store              `bson:"store,omitempty" json:"store,omitempty"`
}
