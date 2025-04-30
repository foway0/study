package infra

import (
	"database/sql"
	"fmt"
	"github.com/foway0/study/go-grpc/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitMySQL(config config.Config) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Mysql.User,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database),
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("ping database...")
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return db
}
