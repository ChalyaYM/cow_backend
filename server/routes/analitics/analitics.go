package analitics

import "github.com/gin-gonic/gin"

type Analitics struct {
}

func (b *Analitics) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/analitics")
	genotypedWriter := Genotyped{}
	checkMilksWriter := CheckMilks{}
	checkMilksWriter.WriteRoutes(apiGroup)
	genotypedWriter.WriteRoutes(apiGroup)
	totalGr := apiGroup.Group("/total")
	total := Total{}
	totalGr.GET("/:region_id/regionalStatistics/", total.RegionalStatistics())
	totalGr.GET("/:region_id/regionStatistics/", total.RegionStatistics())
	totalGr.GET("/:farm_id/farmStatistics/", total.FarmStatistics())
}
