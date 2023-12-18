package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joeytatu/mongodb-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := []models.User{}

	if err := uc.session.DB("mongodb-golang").C("users").Find(nil).All(&users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error:", err)
		return
	}

	usersJson, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", usersJson)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}

	json.NewDecoder(r.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	uc.session.DB("mongodb-golang").C("users").Insert(user)

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	objId := bson.ObjectIdHex(id)
	user := models.User{}

	if err := uc.session.DB("mongodb-golang").C("users").FindId(objId).One(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	objId := bson.ObjectIdHex(id)
	user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	if err := uc.session.DB("mongodb-golang").C("users").UpdateId(objId, bson.M{"$set": user}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	updatedUserJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", updatedUserJson)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	objId := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongodb-golang").C("users").RemoveId(objId); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", objId, "\n")
}
