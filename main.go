package main

import (
	"car/orms"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// repo := repositories.NewCarRepositoryMock()
	// serv := services.NewCarService(repo)
	// h := handlers.Newhandler(serv)

	// g := gin.Default()
	// g.POST("/car", h.Create)
	// g.GET("/car", h.FindMany)
	// g.GET("/car/:id", h.FindByID)
	// g.PUT("/car/:id", h.Update)
	// g.DELETE("/car/:id", h.Delete)

	// g.Run(":8080")

	dsn := "host=localhost user=postgres password=password dbname=integration port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&orms.Car{})
}
