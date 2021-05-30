package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

 
func base_handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	request.ParseForm()
	fmt.Println("request.Form: ", request.Form)
}

func capture_handler(response http.ResponseWriter, request *http.Request) {
	base_handler(response, request)

	dx, err := strconv.Atoi(request.FormValue("x"))
	dy, err := strconv.Atoi(request.FormValue("y"))
	width, err := strconv.Atoi(request.FormValue("w"))
	height, err := strconv.Atoi(request.FormValue("h"))

	fmt.Println(dx, dy, width, height)

	if width < 1 || height < 1 {
		bounds := screenshot.GetDisplayBounds(0)
		width, height = bounds.Dx(), bounds.Dy()
	}

	img, err := screenshot.Capture(dx, dy, width, height)
	if err != nil {
		panic(err)
	}

	file_name := fmt.Sprintf("capture-%d-%d-%d-%d-%d.png", time.Now().Unix(), dx, dy, width, height)
	temp_img_path := filepath.Join(os.TempDir(), file_name)
	fmt.Println(temp_img_path)
	file, _ := os.Create(temp_img_path)
	defer file.Close()
	png.Encode(file, img)
	fmt.Println(file.Name())

	response.Header().Set("content-type", "image/png")
	data, _ := ioutil.ReadFile(file.Name())
	response.Write(data)
}

func localimage_handler(response http.ResponseWriter, request *http.Request) {
	base_handler(response, request)

	file_path := request.FormValue("filePath")
	file_path, err := url.QueryUnescape(file_path)
	file, err := ioutil.ReadFile(file_path)
	
	if err!=nil {
		fmt.Println(err)
		response.WriteHeader(http.StatusNotFound)
	}else{
		response.Header().Set("content-type", "image/png")
		response.Write(file)
	}
}

func start_http_server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", base_handler)

	mux.HandleFunc("/capture", capture_handler)
	mux.HandleFunc("/localimage", localimage_handler)

	server := &http.Server{
		Addr:    ":10707",
		Handler: mux,
	}
	fmt.Println("Starting httpserver ...")
	server.ListenAndServe()
}

func main() {
	start_http_server()
}


 