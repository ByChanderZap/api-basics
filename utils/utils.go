package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	if code > 499 {
		log.Println("Responded with 5XX error: ", err.Error())
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	WriteJSON(w, code, errResponse{
		Error: err.Error(),
	})
}

// func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
// 	data, err := json.Marshal(payload)
// 	if err != nil {
// 		log.Printf("Failed to marshall json response %v\n", err)
// 		w.WriteHeader(500)
// 		return
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(data)
// }

var (
	Validate   = validator.New()
	Translator ut.Translator
)

func InitValidator() {
	eng := en.New()
	uni := ut.New(eng, eng)
	Translator, _ = uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(Validate, Translator)
}

func ParseUUIDParam(r *http.Request, param string) (uuid.UUID, error) {
	pId := chi.URLParam(r, param)
	parsedId, err := uuid.Parse(pId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid %s id", param)
	}
	return parsedId, nil
}
