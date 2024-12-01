package main

import (
	page "app/page"
	"log"
	"net/http"
	"os"
)

const (
	staticFilesRoot      = "./static"
	staticFilesURLPrefix = "/static/"
	configFilesRoot      = "./build"
	configFilesURLPrefix = "/build/"
	portEnv              = "PORT"
	defaultPort          = "8080"
)

func main() {
	page.PageRoutes()
	port := getServerPort()
	hostStaticFiles()
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

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

func hostStaticFiles() {
	staticAssetFolders := []string{"html", "font", "style", "images/assets"}
	for _, folder := range staticAssetFolders {
		registerStaticFileHandler(staticFilesRoot, staticFilesURLPrefix, folder)
	}

	configAssetFolders := []string{"static/images/webp", "static/images/jpg", "static/images/minified"}
	for _, folder := range configAssetFolders {
		registerStaticFileHandler(configFilesRoot, configFilesURLPrefix, folder)
	}
}
