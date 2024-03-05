package handlers

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Status int `json:"status"`
}

type ErrorResponse struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func WriteSuccess(w http.ResponseWriter) error {
	response := SuccessResponse{
		Status: http.StatusOK,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}

func WriteError(w http.ResponseWriter, statusCode int, message error) error {
	response := ErrorResponse{
		Status: statusCode,
		Err:    message.Error(),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil
}
