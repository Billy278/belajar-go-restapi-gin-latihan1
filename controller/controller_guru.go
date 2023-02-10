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

func FindAllGuru(c *gin.Context) {
	guru := []domain.Guru{}
	app.DB.Find(&guru)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}
func FindByIdGuru(c *gin.Context) {
	id := c.Param("IdGuru")
	guru := domain.Guru{}
	err := app.DB.First(&guru, id).Error
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
	guru.Id_guru, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func CreateGuru(c *gin.Context) {
	guru := domain.Guru{}
	err := c.ShouldBindJSON(&guru)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	app.DB.Create(&guru)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func UpdateGuru(c *gin.Context) {
	id := c.Param("IdGuru")
	guru := domain.Guru{}
	err := c.ShouldBindJSON(&guru)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	rows := app.DB.Model(&guru).Where("id_guru=?", id).Updates(&guru).RowsAffected
	if rows == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	guru.Id_guru, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func DeleteGuru(c *gin.Context) {
	guru := domain.Guru{}
	id := c.Param("IdGuru")

	rows := app.DB.Delete(&guru, id).RowsAffected
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
