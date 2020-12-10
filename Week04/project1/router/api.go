package router

import (
	"apiserver/handler/sd"
	"net/http"
)

func RegisterApi() {
	h := http.NewServeMux()
	h.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir("admin"))))

	h.HandleFunc("/health", sd.HealthCheck)
	h.HandleFunc("/disk", sd.DiskCheck)
	h.HandleFunc("/cpu", sd.CPUCheck)
	h.HandleFunc("/ram", sd.RAMCheck)
}
