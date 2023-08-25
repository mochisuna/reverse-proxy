package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type ResponseBody struct {
	Header string `json:"header"`
}

type Handler struct {
	render    *render.Render
	logger    *logrus.Logger
	validator *validator.Validate
}

func NewHandler() *Handler {
	return &Handler{
		render:    render.New(render.Options{}),
		logger:    logrus.New(),
		validator: validator.New(),
	}
}

func (h *Handler) configLog(fileName string) {
	logfile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	h.logger.SetOutput(multiLogFile)
}

func (h *Handler) processRequests(w http.ResponseWriter, r *http.Request) {
	response := &ResponseBody{
		Header: fmt.Sprint(r.Header),
	}
	h.logger.Infof("response: %v", response)
	w.Header().Set("Content-Type", "application/json")
	h.render.JSON(w, http.StatusOK, response)
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	h := NewHandler()
	h.configLog("./log/test.log")
	h.logger.Info("Start API server")

	router := chi.NewRouter()
	router.Get("/health", h.healthCheck)
	router.Get("/", h.processRequests)
	if err := http.ListenAndServe(":80", router); err != nil {
		h.logger.Fatal("Error starting server:", err)
	}
}
