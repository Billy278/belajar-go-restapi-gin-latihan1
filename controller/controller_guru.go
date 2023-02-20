package controller

import (
	"net/http"

	"github.com/Billy278/belajar-go-restapi-gin-latihan1/app"
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/model/domain"
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/model/web"
	"github.com/gin-gonic/gin"
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
	//err := app.DB.First(&guru, id).Error
	sql := "Select id_guru,name,birth_day,married,no_hp from gurus where id_guru=?"
	err := app.DB.Raw(sql, id).Error
	row := app.DB.Raw(sql, id).Scan(&guru).RowsAffected

	if row == 0 {

		c.AbortWithStatusJSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   err,
		})
		return
	}

	guru.Id_guru = id
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
	//app.DB.Create(&guru)
	sql := "Insert into gurus(id_guru,name,birth_day,married,no_hp) Values(?,?,?,?,?)"

	app.DB.Exec(sql, guru.Id_guru, guru.Name, guru.Birth_day, guru.Married, guru.No_hp)
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
	//	rows := app.DB.Model(&guru).Where("id_guru=?", id).Updates(&guru).RowsAffected
	sql := "Update gurus set name=?,birth_day=?,married=?,no_hp=? Where id_guru=?"
	err = app.DB.Raw(sql, guru.Name, guru.Birth_day, guru.Married, guru.No_hp, id).Scan(&guru).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	guru.Id_guru = id
	c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func DeleteGuru(c *gin.Context) {
	id := c.Param("IdGuru")

	//rows := app.DB.Delete(&guru, id).RowsAffected
	sql := "Delete from gurus Where id_guru=?"
	rows := app.DB.Exec(sql, id).RowsAffected
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
