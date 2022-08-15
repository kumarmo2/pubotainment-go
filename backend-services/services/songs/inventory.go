package songs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	songServices "pubwebservice/business/songs"
	"pubwebservice/dtos/songs"
)

func AddSongToInventory(c *gin.Context) {
	var request songs.AddSongToInventory
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if request.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name cannot be empty"})
		c.Abort()
		return
	}

	companyId := c.Keys["companyId"].(int64)
	request.CompanyId = companyId

	inventoryId, err := songServices.AddSongToInventory(&request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": inventoryId})

	/*
	   * request validations
	   * check if name already exists in table songs_sorted_by_name
	       * if yes, return error
	       * if no, generate the model and insert in the main and sorted_by_name table.
	*/

}
