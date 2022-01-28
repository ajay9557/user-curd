package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
)

type Handler struct {
	handler services.User
}

func New(s services.User) Handler {
	return Handler{s}
}

func (h Handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	userId := params["id"]

	id, err := strconv.Atoi(userId)

	if err != nil {
		fmt.Println("Invalid User Id")
		return
	}
	user, err := h.handler.GetUserById(id)

	if err != nil {
		fmt.Println("User id not found")
		return
	}

	jsonData, _ := json.Marshal(user)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"data": {"user": %v}}`, string(jsonData))))
}

func (h Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	users, err := h.handler.GetUsers()

	if err != nil {
		log.Panic("Could not fetch users")
		return
	}

	jsonData, _ := json.Marshal(users)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"data": {"users": %v}}`, string(jsonData))))
}

func (h Handler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		log.Panic("cannot update user data")
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	convId, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Invalid Id")
	}

	_, err = h.handler.UpdateUser(convId, user)

	if err != nil {
		log.Panic("error in updation")
	}

	_, _ = w.Write([]byte(`{"data": "user updated successfully"}`))

}

func (h Handler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		log.Panic("Invalid Id")
		return
	}

	_, err = h.handler.DeleteUser(convId)

	if err != nil {
		log.Panic("error while deleting user")
		return
	}

	_, _ = w.Write([]byte(`{"data": "user deleted successfully"}`))

}

func (h Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		log.Panic("cannot create user data")

		return
	}

	_, err = h.handler.CreateUser(user)
	if err != nil {
		log.Panic("error while creating user")

		return
	}

	_, _ = w.Write([]byte(`{"data": "user created successfully"}`))

}
