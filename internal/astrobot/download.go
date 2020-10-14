package astrobot

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

var imagesFilepath string = "/Users/drpaneas/github/starlordgr/static/images/post/"

func constructImageFilePath(filename string) string {
	return imagesFilepath + filename
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// DownloadImage takes a URL and a title and downloads the image into a specified path
func DownloadImage(image, title string) {
	filename := GetFilename(image, title)
	filepath := constructImageFilePath(filename)
	err := downloadFile(filepath, image)
	if err != nil {
		log.Println(err)
	}
	// Open a test image.
	src, err := imaging.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the cropped image to width = 224 preserving the aspect ratio.
	// destImage := imaging.Resize(src, 224, 0, imaging.Lanczos)
	destImage := imaging.Thumbnail(src, 224, 200, imaging.Lanczos)

	// Save the resulting image
	if strings.Contains(filepath, ".jpg") {
		err = imaging.Save(destImage, filepath, imaging.JPEGQuality(75))
	} else {
		err = imaging.Save(destImage, filepath)
	}
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
