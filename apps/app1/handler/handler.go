package handler

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("App1, OK"))
}

func GetAuthor(w http.ResponseWriter, req *http.Request) {
	// ctx, err := context.WithTimeout(req.Context(), time.Minute*2)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Error creating context"))
	// 	return
	// }

	// db.New(db.DB).GetAuthor(ctx, 1)
}
