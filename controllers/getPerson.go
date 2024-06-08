package controllers

import (
	"infilon-task/internal"
	"infilon-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {

	id := c.Param("id")

	rows, err := internal.DB.Query(`
    SELECT 
      p.name,
      ph.number AS phone_number,
      a.city,
      a.state,
      a.street1,
      a.street2,
      a.zip_code
    FROM 
        person p
    JOIN 
        phone ph ON p.id = ph.person_id
    JOIN 
        address_join aj ON p.id = aj.person_id
    JOIN 
        address a ON aj.address_id = a.id
    WHERE 
        p.id = ?;
  `, id)
	internal.HandleErr(err, http.StatusInternalServerError, c)
	defer rows.Close()

	var person models.Person
	for rows.Next() {
		err = rows.Scan(&person.Name, &person.PhoneNumber, &person.City, &person.State, &person.Street1, &person.Street2, &person.ZipCode)
		internal.HandleErr(err, http.StatusInternalServerError, c)
	}

	if person.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, person)
}
