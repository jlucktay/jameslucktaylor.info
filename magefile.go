// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

var Default = Deploy

const (
	site = "https://jameslucktaylor.info"
)

// This is a demo.
func Demo() {
	fmt.Println("Hello world!")
}

// Deploy the web app to Google Cloud using the SDK.
// Assumes that credentials etc are already set up.
func Deploy() {
	sh.Run("gcloud", "app", "deploy", "--quiet")
}

// Run a quick responsiveness test against the site.
func TestSite() {
	sh.RunV("hey", "-z", "3s", site)
}
