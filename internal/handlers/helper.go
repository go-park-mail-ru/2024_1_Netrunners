package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
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

func WriteError(w http.ResponseWriter, err error) error {
	fmt.Println(err)
	statusCode, err := myerrors.ParseError(err)

	response := ErrorResponse{
		Status: statusCode,
		Err:    err.Error(),
	}

	jsonResponse, err := json.Marshal(response)
	fmt.Println(err)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
