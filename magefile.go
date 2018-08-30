// +build mage

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = DefaultTarget

const (
	domain = "jameslucktaylor.info"
)

var (
	site = fmt.Sprintf("https://%s", domain)
)

const (
	listAppVersionsExit = iota
	unmarshalAppVersionsExit
	deployExit
	readDirExit
	regexCompileExit
)

// Deploys and tests the web app.
func DefaultTarget() {
	mg.Deps(Deploy)
	mg.Deps(TestSite, PruneOldVersions)
}

// Deploys the web app to Google Cloud using the SDK.
// Assumes that credentials etc are already set up.
func Deploy() {
	deployErr := sh.Run("gcloud", "app", "deploy", "--quiet")
	if deployErr != nil {
		mg.Fatal(deployExit, deployErr)
	}
}

// Runs a quick responsiveness test against the site. Sends the output from
// 'hey' to stdout.
func TestSite() {
	sh.RunV("hey", "-z", "3s", site)
}

// Runs a load test against the site. Sends the output from 'go-wrk' to stdout.
func TestLoad() {
	sh.RunV("go-wrk", "-i", "-c=200", "-t=8", "-n=10000", site)
}

// Finds old versions of the web app which no longer have any traffic
// allocation, and prunes them.
func PruneOldVersions() {
	appVersionsOut, appVersionsErr := sh.Output("gcloud", "app", "versions",
		"list", "--format=json")
	if appVersionsErr != nil {
		mg.Fatal(listAppVersionsExit, appVersionsErr)
	}

	type appVersion struct {
		Id            string
		Traffic_split float32
	}

	var appVersions []appVersion
	unmarshalErr := json.Unmarshal([]byte(appVersionsOut), &appVersions)
	if unmarshalErr != nil {
		mg.Fatal(unmarshalAppVersionsExit, unmarshalErr)
	}

	for _, av := range appVersions {
		if av.Traffic_split == 0 {
			sh.Run("gcloud", "app", "versions", "delete", av.Id, "--quiet")
		}
	}
}

// Cleans up various bits of cruft.
func Clean() {
	mg.Deps(DeleteLighthouseReports, PruneOldVersions)
}

// Deletes the HTML reports generated when Lighthouse runs.
func DeleteLighthouseReports() {
	files, readDirErr := ioutil.ReadDir(".")
	if readDirErr != nil {
		mg.Fatal(readDirExit, readDirErr)
	}

	r, regexErr := regexp.Compile(`^jameslucktaylor\.info_.*\.report\.html$`)
	if regexErr != nil {
		mg.Fatal(regexCompileExit, regexErr)
	}

	for _, file := range files {
		if r.MatchString(file.Name()) {
			fmt.Println(file.Name())
			sh.Rm(file.Name())
		}
	}
}

// Deploys, validates, tests, cleans up.
func Full() {
	mg.Deps(Deploy)
	mg.Deps(ValidateWeb, ValidateLighthouse, TestSite)
	mg.Deps(Clean)
}

// Installs Lighthouse globally, via NPM.
func LighthouseInstall() {
	sh.RunV("npm", "install", "-g", "lighthouse")
	sh.RunV("npm", "update", "-g", "lighthouse")
}

// Runs Lighthouse against the deployed web app.
func ValidateLighthouse() {
	mg.Deps(LighthouseInstall)
	sh.Run("lighthouse", site, "--view")
}

// Runs various validators from across the web on the deployed web app.
func ValidateWeb() {
	validators := []string{
		"https://validator.w3.org/unicorn/check?ucn_uri=",
		"https://ssllabs.com/ssltest/analyze.html?clearCache=on&d=",
		"https://developers.google.com/speed/pagespeed/insights/?url=",
		"https://search.google.com/test/mobile-friendly?url=",
		"https://developers.facebook.com/tools/debug/og/object/?q=",
		"https://developers.facebook.com/tools/debug/sharing/?q=",
		"https://realfavicongenerator.net/favicon_checker?protocol=https&site=",
	}

	for _, v := range validators {
		sh.Run("open", fmt.Sprintf("%s%s", v, domain))
	}
}
