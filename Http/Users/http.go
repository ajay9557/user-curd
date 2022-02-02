package Users

import (
	"Icrud/Services"
	"Icrud/TModels"
	"encoding/json"
	_ "fmt"
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
		w.WriteHeader(http.StatusInternalServerError)
		herr := TModels.ErrorResponse{StCode: http.StatusInternalServerError, Errmessage: "Bad Request. Invalid User id"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	if iid < 1 {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: "Bad Request. Id should be greater than 0"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	UserWithId, er := srvhdlr.Sev.UserById(iid)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: er.Error()}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	responseData := TModels.Response{
		Data: struct {
			User TModels.User `json: "user"`
		}{
			User: UserWithId,
		},
		Message:    "Successfully User Retrieved",
		StatusCode: 200,
	}

	jsonData, er := json.Marshal(responseData)
	if er == nil {
		_, _ = w.Write(jsonData)
	}
}

func (srvhdlr Handler) GetUsers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	AllUsers, err := srvhdlr.Sev.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: err.Error()}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	responseData := TModels.Response{
		Data: struct {
			Users []TModels.User `json: "users"`
		}{
			Users: AllUsers,
		},
		Message:    "Successfully All Users are Retrieved",
		StatusCode: 200,
	}

	jsonData, er := json.Marshal(responseData)
	if er == nil {
		_, _ = w.Write(jsonData)
	}

}

func (srvhdlr Handler) InsertUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u TModels.User
	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	herr := TModels.ErrorResponse{http.StatusBadRequest, "ProductNotAvailable"}
	// 	res, er := json.Marshal(herr)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	if er == nil {
	// 		w.Write(res)
	// 	}
	// 	return
	// }
	// err = json.Unmarshal(body, &u)
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil || reflect.DeepEqual(u, TModels.User{}) {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: "Can't parse the given data"}
		jsonData, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(jsonData)
		}
		return
	}

	// ok, err := srvhdlr.Sev.EmailValidation(u.Email)
	// if err != nil {
	// 	w.Write([]byte(`{data: user inserted failed}`))
	// } else if !ok {
	// 	w.Write([]byte(`{data: email already present}`))
	// } else {

	createdUser, err := srvhdlr.Sev.InsertUser(u)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: err.Error()}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	responseData := TModels.Response{
		Data: struct {
			User TModels.User `json: "user"`
		}{
			User: createdUser,
		},
		Message:    "Successfully User Added",
		StatusCode: 200,
	}

	jsonData, er := json.Marshal(responseData)
	if er == nil {
		_, _ = w.Write(jsonData)
	}

	// w.Write([]byte(`{data: user inserted successfully}`))
	// }

}

func (srvhdlr Handler) DeleteUserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	v := params["id"]

	iid, er := strconv.Atoi(v)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		herr := TModels.ErrorResponse{StCode: http.StatusInternalServerError, Errmessage: "Bad Request. Invalid User id"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	if iid < 1 {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: "Bad Request. Id should be greater than 0"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	_, er = srvhdlr.Sev.DeleteUserById(iid)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: er.Error()}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	responseData := TModels.Response{
		Message:    "Successfully User Deleted",
		StatusCode: 200,
	}

	jsonData, er := json.Marshal(responseData)
	if er == nil {
		_, _ = w.Write(jsonData)
	}

	// w.Write([]byte(`{data: user deleted successfully}`))

}

// func (srvhdlr Handler) UpdateUserById(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(req)
// 	v := params["id"]

// 	iid, er := strconv.Atoi(v)
// 	if er != nil {
// 		herr := TModels.ErrorResponse{http.StatusInternalServerError, "Invalid Id"}
// 		w.WriteHeader(http.StatusInternalServerError)
// 		res, er := json.Marshal(herr)
// 		if er == nil {
// 			w.Write(res)
// 		}
// 		return
// 	}

// 	var u TModels.User
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		herr := TModels.ErrorResponse{http.StatusBadRequest, "ProductNotAvailable"}
// 		res, er := json.Marshal(herr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		if er == nil {
// 			w.Write(res)
// 		}
// 		return
// 	}
// 	err = json.Unmarshal(body, &u)
// 	if (err != nil || reflect.DeepEqual(u, TModels.User{})) {
// 		herr := TModels.ErrorResponse{http.StatusBadRequest, "ProductNotAvailable"}
// 		res, er := json.Marshal(herr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		if er == nil {
// 			w.Write(res)
// 		}
// 		return
// 	}

// 	_, err = srvhdlr.Sev.UpdateUserById(u, iid)
// 	if err != nil {
// 		herr := TModels.ErrorResponse{http.StatusBadRequest, "ProductNotAvailable"}
// 		res, er := json.Marshal(herr)
// 		w.WriteHeader(http.StatusBadRequest)
// 		if er == nil {
// 			w.Write(res)
// 		}
// 		return
// 	}

// 	w.Write([]byte(`{data: user updated successfully}`))

// }

func (srvhdlr Handler) UpdateUserById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	var u TModels.User
	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil || reflect.DeepEqual(u, TModels.User{}) {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: "Can't parse the given data"}
		jsonData, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(jsonData)
		}
		return
	}

	params := mux.Vars(req)
	v := params["id"]

	iid, er := strconv.Atoi(v)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		herr := TModels.ErrorResponse{StCode: http.StatusInternalServerError, Errmessage: "Bad Request. Invalid User id"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	if iid < 1 {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: "Bad Request. Id should be greater than 0"}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	updatedUser, er := srvhdlr.Sev.UpdateUserById(u, iid)
	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		herr := TModels.ErrorResponse{StCode: http.StatusBadRequest, Errmessage: er.Error()}
		res, er := json.Marshal(herr)
		if er == nil {
			_, _ = w.Write(res)
		}
		return
	}

	responseData := TModels.Response{
		Data: struct {
			User TModels.User `json :user`
		}{
			updatedUser,
		},
		StatusCode: 200,
		Message:    "Successfully User Updated",
	}

	jsonData, er := json.Marshal(responseData)
	if er == nil {
		_, _ = w.Write(jsonData)
	}

}
