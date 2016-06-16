package features_test

import (
	"jobtracker/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
)

var testContext = app.Context{
	Port:    3000,
	AppRoot: "../",
	Logger:  app.NilLogger{},
}

func TestFeatures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Features Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// Choose a WebDriver:

	agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	// agoutiDriver = agouti.ChromeDriver()

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
