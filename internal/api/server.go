package api

import (
	"core-control/internal/docker"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	dockerService *docker.DockerService
}

func NewServer(ds *docker.DockerService) *Server {
	return &Server{dockerService: ds}
}

func (s *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/containers", s.handleListContainers)
	mux.HandleFunc("POST /api/containers/{id}/start", s.handleStartContainer)
	mux.HandleFunc("POST /api/containers/{id}/stop", s.handleStopContainer)
	mux.HandleFunc("POST /api/containers/{id}/restart", s.handleRestartContainer)
	mux.HandleFunc("GET /api/containers/{id}/logs", s.handleGetLogs)

	return mux
}

func (s *Server) handleListContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := s.dockerService.GetContainers(r.Context())
	if err != nil {
		log.Printf("Error getting containers: %v", err)
		http.Error(w, "Failed to get containers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}

func (s *Server) handleStartContainer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := s.dockerService.StartContainer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleStopContainer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := s.dockerService.StopContainer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleRestartContainer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := s.dockerService.RestartContainer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetLogs(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	logs, err := s.dockerService.GetContainerLogs(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(logs))
}
