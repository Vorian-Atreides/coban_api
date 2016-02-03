package databases

func createTables() {
	DB.AutoMigrate(&Address{})
	DB.AutoMigrate(&Company{})
	DB.AutoMigrate(&Device{})
	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&TransportType{})
	DB.AutoMigrate(&Station{})
	DB.AutoMigrate(&TransportHistory{})
}

func Migrate()  {
	createTables()
}