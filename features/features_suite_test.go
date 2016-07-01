package features_test

import (
	"fmt"
	"io/ioutil"
	"jobtracker/app"

	log "github.com/Sirupsen/logrus"
	. "github.com/onsi/gomega"

	"github.com/sclevine/agouti"

	"testing"
)

func init() {
	logger := log.New()
	logger.Out = ioutil.Discard
	go app.Start(app.Context{
		Port:    3000,
		AppRoot: "../",
		Logger:  logger,
	})
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

var wrapper = func(name string) func(string, func()) {
	return func(description string, body func()) {
		fmt.Println(name, description)
		body()
	}
}

var Given = wrapper("Given")
var When = wrapper("When")
var Then = wrapper("Then")
var And = wrapper("And")
