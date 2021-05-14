package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/wendylau87/xfers2021/entities"
	"github.com/wendylau87/xfers2021/infrastructure/logger"
	"github.com/wendylau87/xfers2021/usecases"
	"net/http"
)

type KursController struct {
	Usecase *usecases.Usecase
	Logger  logger.Logger
}

func InitKursController(uc *usecases.Usecase, logger logger.Logger) *KursController {
	return &KursController{
		Usecase: uc,
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (c *KursController) GetByDate(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	items, err := c.Usecase.Kurs.GetKursByDate(r.URL.Query().Get("startDate"), r.URL.Query().Get("endDate"))
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (c *KursController) GetBySymbol(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	symbol := chi.URLParam(r, "symbol")

	items, err, code := c.Usecase.Kurs.GetKursByName(symbol, r.URL.Query().Get("startDate"), r.URL.Query().Get("endDate"))
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (c *KursController) CreateKurs(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var createKurs entities.CreateKurs
	err := json.NewDecoder(r.Body).Decode(&createKurs)
	if err != nil {
		c.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}


	kurs, err := c.Usecase.Kurs.CreateKurs(createKurs)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResult, err := json.Marshal(&kurs)
	w.Write(jsonResult)
}

func (c *KursController) UpdateKurs(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var putKurs entities.PutKurs
	err := json.NewDecoder(r.Body).Decode(&putKurs)
	if err != nil {
		c.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}


	kurs, err, code := c.Usecase.Kurs.UpdateKurs(putKurs)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResult, err := json.Marshal(&kurs)
	w.Write(jsonResult)
}

func (c *KursController) DeleteKurs(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	date := chi.URLParam(r, "date")
	err := c.Usecase.Kurs.DeleteKurs(date)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}

func (c *KursController) Indexing(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	err := c.Usecase.Kurs.GenerateKurs()
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
	}
	w.Header().Set("Content-Type", "application/json")

}