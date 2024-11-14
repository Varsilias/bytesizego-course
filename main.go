package main

import (
	"fmt"
	"github.com/Varsilias/bytesizego-course/internal/config"
	"github.com/Varsilias/bytesizego-course/internal/db"
	"github.com/Varsilias/bytesizego-course/internal/todo"
	"github.com/Varsilias/bytesizego-course/internal/transport"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	config.LoadConfig()
	port, err := strconv.Atoi(viper.Get("db_port").(string))
	host := viper.Get("db_host").(string)
	password := viper.Get("db_password").(string)
	dbName := viper.Get("db_name").(string)
	dbUser := viper.Get("db_user").(string)

	if err != nil {
		log.Fatal(fmt.Errorf("error parsing port to int: %v", err))
	}

	conn, err := db.New(dbUser, password, dbName, host, port)

	if err != nil {
		log.Fatal(err)
	}

	svc := todo.NewService(conn)
	server := transport.NewServer(svc)

	if err := server.ServeHTTP(); err != nil {
		log.Fatal(err)
	}

}
