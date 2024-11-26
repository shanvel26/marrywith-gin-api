package controllers

import (
	"marrywith-gin-api/models"
	"marrywith-gin-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonController struct {
	personService services.PersonService
}

func NewPersonController(client *mongo.Client) *PersonController {
	repo := models.NewPersonRepository(client)
	service := services.NewPersonService(repo)
	return &PersonController{personService: service}
}

func (pc *PersonController) CreatePerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.personService.CreatePerson(c, &person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (pc *PersonController) GetPersons(c *gin.Context) {
	persons, err := pc.personService.GetPersons(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}
