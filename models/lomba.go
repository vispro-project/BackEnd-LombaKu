package models

type Lomba struct {
	Id              int64  `gorm:"primaryKey" json:"id"`
	NamaLomba       string `gorm:"type:varchar(50)" json:"nama_lomba"`
	TanggalMulai    string `gorm:"type:varchar(15)" json:"tanggal_mulai"`
	TanggalBerakhir string `gorm:"type:varchar(15)" json:"tanggal_berakhir"`
	Deskripsi       string `gorm:"type:varchar(60)" json:"deskripsi_lomba"`
	Like            int64  `gorm:"type:int(7)" json:"like_lomba"` // Ganti dari LikeLomba menjadi Like
	Peserta         int64  `gorm:"type:int(7)" json:"peserta_lomba"`
	UserId          int64  `gorm:"index"`
}
