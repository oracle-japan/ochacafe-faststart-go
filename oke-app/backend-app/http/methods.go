package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/crud"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/repo"
)

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, crud.GetItems(c))
}

func GetItemById(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, crud.GetItemById(c, id))
}

func UpdateItem(c *gin.Context) {
	var item repo.Items
	if err := c.BindJSON(&item); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, crud.UpdateItem(c, item))
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, crud.DeleteItem(c, id))
}
