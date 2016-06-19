package features_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("Login", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		Expect(page.Navigate("http://localhost:3000")).To(Succeed())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("Allows me to login", func() {
		By("showing a login form", func() {
			Expect(page.FindByLabel("Email")).To(BeFound())
			Expect(page.FindByLabel("Password")).To(BeFound())
			Expect(page.FindByButton("Sign In")).To(BeFound())
		})

		By("logging me in", func() {
			page.FindByLabel("Email").SendKeys("test@example.com")
			page.FindByLabel("Password").SendKeys("password")
			page.FindByButton("Sign In").Click()

			Expect(page.Find("body")).To(MatchText("test@example\\.com"))
		})
	})
})
