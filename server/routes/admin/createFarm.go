package admin

import (
	"cow_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Admin) CreateFarm() func(*gin.Context) {
	return func(c *gin.Context) {
		db := models.GetDb()
		regions := []models.Region{}
		districts := []models.District{}
		hozs := []models.Farm{}
		db.Find(&regions)
		db.Find(&districts)
		db.Where("type = 2").Find(&hozs)
		c.HTML(http.StatusOK, "AdminCreateFarmPage.tmpl", gin.H{
			"title":     "Создание хозяйства",
			"holds":     hozs,
			"regions":   regions,
			"districts": districts})
	}
}
