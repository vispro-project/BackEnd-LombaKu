package lombacontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/vispro-project/BackEnd-LombaKu.git/config"
	"github.com/vispro-project/BackEnd-LombaKu.git/helper"
	"github.com/vispro-project/BackEnd-LombaKu.git/models"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var lomba []models.Lomba

	if err := models.DB.Find(&lomba).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, lomba)
}
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}
	var lomba models.Lomba
	if err := models.DB.First(&lomba, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Lomba Tidak Adaa")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, lomba)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var lomba models.Lomba

	// Mengambil nilai cookie dengan nama "token"
	cookie, err := r.Cookie("token")
	if err != nil {
		ResponseError(w, http.StatusUnauthorized, "Cookie token tidak ditemukan: "+err.Error())
		return
	}

	// Mendekode token untuk mendapatkan informasi pengguna
	token, err := jwt.ParseWithClaims(cookie.Value, &config.JWTclaim{}, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if err != nil {
		ResponseError(w, http.StatusUnauthorized, "Gagal mendekode token: "+err.Error())
		return
	}

	// Mengakses informasi pengguna dari claims
	claims, ok := token.Claims.(*config.JWTclaim)
	if !ok {
		ResponseError(w, http.StatusUnauthorized, "Gagal mendapatkan claims dari token")
		return
	}

	// Set nilai UserID di model Lomba
	lomba.UserId = claims.UserId

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lomba); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&lomba).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, lomba)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var lomba models.Lomba
	if err := models.DB.First(&lomba, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, " lomba trre")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err := models.DB.Delete(&lomba).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, map[string]string{"message": "Lomba deleted successfully"})
}

func SearchLombaByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		ResponseError(w, http.StatusBadRequest, "Name parameter is required")
		return
	}

	var lomba []models.Lomba
	if err := models.DB.Where("nama_lomba LIKE ?", "%"+name+"%").Find(&lomba).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Lomba Tidak Adaaa")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ResponseJson(w, http.StatusOK, lomba)
}
