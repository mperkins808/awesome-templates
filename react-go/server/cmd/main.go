package main

import (
	// "embed"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/mperkins808/awesome-templates/server/pkg/api/healthy"
	"github.com/mperkins808/awesome-templates/server/pkg/util"
	vueglue "github.com/mperkins808/vite-go"
	log "github.com/sirupsen/logrus"
)

// this is not for vite, but to help our
// makefile stop the process:
var pidFile string
var pidDeleteChan chan os.Signal

var vueData *vueglue.VueGlue

func main() {

	config := util.GenConfig(os.Getenv("ENVIRONMENT"))

	pid := strconv.Itoa(os.Getpid())
	_ = os.WriteFile(pidFile, []byte(pid), 0644)
	util.WaitForSignal(pidFile, pidDeleteChan)
	defer util.CleanupPID(pidFile)

	glue, err := vueglue.NewVueGlue(&config)

	if err != nil {
		log.Fatalln(err)
		return
	}

	// Set up our router
	mux := chi.NewRouter()

	// Set up a file server for our assets.
	vueData = glue
	fsHandler, err := glue.FileServer()
	if err != nil {
		log.Errorf("could not set up static file server %v", err)
		return
	}

	// Serve Files
	mux.Handle(config.URLPrefix+"*", fsHandler)
	mux.Handle("/*", logRequest(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.PageWithAVue(w, r, vueData)
	})))

	// API Routes
	mux.Get("/api/healthy", healthy.Healthy)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	log.Infof("Starting server on :%v", PORT)
	err = http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		diff := time.Since(start)
		log.Infof("%s - %s %s %s completed in %v", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI(), diff)
	})
}
