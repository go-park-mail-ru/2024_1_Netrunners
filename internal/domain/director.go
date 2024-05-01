package domain

import "time"

type DirectorToAdd struct {
	Name     string    `name:"name"`
	Avatar   string    `name:"avatar"`
	Birthday time.Time `name:"birthday"`
}
