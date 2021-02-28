package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jorgeAM/goHexagonal/internal/creating"
	"github.com/jorgeAM/goHexagonal/internal/platform/bus/inmemory"
	"github.com/jorgeAM/goHexagonal/internal/platform/server"
	"github.com/jorgeAM/goHexagonal/internal/platform/storage/mysql"
)

const (
	host            = "localhost"
	port            = 3000
	shutdownTimeout = 10 * time.Second

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
	service := creating.NewCourseService(repository)

	commandBus := inmemory.NewCommadBus()

	createCourseCommandHandler := creating.NewCourseCommandHandler(service)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.NewServer(context.Background(), host, port, shutdownTimeout, commandBus)

	if err := srv.Run(ctx); err != nil {
		log.Fatalf("something got wrong when we try to run web server %v", err)
	}
}
