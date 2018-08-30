// +build mage

package main

import (
	"encoding/json"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	site = "https://jameslucktaylor.info"
)

var Default = DefaultTarget

const (
	listAppVersionsExit = iota
	unmarshalAppVersionsExit
	deployExit
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
