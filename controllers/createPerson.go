package controllers

import (
	"infilon-task/internal"
	"infilon-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {

	var data models.Person
	err := c.ShouldBindJSON(&data)
	internal.HandleErr(err, http.StatusBadRequest, c)

	// DB transaction
	tx, err := internal.DB.Begin()
	internal.HandleErr(err, http.StatusInternalServerError, c)

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			tx.Commit()
		}
	}()

	// Insert into person table
	insertPerson, err := tx.Prepare("INSERT INTO person (name) VALUES (?)")
	internal.HandleErr(err, http.StatusInternalServerError, c)
	defer insertPerson.Close()
	result, err := insertPerson.Exec(data.Name)
	internal.HandleErr(err, http.StatusInternalServerError, c)

	personID, err := result.LastInsertId()
	internal.HandleErr(err, http.StatusInternalServerError, c)

	// Insert into phone table
	insertPhone, err := tx.Prepare("INSERT INTO phone (person_id, number) VALUES (?, ?)")
	internal.HandleErr(err, http.StatusInternalServerError, c)
	defer insertPhone.Close()
	_, err = insertPhone.Exec(personID, data.PhoneNumber)
	internal.HandleErr(err, http.StatusInternalServerError, c)

	// Insert into address table
	insertAddress, err := tx.Prepare("INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)")
	internal.HandleErr(err, http.StatusInternalServerError, c)
	defer insertAddress.Close()
	result, err = insertAddress.Exec(data.City, data.State, data.Street1, data.Street2, data.ZipCode)
	internal.HandleErr(err, http.StatusInternalServerError, c)
	addressID, err := result.LastInsertId()
	internal.HandleErr(err, http.StatusInternalServerError, c)

	// Insert into address_join table
	insertAddressJoin, err := tx.Prepare("INSERT INTO address_join (person_id, address_id) VALUES (?, ?)")
	internal.HandleErr(err, http.StatusInternalServerError, c)
	defer insertAddressJoin.Close()

	_, err = insertAddressJoin.Exec(personID, addressID)
	internal.HandleErr(err, http.StatusInternalServerError, c)
}
