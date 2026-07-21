package main

import (
	"book_recsys/models"
	"book_recsys/server"
	"book_recsys/settings"
	"log"
)

func main() {

	cfg, err := settings.NewConfig()
	if err != nil {
		panic("Erro ao carregar configurações: " + err.Error())
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		panic("Erro ao criar servidor: " + err.Error())
	}

	log.Println("Fazendo migrations....")
	srv.DBRelClient.Con.AutoMigrate(&models.Book{})

	log.Println("Subindo o servidor...")
	if err := srv.Start(); err != nil {
		panic("Erro ao iniciar servidor: " + err.Error())
	}

}
