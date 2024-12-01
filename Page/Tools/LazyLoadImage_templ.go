// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package tools

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func OptimizeImage(img string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if isImageFormatValid(img) {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 1)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("/config/context/images/webp/" + fileName(img) + ".webp")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `Page/Tools/LazyLoadImage.templ`, Line: 23, Col: 103}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 2)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(img)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `Page/Tools/LazyLoadImage.templ`, Line: 24, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 3)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(createImageName(img))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `Page/Tools/LazyLoadImage.templ`, Line: 24, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 4)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templ.WriteWatchModeString(templ_7745c5c3_Buffer, 5)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func createImageName(imageSource string) string {
	s := imageSource
	lastSlash := strings.LastIndex(s, "/")
	if lastSlash != -1 {
		s = s[lastSlash+1:]
	}
	lastDot := strings.LastIndex(s, ".")
	if lastDot != -1 {
		s = s[:lastDot]
	}
	return s
}

func fileName(imageSource string) string {
	s := imageSource
	lastSlash := strings.LastIndex(s, "/")
	if lastSlash != -1 {
		s = s[lastSlash+1:]
	}
	return s
}

func isImageFormatValid(src string) bool {
	if len(src) != 0 {
		validImageFormats := []string{".png", ".jpg", ".jpeg", ".webp", ".gif", ".bmp", ".tiff"}
		for _, ext := range validImageFormats {
			if strings.HasSuffix(strings.ToLower(src), ext) {
				return true
			}
		}
	}
	return false
}

func ConvertAllImages() {

}

func OptimizeImages() {
	dirPath := "./static/images/assets"
	webpOutputDir := "./static/images/assets/context/webp"
	jpegOutputDir := "./static/images/assets/context/jpg"

	createDir(webpOutputDir)
	createDir(jpegOutputDir)

	err := processImages(dirPath, webpOutputDir, jpegOutputDir)
	if err != nil {
		log.Fatalf("Error processing images: %v", err)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

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
		if d.IsDir() {
			if strings.Contains(path, "/context/") {
				return filepath.SkipDir
			}
			return nil
		}

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

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Error creating directory %s: %v\n", dir, err)
		}
	}
}

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

func saveAsWebP(path string, img image.Image) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating WebP file: %v", err)
	}
	defer outputFile.Close()

	err = webp.Encode(outputFile, img, &webp.Options{Lossless: true})
	if err != nil {
		return fmt.Errorf("error encoding WebP: %v", err)
	}

	return nil
}

var _ = templruntime.GeneratedTemplate
