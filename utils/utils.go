package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(payload); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return fmt.Errorf("malformed JSON at position %d", syntaxErr.Offset)
		case errors.As(err, &unmarshalTypeErr):
			log.Println("unmarshal type error", unmarshalTypeErr.Value)
			return fmt.Errorf("invalid value for field '%s', expected type '%s', received type '%s'",
				unmarshalTypeErr.Field,
				unmarshalTypeErr.Type,
				unmarshalTypeErr.Value,
			)
		case errors.Is(err, io.EOF):
			return fmt.Errorf("empty request body")
		case strings.Contains(err.Error(), "json: unknown field"):
			fieldName := extractUnknownFieldName(err.Error())
			return fmt.Errorf("unknown field '%s' in request body", fieldName)
		default:
			return fmt.Errorf("invalid JSON payload")
		}
	}

	return nil
}

func extractUnknownFieldName(errMsg string) string {
	start := strings.Index(errMsg, "\"") + 1
	end := strings.LastIndex(errMsg, "\"")
	if start > 0 && end > start {
		return errMsg[start:end]
	}
	return "unknown"
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
