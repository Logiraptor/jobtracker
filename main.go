package main

import (
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.ServeFile(rw, req, "public/index.html")
	})

	http.HandleFunc("/pdf", func(rw http.ResponseWriter, req *http.Request) {
		var buf = make([]byte, 32)
		rand.Read(buf)
		randomString := hex.EncodeToString(buf)
		f, err := os.Create(filepath.Join(os.TempDir(), randomString+".html"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		f.WriteString("<html></html>")
		f.Close()

		pdfName := f.Name() + ".pdf"

		log.Println("wkhtmltopdf", f.Name(), pdfName)
		output, err := exec.Command("wkhtmltopdf", f.Name(), pdfName).CombinedOutput()
		log.Println(string(output))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		http.ServeFile(rw, req, pdfName)
		os.Remove(f.Name())
		os.Remove(pdfName)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("App started on port: %s", port)
	log.Println(http.ListenAndServe(":"+port, nil))
}
