package crud

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/db"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/repo"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
)

func GetItems(ctx *gin.Context) []repo.Items {
	db := db.GetDBInfo()
	items := []repo.Items{}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	result := db.WithContext(ctx.Request.Context()).Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return items
}
func GetItemById(ctx *gin.Context, id string) repo.Items {
	db := db.GetDBInfo()
	item := repo.Items{}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	result := db.WithContext(ctx.Request.Context()).First(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return item
}

func UpdateItem(ctx *gin.Context, items repo.Items) int64 {
	db := db.GetDBInfo()
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	result := db.WithContext(ctx.Request.Context()).Save(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}

func DeleteItem(ctx *gin.Context, id string) int64 {
	db := db.GetDBInfo()
	item := repo.Items{}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	result := db.WithContext(ctx.Request.Context()).Delete(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}
