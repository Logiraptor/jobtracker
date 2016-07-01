package app

import (
	"io"
	"net/http"
	"os/exec"

	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
)

type PdfController struct {
	*log.Logger
}

func NewPdfController(logger *log.Logger) PdfController {
	return PdfController{
		Logger: logger,
	}
}

func (p PdfController) Generate(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/pdf")

	cmd := exec.Command("wkhtmltopdf", "-", "-")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	io.WriteString(stdin, "<html></html>")
	stdin.Close()

	io.Copy(rw, stdout)
	cmd.Wait()
}

func (p PdfController) Register(mux *mux.Router) {
	mux.Path("/generate_pdf").HandlerFunc(p.Generate)
}
