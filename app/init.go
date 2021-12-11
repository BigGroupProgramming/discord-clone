package app

import (
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Server struct {
	DB *gorm.DB
}
func Init() (Server, error) {
	var (
		s Server
		err error
	)

	s.DB, err = initDB()
	if err != nil {
		return Server{}, fmt.Errorf("failed to initDB | %v", err)
	}

	return s, nil
}

func initDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening db | %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db | %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to db | %v", err)
	}

	return gormDB, nil
}