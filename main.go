package main

import (
	page "app/Page"
	tools "app/Page/Tools"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	staticFilesRoot      = "./static"
	staticFilesURLPrefix = "/static/"
	configFilesRoot      = "./config"
	configFilesURLPrefix = "/config/"
	portEnv              = "PORT"
	defaultPort          = "8080"
)

func registerStaticFileHandler(rootDirectory, urlPrefix, folder string) {
	path := urlPrefix + folder + "/"
	fileServer := http.FileServer(http.Dir(rootDirectory + "/" + folder))
	http.Handle(path, http.StripPrefix(path, fileServer))
}

func getServerPort() string {
	if port := os.Getenv(portEnv); port != "" {
		return ":" + port
	}
	return ":" + defaultPort
}

func main() {
	page.PageRoutes()

	staticAssetFolders := []string{"html", "font", "style", "images/assets"}
	for _, folder := range staticAssetFolders {
		registerStaticFileHandler(staticFilesRoot, staticFilesURLPrefix, folder)
	}

	configAssetFolders := []string{"context/images/webp", "context/images/jpg", "context/images/minified"}
	for _, folder := range configAssetFolders {
		registerStaticFileHandler(configFilesRoot, configFilesURLPrefix, folder)
	}

	port := getServerPort()
	fmt.Printf("Listening on %s\n", port)

	go tools.OptimizeImages()
	go tools.LoadSettings()

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
