package main

import (
	. "app/Page"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"strings"
	"io"
	"bytes"
	"github.com/chai2010/webp"
)

const (
	staticDir   = "./static"
	staticPath  = "/static/"
	portEnv     = "PORT"
	defaultPort = "8080"
)

func serveStaticFiles(folder string) {
	path := staticPath + folder + "/"
	fileServer := http.FileServer(http.Dir(staticDir + "/" + folder))
	http.Handle(path, http.StripPrefix(path, fileServer))
}

func getPort() string {
	port := os.Getenv(portEnv)
	if port == "" {
		port = defaultPort
	}
	return ":" + port
}

func main() {
	dirPath := "./static/images/assets"
	// Output directories
	webpOutputDir := "./static/images/assets/context/webp"
	jpegOutputDir := "./static/images/assets/context/jpg"

	createDir(webpOutputDir)
	createDir(jpegOutputDir)

	err := processImages(dirPath, webpOutputDir, jpegOutputDir)
	if err != nil {
		log.Fatalf("Error processing images: %v", err)
	}

	// Setup routes and static file serving
	Routes()
	serveStaticFiles("html")
	serveStaticFiles("font")
	serveStaticFiles("style")
	serveStaticFiles("images")

	port := getPort()
	fmt.Printf("Listening on %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}


// Check if file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// Compare two images
func compareImages(path1, path2 string) bool {
	file1, err1 := os.Open(path1)
	if err1 != nil {
		return false
	}
	defer file1.Close()

	file2, err2 := os.Open(path2)
	if err2 != nil {
		return false
	}
	defer file2.Close()

	// Read entire files
	data1, err1 := io.ReadAll(file1)
	data2, err2 := io.ReadAll(file2)

	if err1 != nil || err2 != nil {
		return false
	}

	return bytes.Equal(data1, data2)
}

func processImages(dirPath, webpOutputDir, jpegOutputDir string) error {
	sem := make(chan struct{}, runtime.NumCPU())
	var wg sync.WaitGroup
	var processErr error
	var mu sync.Mutex

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			// Prevent walking into context subdirectories
			if strings.Contains(path, "/context/") {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if path is directly under assets
		relPath, err := filepath.Rel(dirPath, path)
		if err != nil || strings.Contains(relPath, string(filepath.Separator)) {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
			return nil
		}

		wg.Add(1)
		sem <- struct{}{}

		go func(path string) {
			defer wg.Done()
			defer func() { <-sem }()

			file, err := os.Open(path)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", path, err)
				return
			}
			defer file.Close()

			var img image.Image
			switch ext {
			case ".jpg", ".jpeg":
				img, err = jpeg.Decode(file)
			case ".png":
				img, err = png.Decode(file)
			case ".webp":
				img, err = webp.Decode(file)
			}

			if err != nil {
				log.Printf("Error decoding file %s: %v\n", path, err)
				return
			}

			baseName := filepath.Base(path)
			outputWebP := filepath.Join(webpOutputDir, baseName+".webp")
			outputJPEG := filepath.Join(jpegOutputDir, baseName+".jpg")

			// Check if WebP file already exists and matches
			if !fileExists(outputWebP) || !compareImages(path, outputWebP) {
				if err = saveAsWebP(outputWebP, img); err != nil {
					log.Printf("Error saving WebP file %s: %v\n", outputWebP, err)
					mu.Lock()
					processErr = err
					mu.Unlock()
				}
			}

			// Check if JPEG file already exists and matches
			if !fileExists(outputJPEG) || !compareImages(path, outputJPEG) {
				if err = saveAsJPEG(outputJPEG, img); err != nil {
					log.Printf("Error saving JPEG file %s: %v\n", outputJPEG, err)
					mu.Lock()
					processErr = err
					mu.Unlock()
				}
			}
		}(path)

		return nil
	})

	wg.Wait()

	if err != nil {
		return err
	}

	return processErr
}
// Creates a directory if it doesn't exist
func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Error creating directory %s: %v\n", dir, err)
		}
	}
}

// Saves an image as JPEG
func saveAsJPEG(path string, img image.Image) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating JPEG file: %v", err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return fmt.Errorf("error encoding JPEG: %v", err)
	}

	return nil
}

// Saves an image as WebP
func saveAsWebP(path string, img image.Image) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating WebP file: %v", err)
	}
	defer outputFile.Close()

	// Encode as WebP (lossless)
	err = webp.Encode(outputFile, img, &webp.Options{Lossless: true})
	if err != nil {
		return fmt.Errorf("error encoding WebP: %v", err)
	}

	return nil
}
