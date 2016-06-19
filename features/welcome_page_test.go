package features_test

import (
	"jobtracker/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("WelcomePage", func() {
	var page *agouti.Page

	BeforeEach(func() {
		app.Start(testContext)

		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("renders the home page", func() {
		By("showing a welcome message", func() {
			Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			Expect(page.Find("body")).To(MatchText("Welcome to Job Tracker"))
		})

		By("showing a sign up link", func() {
			Expect(page.FindByLink("Sign Up")).To(BeFound())
		})
	})
})
