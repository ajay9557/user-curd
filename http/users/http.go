package users

import (
	"user-curd/models"
	"user-curd/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserHandler struct {
	serv services.Services
}

func New(service services.Services) UserHandler {
	return UserHandler{serv: service}
}

func (u UserHandler) PostUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("invalid body"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user.Id == 0 {
		_, _ = rw.Write([]byte("Id shouldn't be zero"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok, err := u.serv.EmailValidation(user.Email)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("error generated"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		_, _ = rw.Write([]byte("email already present - could not create user"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		err = u.serv.InsertUserDetails(user)
		if err != nil {
			fmt.Println(err)
			_, _ = rw.Write([]byte("Database error"))
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("User created"))
	}
}

func (u UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	allUsers, err := u.serv.FetchAllUserDetails()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("error generated"))
	} else {
		b, err := json.Marshal(allUsers)
		if err != nil {
			fmt.Println(err)
			return
		}
		rw.WriteHeader(http.StatusOK)
		_, err = rw.Write(b)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (u UserHandler) GetUserById(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		_, _ = rw.Write([]byte("invalid parameter id"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := u.serv.FetchUserDetailsById(int_id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal error"))
		return
	} else {
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			rw.Write([]byte("error in json"))
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		_, err = rw.Write(b)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func (u UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var user *models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("invalid body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Id == 0 {
		_, _ = rw.Write([]byte("Id shouldn't be zero"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok, err := u.serv.EmailValidation(user.Email)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("error generated"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		_, _ = rw.Write([]byte("email already present - could not create user"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		err = u.serv.UpdateUserDetails(*user)
		if err != nil {
			fmt.Println(err)
			_, _ = rw.Write([]byte("Database error"))
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("User updated"))
	}
}

func (u UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	int_id, _ := strconv.Atoi(id)
	if int_id == 0 {
		_, _ = rw.Write([]byte("Id shouldn't be zero"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err := u.serv.DeleteUserDetailsById(int_id)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("error generated"))
		return
	} else {
		rw.WriteHeader(http.StatusOK)
		_, err = rw.Write([]byte("User deleted successfully"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
