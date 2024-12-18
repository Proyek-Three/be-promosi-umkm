package module

import (
	"context"
	"fmt"
	"time"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	// Set ID yang baru dibuat
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

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

// Fungsi untuk memasukkan admin baru
func InsertAdmin(db *mongo.Database, col string, username, password, email string) (primitive.ObjectID, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}

	admin := bson.M{
		"username": username,
		"email":     email,
		"password":  string(hashedPassword),
	}

	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// Fungsi untuk mendapatkan admin berdasarkan username
func GetAdminByUsername(db *mongo.Database, col string, username string) (*model.User, error) {
	var admin model.User
	err := db.Collection(col).FindOne(context.Background(), bson.M{"username": username}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// Fungsi untuk mendapatkan admin berdasarkan email
func GetAdminByEmail(db *mongo.Database, col string, email string) (*model.User, error) {
	var admin model.User
	err := db.Collection(col).FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func ValidatePassword(hashedPassword, plainPassword string) bool {
	// Gunakan bcrypt untuk membandingkan hash password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // Jika tidak ada error, password valid
}