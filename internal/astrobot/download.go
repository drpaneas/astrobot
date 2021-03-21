package astrobot

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)


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

func imageFilename(imageLink string) string {
	//Split by first delimeter
	delimiter := "/"

	tmpString := strings.Split(imageLink, delimiter)
	length := len(tmpString)
	newString := tmpString[length-1:]
	return newString[0]
}

func removeQuote(filename string) string {
	if strings.Contains(filename, "'") {
		return strings.ReplaceAll(filename, "'", "-")
	}
	return filename
}

// DownloadImage takes a URL and downloads the image into a specified path
func DownloadImage(image, filepath string) error {
	err := downloadFile(filepath, image)
	if err != nil {
		log.Println(err)
	}
	// Open a test image.
	src, err := imaging.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open image %s: %v", filepath, err)
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
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}
