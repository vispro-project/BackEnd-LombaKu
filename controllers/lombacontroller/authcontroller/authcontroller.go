package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vispro-project/BackEnd-LombaKu.git/config"
	"github.com/vispro-project/BackEnd-LombaKu.git/helper"
	"github.com/vispro-project/BackEnd-LombaKu.git/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// cek user data user berdasarkan username

	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau Passworn salah"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau Passworn salah"}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	// generate token
	expTime := time.Now().Add(time.Minute * 60)
	claims := &config.JWTclaim{
		Username: user.Username,
		UserId:   user.Id, // Menambahkan UserId ke dalam claims
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	//deklarasi algoritma signIn

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//signed token

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke coockie

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
	return
}
func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJson(w, http.StatusOK, response)

}
func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token ke coockie

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "LogOut Berhasil"}
	helper.ResponseJson(w, http.StatusOK, response)
	return
}
