package router

import "github.com/roskee/gotbot/entity"

// Handler is an abstract to hold a certain bot command and its implementation
type Handler struct {
	Name     string `json:"name"`
	Function func(update entity.Update)
}
