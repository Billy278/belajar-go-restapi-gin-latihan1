package main

import (
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/app"
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	app.NewDB()
	router.GET("/api/siswa", controller.FindAllSiswa)
	router.GET("/api/siswa/:id", controller.FindByIdSiswa)
	router.POST("/api/siswa", controller.CreateSiswa)
	router.PUT("/api/siswa/:id", controller.UpdateSiswa)
	router.DELETE("/api/siswa/:id", controller.DeleteSiswa)

	router.GET("/api/guru", controller.FindAllGuru)
	router.POST("/api/guru", controller.CreateGuru)
	router.GET("/api/guru/:IdGuru", controller.FindByIdGuru)
	router.PUT("/api/guru/:IdGuru", controller.UpdateGuru)
	router.DELETE("/api/guru/:IdGuru", controller.DeleteGuru)

	router.Run(":9000")
}
