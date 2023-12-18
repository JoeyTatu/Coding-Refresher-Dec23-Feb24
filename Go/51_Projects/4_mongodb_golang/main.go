package main

import (
	"log"
	"net/http"

	"github.com/joeytatu/mongodb-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users", uc.GetAllUsers)
	r.POST("users/user", uc.CreateUser)
	r.GET("/users/user/:id", uc.GetUserById)
	r.PUT("/users/user/:id", uc.UpdateUser)   // same as PATCH
	r.PATCH("/users/user/:id", uc.UpdateUser) // same as PUT
	r.DELETE("/users/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017/mongodb-users")
	if err != nil {
		log.Fatal("Error:", err)
		return nil
	}
	return s
}
