package healthy

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Healthy(w http.ResponseWriter, r *http.Request) {
	log.Info("Healthy endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
