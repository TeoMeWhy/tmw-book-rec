package server

import (
	booksmgmtctr "book_recsys/controllers/books_mgmt_ctr"
	reldb "book_recsys/repositories/reldb"
	booksmgmtservice "book_recsys/services/books_mgmt_service"
	"book_recsys/settings"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	DBRelClient     *reldb.DBClient
	BookMGMTService *booksmgmtservice.BookService

	BookMGMTController *booksmgmtctr.BookMgmtController
	URI                string
	Port               string
	API                *fiber.App
}

func (s *Server) Start() error {
	s.API = fiber.New()

	s.API.Get("/books", s.BookMGMTController.GetAllBooks)
	s.API.Get("/books/search", s.BookMGMTController.GetBookByQuery)
	s.API.Get("/books/:id", s.BookMGMTController.GetBookByID)
	s.API.Post("/books", s.BookMGMTController.RegisterBook)
	s.API.Put("/books/:id", s.BookMGMTController.UpdateBook)

	s.API.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "OK"})
	})

	addr := s.URI + ":" + s.Port
	return s.API.Listen(addr)
}

func NewServer(cfg *settings.Config) (*Server, error) {

	DBClient, err := reldb.NewDBClient(cfg)
	if err != nil {
		log.Println("Erro ao criar DBClient:", err)
		return nil, err
	}

	BookMGMTService := booksmgmtservice.NewBookService(DBClient)

	ctrBookMgmt := booksmgmtctr.NewBookMgmtController(BookMGMTService)

	server := &Server{
		URI:                cfg.ServerURI,
		Port:               cfg.ServerPort,
		DBRelClient:        DBClient,
		BookMGMTService:    BookMGMTService,
		BookMGMTController: ctrBookMgmt,
		API:                fiber.New(),
	}

	return server, nil
}
