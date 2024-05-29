package domain

import "time"

//easyjson:json
type DirectorToAdd struct {
	Name     string    `name:"name"`
	Avatar   string    `name:"avatar"`
	Birthday time.Time `name:"birthday"`
}
