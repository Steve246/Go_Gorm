package model

type Menu_Price struct {
	Id string `gorm:"primaryKey"`
	// Menu      Menu
	Price     int
	BaseModel BaseModel `gorm:"embedded"`
}

func (mp Menu_Price) TableName() string {
	//ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "m_menu_price"
}
