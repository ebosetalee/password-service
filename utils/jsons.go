package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"net/http"

	"github.com/ebosetalee/password-service.git/types"
)

var ErrEmptyBody = errors.New("body must not be empty")

func ReadJSON(r *http.Request, dst interface{}) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	if r.Body == nil {
		return ErrEmptyBody
	}

	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %v", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return ErrEmptyBody

		default:
			return err
		}

	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v types.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = "Something went wrong try again later"
	}
	log.Println("status code:", status)
	// log.Println(err)
	response := types.Response{
		Code:    status,
		Message: errMsg,
	}

	WriteJSON(w, status, response)
}
