package user

import (
	"fmt"
	"net/http"
)

type Service interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
}
