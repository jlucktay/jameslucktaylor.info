// +build mage

package main

import (
	"encoding/json"
	"fmt"

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
)

type (
	WebApp mg.Namespace
)

// Deploys and tests the web app.
func Def() {
	mg.Deps(WebApp.Deploy)
	mg.Deps(Clean)
}

// Deploys the web app to Google Cloud using the SDK.
// Assumes that credentials etc are already set up.
func (WebApp) Deploy() error {
	// return sh.Run("gcloud", "app", "deploy", fmt.Sprintf("--project=%s", gcpProject), "--quiet", "--verbosity=critical", "--promote")
	return sh.Run("hugo", "--version")
}

// Finds old versions of the web app which no longer have any traffic
// allocation, and prunes them.
func (WebApp) Prune() error {
	appVersionsOut, appVersionsErr := sh.Output("gcloud", "app", "versions", "list", "--format=json", "--service=hugo")
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
func Clean() error {
	return sh.Run("rm", "-rf", "public")
}
