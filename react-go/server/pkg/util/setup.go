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

// flag.StringVar(&config.Environment, "env", "development", "development|production")
// flag.StringVar(&config.JSProjectPath, "assets", "../app", "location of javascript files.")
// flag.BoolVar(&config.JSInExternalDir, "external", true, "location of javascript files.")
// flag.StringVar(&config.DevServerDomain, "domain", "localhost", "Domain of the dev server.")
// flag.BoolVar(&config.HTTPS, "https", false, "Expect dev server to use HTTPS")
// flag.StringVar(&config.AssetsPath, "dist", "", "dist directory relative to the JS project directory.")
// flag.StringVar(&config.EntryPoint, "entryp", "", "relative path of the entry point of the js app.")
// flag.StringVar(&config.Platform, "platform", "", "vue|react|svelte")
// flag.StringVar(&pidFile, "pid", "", "location of optional pid file.")

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
