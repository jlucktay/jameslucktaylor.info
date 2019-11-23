// +build mage

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	domain     = "jameslucktaylor.info"
	gcpProject = "jameslucktaylor-info"
)

var (
	Default = Def
	Aliases = map[string]interface{}{
		"c": Clean,
		// 	"i":     Install,
		// 	"build": Install,
		// 	"ls":    List,
	}

	site = fmt.Sprintf("https://%s", domain)
)

// Exit codes
const (
	listAppVersionsExit = iota
	unmarshalAppVersionsExit
	readDirExit
	regexCompileExit
	curlExit
	unmarshalCurlExit
)

type (
	WebApp mg.Namespace
)

// Deploys and tests the web app.
func Def() {
	mg.Deps(WebApp.Deploy)
	mg.Deps(WebApp.Test, WebApp.Prune)
}

// Deploys the web app to Google Cloud using the SDK.
// Assumes that credentials etc are already set up.
func (WebApp) Deploy() error {
	return sh.Run("gcloud", "app", "deploy", fmt.Sprintf("--project=%s", gcpProject), "--quiet", "--verbosity=critical", "--promote")
}

// Runs a quick responsiveness test against the deployed web app. Sends the output from 'hey' to stdout.
func (WebApp) Test() error {
	return sh.RunV("hey", "-z", "3s", site)
}

// Runs a load test against the deployed web app. Sends the output from 'go-wrk' to stdout.
func (WebApp) TestLoad() error {
	return sh.RunV("go-wrk", "-i", "-t=8", "-n=10000", site)
}

// Finds old versions of the web app which no longer have any traffic
// allocation, and prunes them.
func (WebApp) Prune() error {
	appVersionsOut, appVersionsErr := sh.Output("gcloud", "app", "versions", "list", "--format=json")
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

	versionsDeleteArgs := []string{"app", "versions", "delete"}

	for _, av := range appVersions {
		if av.Traffic_split == 0 {
			versionsDeleteArgs = append(versionsDeleteArgs, av.Id)
		}
	}

	versionsDeleteArgs = append(versionsDeleteArgs, "--quiet")

	if len(versionsDeleteArgs) > 4 {
		return sh.Run("gcloud", versionsDeleteArgs...)
	}

	return nil
}

// Cleans up various bits of cruft.
func Clean() {
	mg.Deps(DeleteLighthouseReports, WebApp.Prune)
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
	mg.Deps(WebApp.Deploy)
	mg.Deps(ValidateWeb, ValidateLighthouse, WebApp.Test)
	mg.Deps(Clean)
}

// Installs Lighthouse globally, via NPM.
func LighthouseInstall() {
	sh.RunV("npm", "update", "-g", "lighthouse")
}

// Runs Lighthouse against the deployed web app.
func ValidateLighthouse() error {
	mg.Deps(LighthouseInstall)
	return sh.Run("lighthouse", site, "--quiet", "--view")
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

// Runs validations only.
func Validate() {
	mg.Deps(ValidateWeb, ValidateLighthouse, WebApp.Test)
	mg.Deps(Clean)
}

// Runs up the web app locally.
func Dev() error {
	return sh.Run("go", "run", "main.go")
}

// Runs OWASP ZAP against the deployed web app.
func Zap() error {
	zapScript := "/Applications/OWASP ZAP.app/Contents/Java/zap.sh"

	_, zapErr := os.Stat(zapScript)
	if zapErr != nil {
		return zapErr
	}

	return sh.RunV(zapScript, "-cmd", "-quickurl", site)
}

// Validates the structured data that the web app exposes.
func ValidateData() {
	curlOut, curlErr := sh.Output("curl", "--silent", "--header", "Accept: application/json", fmt.Sprintf("http://linter.structured-data.org/?url=%s", site))
	if curlErr != nil {
		mg.Fatal(curlExit, curlErr)
	}

	type curlOutput struct {
		Messages []string
	}

	var curl curlOutput
	unmarshalErr := json.Unmarshal([]byte(curlOut), &curl)
	if unmarshalErr != nil {
		mg.Fatal(unmarshalCurlExit, unmarshalErr)
	}

	if len(curl.Messages) > 0 {
		fmt.Println("Messages from structured data linter:")

		for n, m := range curl.Messages {
			fmt.Printf("[%v] %v\n", n, m)
		}
	}
}

// Does everything.
func KitchenSink() {
	mg.Deps(WebApp.Deploy)
	mg.Deps(ValidateWeb, ValidateLighthouse, ValidateData, WebApp.Test, WebApp.TestLoad, Zap)
	mg.Deps(Clean)
}
