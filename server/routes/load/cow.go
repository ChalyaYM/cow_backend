package load

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const COW_CSV_PATH = "./csv/cows/"

var cowUniqueIndex uint64 = 0

func (l *Load) Cow() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		csv, ok := form.File["csv"]
		if !ok {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		uploadedName := COW_CSV_PATH + "cow_" + now.Format(time.Stamp) + "_" + strconv.FormatUint(cowUniqueIndex, 10) + ".csv"
		if err := c.SaveUploadedFile(csv[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		cowUniqueIndex++
		c.JSON(200, "OK")
	}
}
