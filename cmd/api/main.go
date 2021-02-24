package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jorgeAM/goHexagonal/internal/platform/server"
	"github.com/jorgeAM/goHexagonal/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 3000

	dbUser = "root"
	dbPass = "123456"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "courses"
)

func main() {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		log.Fatalf("something got wrong when we try to connect to database %v", err)
	}

	repository := mysql.NewCourseRepository(db)
	srv := server.NewServer(host, port, repository)

	if err := srv.Run(); err != nil {
		log.Fatalf("something got wrong when we try to run web server %v", err)
	}
}
