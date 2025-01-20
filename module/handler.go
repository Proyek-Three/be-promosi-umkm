package module

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Fungsi untuk register
func CreateUser(collection *mongo.Collection, name, username, email, password, role, phoneNumber string) (model.Users, error) {
	// Validasi input
	if name == "" || username == "" || email == "" || password == "" {
		return model.Users{}, errors.New("name, username, email, and password are required")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Buat objek user
	user := model.Users{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Role:        role,
		PhoneNumber: phoneNumber,
		Username:    username,
		Email:       email,
		Password:    string(hashedPassword),
	}

	// Masukkan ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to insert user into MongoDB: %w", err)
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

// Fungsi untuk memasukkan admin baru
func RegisUser(db *mongo.Database, col string, user model.Users) (primitive.ObjectID, error) {
	if user.Username == "" || user.Password == "" || user.Email == "" || user.Store.StoreName == "" {
		return primitive.NilObjectID, fmt.Errorf("username, password, email, and store name cannot be empty")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = string(hashedPassword)

	// Set Store ID jika belum ada
	if user.Store.ID.IsZero() {
		user.Store.ID = primitive.NewObjectID()
	}

	// Simpan ke database
	result, err := db.Collection(col).InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// Fungsi untuk mendapatkan admin berdasarkan username
func GetUserByUsernameOrEmail(db *mongo.Database, col, username, email string) (*model.User, error) {
	var admin model.User
	err := db.Collection(col).FindOne(context.Background(), bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func GetUserByUsername(db *mongo.Database, collectionName string, username string) (*model.Users, error) {
	var user model.Users
	collection := db.Collection(collectionName)

	// Query filter berdasarkan username
	filter := bson.M{"username": username}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User tidak ditemukan
		}
		return nil, err // Error saat query
	}

	return &user, nil
}


func ValidatePassword(hashedPassword, plainPassword string) bool {
	// Gunakan bcrypt untuk membandingkan hash password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil // Jika tidak ada error, password valid
}
