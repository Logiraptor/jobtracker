package features_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("SignUp", func() {
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

	It("Allows me to sign up", func() {
		By("showing a sign up link", func() {
			Expect(page.FindByLink("Sign Up")).To(BeFound())
			page.FindByLink("Sign Up").Click()
		})

		By("showing a sign up form", func() {
			Expect(page.FindByLabel("Email")).To(BeFound())
			Expect(page.FindByLabel("Password")).To(BeFound())
			Expect(page.FindByLabel("Current Password")).To(BeFound())
			Expect(page.FindByButton("Sign Up")).To(BeFound())
		})

		By("logging me in", func() {
			page.FindByLabel("Email").SendKeys("test@example.com")
			page.FindByLabel("Password").SendKeys("password")
			page.FindByLabel("Current Password").SendKeys("password")
			page.FindByButton("Sign Up").Click()

			Expect(page.Find("body")).To(MatchText("test@example\\.com"))
		})
	})
})
