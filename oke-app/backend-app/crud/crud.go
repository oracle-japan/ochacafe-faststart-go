package crud

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/db"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/repo"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetItems(ctx *gin.Context) []repo.Items {
	dbInfo := db.GetDbInfo()
	items := []repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
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
	dbInfo := db.GetDbInfo()
	item := repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
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
	dbInfo := db.GetDbInfo()
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
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
	dbInfo := db.GetDbInfo()
	item := repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
	result := db.WithContext(ctx.Request.Context()).Delete(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}
