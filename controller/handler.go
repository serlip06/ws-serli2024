package controller

import (
	//"log"
	//"context"
	//"encoding/json"
	// "go.mongodb.org/mongo-driver/mongo"
	"time"
	"github.com/gofiber/fiber/v2"
	inimodel "github.com/serlip06/ujicobapackage/model"
	cek "github.com/serlip06/ujicobapackage/module"
	//"golang.org/x/crypto/bcrypt"
	//"go.mongodb.org/mongo-driver/bson"
)

var db = cek.MongoConnectdb("tesdb2024") // Ganti dengan nama database Anda


// Fungsi untuk registrasi pengguna baru langsung (tanpa acc)
func SignupHandler(c *fiber.Ctx) error {
	if c.Method() != "POST" {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
	}

	var registration inimodel.SignupRequest
	if err := c.BodyParser(&registration); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	// Membuat objek registrasi pengguna tanpa hash password
	pengguna := inimodel.Pengguna{
		Username:  registration.Username,
		Password:  registration.Password, // Menyimpan password dalam bentuk plaintext
		CreatedAt: time.Now(),
	}

	// Simpan data pengguna ke dalam koleksi "penggunas" (user resmi)
	err := cek.SavePengguna(pengguna, db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user")
	}

	// Return sukses setelah registrasi
	return c.Status(fiber.StatusCreated).JSON(map[string]string{"message": "Registration successful"})
}

// Fungsi untuk login pengguna
func SigninHandler(c *fiber.Ctx) error {
	if c.Method() != "POST" {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
	}

	var signinData inimodel.SigninRequest
	if err := c.BodyParser(&signinData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	// Panggil fungsi untuk mencari pengguna berdasarkan username
	pengguna, err := cek.FindPenggunaByUsername(signinData.Username, db)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Verifikasi password plaintext
	if pengguna.Password != signinData.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Return sukses setelah login berhasil
	return c.Status(fiber.StatusOK).JSON(inimodel.AccessResponse{
		Status:  "success",
		Message: "Login successful",
	})
}


// Fungsi untuk mengambil pengguna berdasarkan username
// func GetPenggunaByUsername(c *fiber.Ctx) error {
// 	// Ambil username dari parameter URL
// 	username := c.Params("username")

// 	// Panggil fungsi FindPenggunaByUsername untuk mencari pengguna
// 	pengguna, err := cek.FindPenggunaByUsername(username, db)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error": "User not found",
// 		})
// 	}

// 	// Mengembalikan data pengguna yang ditemukan
// 	return c.Status(fiber.StatusOK).JSON(pengguna)
// }

































// Fungsi untuk registrasi pengguna baru
// func SignupHandler(c *fiber.Ctx) error {
// 	if c.Method() != "POST" {
// 		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
// 	}

// 	var registration inimodel.UnverifiedUsers
// 	if err := c.BodyParser(&registration); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
// 	}

// 	// Hash password sebelum disimpan
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
// 	}
// 	registration.Password = string(hashedPassword)
// 	registration.SubmittedAt = time.Now()

// 	// Simpan data ke unverified users
// 	if err := cek.SaveUnverifiedUsers(registration, db); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save registration")
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(map[string]string{"message": "Registration successful"})
// }


// // Fungsi untuk login pengguna
// func SigninHandler(c *fiber.Ctx) error {
// 	if c.Method() != "POST" {
// 		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
// 	}

// 	var signinData inimodel.SigninRequest
// 	if err := c.BodyParser(&signinData); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
// 	}

// 	collection := db.Collection("penggunas")
// 	var pengguna inimodel.Pengguna
// 	err := collection.FindOne(context.Background(), bson.M{"username": signinData.Username}).Decode(&pengguna)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
// 	}

// 	if !verifyPassword(pengguna.Password, signinData.Password) {
// 		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(inimodel.AccessResponse{
// 		Status:  "success",
// 		Message: "Login successful",
// 	})
// }

// // Fungsi untuk memverifikasi password
// func verifyPassword(storedPassword string, inputPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
// 	return err == nil
// }

// func ConfirmRegistrationHandler(c *fiber.Ctx) error {
//     id := c.Params("id") // Ambil ID dari URL

//     if id == "" {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//             "error": "ID is required",
//         })
//     }

//     // Panggil fungsi ConfirmRegistration
//     unverifiedUser, pengguna, err := cek.ConfirmRegistration(id, db)
//     if err != nil {
//         log.Printf("Error in ConfirmRegistration: %v", err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to confirm registration",
//         })
//     }

//     // Hash password jika belum di-hash sebelumnya
//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pengguna.Password), bcrypt.DefaultCost)
//     if err != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to hash password",
//         })
//     }
//     pengguna.Password = string(hashedPassword)

//     // Simpan pengguna ke koleksi `penggunas`
//     collection := db.Collection("penggunas")
//     _, err = collection.InsertOne(context.Background(), pengguna)
//     if err != nil {
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//             "error": "Failed to save user to penggunas",
//         })
//     }

//     return c.Status(fiber.StatusOK).JSON(fiber.Map{
//         "message":         "Registration confirmed",
//         "unverified_user": unverifiedUser,
//         "pengguna":        pengguna,
//     })
// }
