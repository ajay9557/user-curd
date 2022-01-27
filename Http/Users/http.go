package Users

import (
	"Icrud/Services"
	"Icrud/TModels"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	Sev Services.ISUser
}

func (srvhdlr Handler) UserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	v := params["id"]

	iid, er := strconv.Atoi(v)
	if er != nil {
		herr := TModels.RetErr{http.StatusInternalServerError, "Invalid Id"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(herr)
		if er == nil {
			w.Write(res)
		}
		return
	}

	UserWithId, err := srvhdlr.Sev.UserById(iid)
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	res, er := json.Marshal(UserWithId)
	if er == nil {
		w.Write(res)
	}
}

func (srvhdlr Handler) GetUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	AllUsers, err := srvhdlr.Sev.GetUsers()
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	res, er := json.Marshal(AllUsers)
	if er == nil {
		w.Write(res)
	}

}

func (srvhdlr Handler) InsertUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u TModels.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}
	err = json.Unmarshal(body, &u)
	if err != nil || reflect.DeepEqual(u, TModels.User{}) {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	// ok, err := srvhdlr.Sev.EmailValidation(u.Email)
	// if err != nil {
	// 	w.Write([]byte(`{data: user inserted failed}`))
	// } else if !ok {
	// 	w.Write([]byte(`{data: email already present}`))
	// } else {

	_, err = srvhdlr.Sev.InsertUser(u)

	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	w.Write([]byte(`{data: user inserted successfully}`))
	// }

}

func (srvhdlr Handler) DeleteUserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	v := params["id"]

	iid, er := strconv.Atoi(v)
	if er != nil {
		herr := TModels.RetErr{http.StatusInternalServerError, "Invalid Id"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(herr)
		if er == nil {
			w.Write(res)
		}
		return
	}

	_, err := srvhdlr.Sev.DeleteUserById(iid)
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	w.Write([]byte(`{data: user deleted successfully}`))

}

func (srvhdlr Handler) UpdateUserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	v := params["id"]

	iid, er := strconv.Atoi(v)
	if er != nil {
		herr := TModels.RetErr{http.StatusInternalServerError, "Invalid Id"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(herr)
		if er == nil {
			w.Write(res)
		}
		return
	}

	var u TModels.User
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}
	err = json.Unmarshal(body, &u)
	if (err != nil || reflect.DeepEqual(u, TModels.User{})) {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	_, err = srvhdlr.Sev.UpdateUserById(u, iid)
	if err != nil {
		herr := TModels.RetErr{http.StatusBadRequest, "ProductNotAvailable"}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}

	w.Write([]byte(`{data: user updated successfully}`))

}
