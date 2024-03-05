package handlers

import (
	"fmt"
	"net/http"
)

type MainPageHandlers struct{}

func InitMainPageHandlers() *MainPageHandlers {
	return &MainPageHandlers{}
}

func (mainPageHandlers *MainPageHandlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
}
