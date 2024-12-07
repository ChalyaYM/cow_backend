package admin

import (
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

func (s *Admin) WriteRoutes(rg *gin.RouterGroup) {
	apiGroup := rg.Group("/admin")
	apiGroup.GET("/cowTable", s.CheckCowTable())
	apiGroup.GET("/checkUsers", s.CheckUsersTable())
	apiGroup.GET("/createUser", s.CreateUser())
	apiGroup.GET("/checkHoldings", s.CheckHozTable(1))
	apiGroup.GET("/createHolding", s.CreateHolding())
	apiGroup.GET("/checkHozs", s.CheckHozTable(2))
	apiGroup.GET("/checkFarms", s.CheckHozTable(3))
	apiGroup.POST("/approveCows", s.ApproveCows())
	apiGroup.POST("/newUser", s.NewUser())
	apiGroup.POST("/newHolding", s.NewHolding())
	apiGroup.GET("", s.Index())
}
