package model

type Menu struct {
	Id        string    `gorm:"primaryKey"`
	Menu_Text string    `gorm:"size:50; not null"`
	BaseModel BaseModel `gorm:"embedded"`
}

func (m Menu) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_menu"
}
