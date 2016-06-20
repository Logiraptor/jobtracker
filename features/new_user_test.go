package features_test

import (
	"regexp"
	"testing"

	"github.com/manveru/faker"

	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

func TestSignUpOutAndIn(t *testing.T) {
	fake, _ := faker.New("en")
	email := fake.Email()
	password := fake.Characters(20)

	EndToEndTest("User Registration, Logout, Login", t, func(page *agouti.Page) {
		page.Navigate("http://localhost:3000")

		When("I sign up", func() {
			page.FindByLink("Sign Up").Click()

			page.FindByLabel("Email").SendKeys(email)
			page.FindByLabel("Password").SendKeys(password)
			page.FindByLabel("Current Password").SendKeys(password)
			page.FindByButton("Sign Up").Click()
		})

		Then("I see my email", func() {
			Expect(page.Find("body")).To(MatchText(regexp.QuoteMeta(email)))
		})

		When("I log out", func() {
			page.FindByLink("Log Out").Click()
		})

		Then("I don't see my email", func() {
			Expect(page.Find("body")).NotTo(MatchText(regexp.QuoteMeta(email)))
		})

		When("I sign in", func() {
			page.FindByLabel("Email").SendKeys(email)
			page.FindByLabel("Password").SendKeys(password)
			page.FindByButton("Log In").Click()
		})

		Then("I see my email", func() {
			Expect(page.Find("body")).To(MatchText(regexp.QuoteMeta(email)))
		})
	})
}
