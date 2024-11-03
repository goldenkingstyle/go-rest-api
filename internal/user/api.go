package user

import "net/http"

func RegisterHandlers(m *http.ServeMux) {

	service := NewService()

	m.HandleFunc("GET /user/{id}", service.GetUser)

}
