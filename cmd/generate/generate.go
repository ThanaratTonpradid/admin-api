package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"mini-api/config"
)

func main() {
	config.AutoReadConfig("local.env")
	workerConfig := config.NewAPIConfig()

	g := gen.NewGenerator(gen.Config{
		OutPath:       "./query",
		FieldNullable: true,
		FieldSignable: true,
	})

	db, _ := gorm.Open(mysql.Open(workerConfig.MySQLDSN))
	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
}
