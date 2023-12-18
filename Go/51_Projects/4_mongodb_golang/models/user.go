package models

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Sex    string        `json:"sex" bson:"sex"`
	Gender string        `json:"gender" bson:"gender"`
	DOB    time.Time     `json:"dob" bson:"dob"`
	Age    int16         `json:"age" bson:"age"`
}

// START - AI generated
func (u *User) CalculateAge() {
	now := time.Now()
	age := now.Year() - u.DOB.Year()

	if now.YearDay() < u.DOB.YearDay() {
		age-- // Check if birthday has already happened
	}

	u.Age = int16(age)
}

func (u *User) MarshallJson() ([]byte, error) {
	u.CalculateAge()
	type Alias User
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}

// END - AI generated
