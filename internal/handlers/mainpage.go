package handlers

import "net/http"

type MainPageHandlers struct{}

func InitMainPageHandlers() *MainPageHandlers {
	return &MainPageHandlers{}
}

func (mainPageHandlers *MainPageHandlers) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
