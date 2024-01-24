package models

type User struct {
	Id          int64   `gorm:"primaryKey" json:"id"`
	NamaLengkap string  `gorm:"varchar(50)" json:"nama_lengkap"`
	Username    string  `gorm:"type:varchar(50);unique" json:"username"`
	Password    string  `gorm:"type:varchar(60)" json:"password"`
	Email       string  `gorm:"type:varchar(100);unique" json:"email"`
	Lombas      []Lomba `gorm:"foreignKey:UserId"`
}
