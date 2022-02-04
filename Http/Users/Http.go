package Users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-curd/model"
	service "user-curd/services"

	"github.com/gorilla/mux"
)

type Handler struct {
	S service.User
}

func (serv Handler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	v := params["id"]

	id, err := strconv.Atoi(v)

	if err != nil {
		w.Write([]byte("invalid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usr, err := serv.S.SearchByUserId(id)

	if err != nil {
		w.Write([]byte("id not found"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(usr)
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (serv Handler) Create(w http.ResponseWriter, r *http.Request) {
	var users model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &users)
	if err != nil {
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usr, err := serv.S.InsertUserDetails(&users)
	res, _ := json.Marshal(usr)
	if err != nil {
		_, _ = w.Write([]byte("could not create User"))
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		w.Write(res)
	}
}

func (serv Handler) DeleteId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response model.Response
	params := mux.Vars(r)

	v := params["id"]

	id, err := strconv.Atoi(v)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.StatusCode = http.StatusBadRequest
		response.Message = "invalid id"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		return
	}
	err = serv.S.DeleteByUserId(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.StatusCode = http.StatusInternalServerError
		response.Message = err.Error()
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		return
	}
	response.StatusCode = http.StatusOK
	response.Message = "user deleted"
	response.Data = nil
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
	w.WriteHeader(response.StatusCode)

}
func (serv Handler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var response model.Response
	body, _ := ioutil.ReadAll(r.Body)
	var usr model.User
	err := json.Unmarshal(body, &usr)

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "invalid body"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = rw.Write(b)
		rw.WriteHeader(response.StatusCode)
		return
	}
	user, err := serv.S.UpdateByUserId(&usr)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "could not update user"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = rw.Write(b)
		rw.WriteHeader(response.StatusCode)
		return
	}
	response.StatusCode = http.StatusOK
	response.Message = "user updated"
	response.Data = user
	b, _ := json.Marshal(response)
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func (serv Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	usr, err := serv.S.SearchAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Retrieving Failed."))
	}
	res, err := json.Marshal(usr)
	if err != nil {
		_, _ = w.Write([]byte("could not get User"))
		w.WriteHeader(http.StatusInternalServerError)
	} else {

		w.Write(res)
	}
}
