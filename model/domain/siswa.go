package domain

type Siswa struct {
	Id      int    `gorm:"primarykey" json:"id"`
	Name    string `grom:"type:varchar(200)" json:"name"`
	Jurusan string `grom:"type:varchar(200)" json:"jurusan"`
	Stambuk string `grom:"type:varchar(200)" json:"stambuk"`
}
