package util

import (
	"os"
	"os/signal"
	"syscall"

	vueglue "github.com/mperkins808/vite-go"

	log "github.com/sirupsen/logrus"
)

func WaitForSignal(pidFile string, pidDeleteChan chan os.Signal) {
	if pidFile == "" {
		return
	}
	pidDeleteChan = make(chan os.Signal, 1)
	signal.Notify(pidDeleteChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-pidDeleteChan
		log.Info("Deleted pid file")
		_ = os.Remove(pidFile)
		os.Exit(0)

	}()
}

func GenConfig(env string) vueglue.ViteConfig {
	var config vueglue.ViteConfig
	if env == "production" {
		config.Environment = "production"
		config.JSProjectPath = "../app/assets"
		config.JSInExternalDir = true
		config.DevServerDomain = "localhost"
		config.HTTPS = false
		config.AssetsPath = "dist"
		config.EntryPoint = ""
		config.Platform = ""
	} else {
		config.Environment = "development"
		config.JSProjectPath = "../app"
		config.JSInExternalDir = true
		config.DevServerDomain = "localhost"
		config.HTTPS = false
		config.AssetsPath = "dist"
		config.EntryPoint = ""
		config.Platform = ""
	}

	// good to know
	if env != "production" && env != "development" {
		log.Warn("Environment variable ENVIRONMENT must be either development or production. Defaulting to development")
	}

	// derived
	if config.Environment == "production" {
		// Running inside docker
		config.FS = os.DirFS("/app")
		config.URLPrefix = "/assets/"
	} else if config.Environment == "development" {
		config.FS = os.DirFS("../app")
	} else {
		log.Fatalln("unsupported environment setting")
	}

	return config
}
