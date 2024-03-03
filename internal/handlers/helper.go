package handlers

import (
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message error) error {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message.Error()))
	if err != nil {
		return err
	}
	return nil
}
