package module

import (
	"context"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Proyek-Three/be-promosi-umkm/model"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fungsi untuk membuat user baru
func CreateUser(collection *mongo.Collection, username, email, password string) (model.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	// Buat objek user
	user := model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Masukkan ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Fungsi untuk mengecek apakah email sudah ada
func IsEmailExist(collection *mongo.Collection, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func InsertAdmin(db *mongo.Database, col string, username string, password string, email string) (insertedID primitive.ObjectID, err error) {
    admin := bson.M{
    "user_name" : username,
    "email"        : email,
    "password"    : password,
    }
    result, err := db.Collection(col).InsertOne(context.Background(), admin)
    if err != nil {
        fmt.Printf("InsertAdmin: %v\n", err)
        return
    }
    insertedID = result.InsertedID.(primitive.ObjectID)
    return insertedID, nil
}