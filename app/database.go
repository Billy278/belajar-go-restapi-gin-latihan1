package app

import (
	"github.com/Billy278/belajar-go-restapi-gin-latihan1/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/belajar_golang_restapi_gin_latihan1?parseTime=true"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Siswa{}, &domain.Guru{})
	DB = db

}
