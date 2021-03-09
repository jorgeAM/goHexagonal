package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"

	"github.com/jorgeAM/goHexagonal/internal/creating"
	"github.com/jorgeAM/goHexagonal/internal/platform/bus/inmemory"
	"github.com/jorgeAM/goHexagonal/internal/platform/server"
	"github.com/jorgeAM/goHexagonal/internal/platform/storage/mysql"
)

const (
	shutdownTimeout = 10 * time.Second
	dbTimeout       = 5 * time.Second
)

func main() {
	var (
		host   = os.Getenv("HOST")
		port   = os.Getenv("PORT")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	)

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		log.Fatalf("something got wrong when we try to connect to database %v", err)
	}

	repository := mysql.NewCourseRepository(db, dbTimeout)
	service := creating.NewCourseService(repository)

	commandBus := inmemory.NewCommadBus()

	createCourseCommandHandler := creating.NewCourseCommandHandler(service)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	portUint, _ := strconv.Atoi(port)

	ctx, srv := server.NewServer(context.Background(), host, uint(portUint), shutdownTimeout, commandBus)

	if err := srv.Run(ctx); err != nil {
		log.Fatalf("something got wrong when we try to run web server %v", err)
	}
}
