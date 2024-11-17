package app

import (
	"log"
	"net/http"
	"posts/internal/config"

	"posts/internal/postgres"
	"posts/internal/service"
	transport "posts/internal/transport/http"
	datab "posts/pkg/db"
)

func RunServer() {

	db, err := datab.NewPostgresDb(datab.InfoDatabase{
		Host:     config.HOST,
		Port:     config.PORT,
		User:     "postgres",
		Dbname:   "posters",
		Password: config.DB_PASS,
		Sslmode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewPosts(db)
	serv := service.NewPosts(repo)
	handl := transport.NewHandler(serv)
	r := handl.InitRouter()
	log.Println("server run on the 8080 port")
	http.ListenAndServe(config.ADDR, r)
}
