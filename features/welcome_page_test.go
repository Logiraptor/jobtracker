package features_test

import (
	"testing"

	. "github.com/sclevine/agouti/matchers"

	. "github.com/onsi/gomega"

	"github.com/sclevine/agouti"
)

func TestWelcomePage(t *testing.T) {
	EndToEndTest("Home Page", t, func(page *agouti.Page) {
		page.Navigate("http://localhost:3000")
		Expect(page.Find("body")).To(MatchText("Welcome to Job Tracker"))
		Expect(page.FindByLink("Sign Up")).To(BeFound())
	})
}
