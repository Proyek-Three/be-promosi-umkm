package module

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/Proyek-Three/bp-promosi-umkm/config"

	"github.com/Proyek-Three/be-promosi-umkm/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Login authenticates an admin and generates a JWT token
func Login(db *mongo.Database, username string, password string) (string, error) {
	var admin model.Admin
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the admin by username
	err := db.Collection("Admin").FindOne(ctx, bson.M{"user_name": username}).Decode(&admin)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("error finding user: %v", err)
	}

	// Check password match
	if admin.Password != password {
		return "", fmt.Errorf("invalid password")
	}

	// Generate JWT token
	token, err := config.GenerateJWT(admin)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	// Save the token to the database
	err = SaveTokenToDatabase(db, "Tokens", admin.ID.Hex(), token)
	if err != nil {
		return "", fmt.Errorf("error saving token to database: %v", err)
	}

	return token, nil
}

// DeleteTokenFromMongoDB deletes a token from the database
func DeleteTokenFromMongoDB(db *mongo.Database, token string) error {
	collection := db.Collection("Tokens")
	filter := bson.M{"token": token}

	// Delete the token
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// SaveTokenToDatabase stores or updates the token in the database
func SaveTokenToDatabase(db *mongo.Database, col string, adminID string, token string) error {
	collection := db.Collection(col)
	filter := bson.M{"admin_id": adminID}
	update := bson.M{
		"$set": bson.M{
			"token":      token,
			"updated_at": time.Now(),
		},
	}

	// Upsert token for the admin
	_, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

// GetAdminByUsername retrieves an admin by their username
func GetAdminByUsername(db *mongo.Database, username string) (*model.Admin, error) {
	var admin model.Admin
	err := db.Collection("admin").FindOne(context.Background(), bson.M{"user_name": username}).Decode(&admin)
	if err == mongo.ErrNoDocuments {
		return nil, nil // Admin not found
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
