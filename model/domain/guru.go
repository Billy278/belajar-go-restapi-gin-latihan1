package domain

import "time"

type Guru struct {
	Id_guru   string    `gorm:"primarykey" json:"id_guru"`
	Name      string    `grom:"type:varchar(200)" json:"name"`
	Birth_day time.Time `grom:"type:date" json:"birth_day"`
	Married   bool      `grom:"type:bool" json:"married"`
	No_hp     string    `grom:"type:varchar(20)" json:"no_hp"`
}
