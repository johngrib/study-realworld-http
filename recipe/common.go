package recipe

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RequestGet(uriAddress string) {
	resp, err := http.Get(uriAddress)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

func RequestGetWithQuery(uriAddress string) {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(uriAddress + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

// RequestHead requests with the HEAD method.
func RequestHead(uriAddress string) {
	resp, err := http.Head(uriAddress)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}

func RequestPostWithUrlEncoded(uriAddress string) {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm(uriAddress, values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func RequestPostWithFileBody(uriAddress, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(uriAddress, "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

func RequestPostWithTextString(uriAddress, text string) {
	reader := strings.NewReader(text)
	resp, err := http.Post(uriAddress, "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}

// RequestPostWithMultipart requests with multipart/form-data
func RequestPostWithMultipart(uriAddress string) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// 파일 이외의 필드는 WriteField 메소드로 등록
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	resp, err := http.Post(uriAddress, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
