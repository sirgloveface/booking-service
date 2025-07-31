package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sirgloveface/booking-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	dsn := os.Getenv("POSTGRES_URL")
	if dsn == "" {
		dsn = "host=dpg-d23sqqngi27c738d4otg-a.oregon-postgres.render.com user=curdoapp_user password=FSOXkbiebFZSnmp1XlZJ7DZ4MkQug2g9 dbname=curdoapp port=5432 sslmode=require"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&model.Booking{}); err != nil {
		log.Fatalf("failed to auto-migrate: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return db
}
