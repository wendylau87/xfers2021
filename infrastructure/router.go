package infrastructure

import (
	"github.com/go-chi/chi"
	"github.com/wendylau87/xfers2021/controllers"
	"github.com/wendylau87/xfers2021/domain"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/infrastructure/sqlhandler"
	"github.com/wendylau87/xfers2021/usecases"

	"net/http"
	"os"


)

// Dispatch is handle routing
func Dispatch(logger logger.Logger, sqlHandler sqlhandler.SQLHandler) {
	dom := domain.Init(logger, sqlHandler)
	uc := usecases.Init(logger, dom)
	kursController := controllers.InitKursController(uc,logger)

	r := chi.NewRouter()
	r.Get("/api/indexing", kursController.Indexing)
	r.Post("/api/kurs", kursController.CreateKurs)
	r.Delete("/api/kurs/{date}", kursController.DeleteKurs)
	r.Put("/api/kurs", kursController.UpdateKurs)
	r.Get("/api/kurs/{symbol}", kursController.GetBySymbol)
	r.Get("/api/kurs", kursController.GetByDate)
	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}

	logger.LogAccess("HTTP served on %s", os.Getenv("SERVER_PORT"))
}
