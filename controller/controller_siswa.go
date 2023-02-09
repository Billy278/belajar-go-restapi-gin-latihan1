package controller

import (
	"net/http"
	"strconv"

	"github.com/Billy278/belajar-go-restapi-gin-latihan1/app"
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/model/domain"
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/model/web"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAllSiswa(c *gin.Context) {
	siswa := []domain.Siswa{}
	app.DB.Find(&siswa)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   siswa,
	})
}
func FindByIdSiswa(c *gin.Context) {
	siswa := domain.Siswa{}
	id := c.Param("id")
	err := app.DB.First(&siswa, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "Not Found",
				Data:   err,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server Error",
				Data:   err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   siswa,
	})
}
func CreateSiswa(c *gin.Context) {
	siswa := domain.Siswa{}
	//dengan memanffatkan c.ShouldBindJSON
	//sudah otomatis request dari client di masukkan ke struct siswa
	err := c.ShouldBindJSON(&siswa)
	//
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	app.DB.Create(&siswa)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   siswa,
	})

}

func UpdateSiswa(c *gin.Context) {
	siswa := domain.Siswa{}
	id := c.Param("id")
	err := c.ShouldBindJSON(&siswa)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	rows := app.DB.Model(&siswa).Where("id=?", id).Updates(&siswa).RowsAffected
	if rows == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	siswa.Id, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   siswa,
	})
}

func DeleteSiswa(c *gin.Context) {
	siswa := domain.Siswa{}
	id := c.Param("id")

	rows := app.DB.Delete(&siswa, id).RowsAffected
	if rows == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		})
		return
	}
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	})
}
