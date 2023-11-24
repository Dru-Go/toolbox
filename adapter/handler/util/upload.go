package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CSVUploadHandler(request *http.Request, form_ref string) (string, error) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	request.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `import`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := request.FormFile(form_ref)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := os.CreateTemp("temp-imports", "upload-*.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	nBytes, err := io.Copy(tempFile, file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// return that we have successfully uploaded our file!
	fmt.Printf("Successfully Written %v amount of data, file name is %s \n", nBytes, tempFile.Name())
	return tempFile.Name(), nil
}
