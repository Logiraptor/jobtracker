package features_test

import (
	"jobtracker/app"
	"jobtracker/app/web"

	. "github.com/onsi/gomega"

	"github.com/sclevine/agouti"

	"testing"
)

var testContext = app.Context{
	Port:    3000,
	AppRoot: "../",
	Logger:  web.NilLogger{},
}

func init() {
	go app.Start(testContext)
}

var agoutiDriver *agouti.WebDriver

func EndToEndTest(name string, t *testing.T, body func(page *agouti.Page)) {
	RegisterTestingT(t)

	agoutiDriver = agouti.ChromeDriver()
	Expect(agoutiDriver.Start()).To(Succeed())
	page, err := agoutiDriver.NewPage()
	Expect(err).NotTo(HaveOccurred())

	body(page)

	Expect(page.Destroy()).To(Succeed())
	Expect(agoutiDriver.Stop()).To(Succeed())
}

var Given = func(description string, body func()) {
	body()
}
var When = Given
var Then = Given
var And = Given
