package handlers

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message error) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message.Error()))
	if err != nil {
		fmt.Printf("creating user error: %v", err)
	}
	return
}
