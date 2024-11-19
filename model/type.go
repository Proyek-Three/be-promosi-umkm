package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users represents the table Users
type Users struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama     string `json:"nama"`
	Role     string `json:"role"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password"`
	NoTelp   string `json:"no_telp"`
	Alamat   string `json:"alamat"`
}

// Product represents the table Produk
type Product struct {
	ID         int            `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaProduk string         `json:"nama_produk"`
	NamaToko   string         `json:"nama_toko"`
	Deskripsi  string         `json:"deskripsi"`
	Gambar     string         `json:"gambar"`
	Harga      int            `json:"harga"`
	Category   NamaCategory   `bson:"nama_category,omitempty" json:"nama_category,omitempty"`
}

// Category represents the table Categori
type Category struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaCategory string `json:"nama_category"`
}
