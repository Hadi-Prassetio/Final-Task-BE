package handlers

import (
	profiledto "_waysbeans_/dto/profile"
	dto "_waysbeans_/dto/result"
	"_waysbeans_/models"
	"_waysbeans_/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// . Declare handlerProfile struct
type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

// . HandlerProfile function
func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

// . FindProfiles Method
func (h *handlerProfile) FindProfiles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	profiles, err := h.ProfileRepository.FindProfiles()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: profiles}

	json.NewEncoder(w).Encode(response)
}

// . GetProfile method
func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)}
	json.NewEncoder(w).Encode(response)
}

// . Create Profile
func (h *handlerProfile) CreateProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := profiledto.CreateProfileRequest{
		Image:    filename,
		Phone:    r.FormValue("phone"),
		Address:  r.FormValue("address"),
		PostCode: r.FormValue("post_code"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	profile := models.Profile{
		Image:    request.Image,
		Phone:    request.Phone,
		Address:  request.Address,
		PostCode: request.PostCode,
		UserID:   userId,
	}

	data, err := h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(data)}
	json.NewEncoder(w).Encode(response)
}

// . convertResponseProfile function
func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{

		ID:       u.ID,
		Phone:    u.Phone,
		Address:  u.Address,
		PostCode: u.PostCode,
		Image:    u.Image,
		UserID:   u.UserID,
		User:     u.User,
	}
}
