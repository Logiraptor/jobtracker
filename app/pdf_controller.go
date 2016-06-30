package app

import (
	"io"
	"jobtracker/app/web"
	"net/http"
	"os/exec"
)

type PdfController struct {
	web.Logger
}

func NewPdfController(logger web.Logger) PdfController {
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
