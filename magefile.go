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
	// Default denotes Mage's default target when invoked without one explicitly.
	Default = Def

	// Aliases links short versions to longer target names.
	Aliases = map[string]interface{}{
		"b": WebApp.Build,
		"c": WebApp.Clean,
		"p": WebApp.Prune,
	}

	site = fmt.Sprintf("https://%s", domain)
)

// Exit codes
const (
	listAppVersionsExit = iota
	unmarshalAppVersionsExit
)

type (
	// WebApp collects all of the targets involving the web app.
	WebApp mg.Namespace
)

// Def is assigned as the 'Default' target, so it builds the web app.
func Def() {
	mg.Deps(WebApp.Build)
}

// Build the web app using Hugo.
func (WebApp) Build() error {
	return sh.RunV(
		"hugo",
		"--baseURL", "//hugo-dot-jameslucktaylor-info.appspot.com",
		"--cleanDestinationDir",
		"--enableGitInfo",
		"--gc",
		"--minify",
		"--source", "hugo",
		"--stepAnalysis",
		"--templateMetrics",
		"--templateMetricsHints",
		"--verbose",
	)
}

// Prune will find old versions of the web app which no longer have any traffic
// allocation, and delete them.
func (WebApp) Prune() error {
	appVersionsOut, appVersionsErr := sh.Output("gcloud", "app", "versions", "list", "--format=json", "--service=hugo")
	if appVersionsErr != nil {
		mg.Fatal(listAppVersionsExit, appVersionsErr)
	}

	type appVersion struct {
		// ID is the site version returned by the GCloud SDK.
		ID string `json:"id"`
		// TrafficSplit is the percentage of traffic directed at this version.
		TrafficSplit float32 `json:"traffic_split"`
	}

	var appVersions []appVersion
	unmarshalErr := json.Unmarshal([]byte(appVersionsOut), &appVersions)
	if unmarshalErr != nil {
		mg.Fatal(unmarshalAppVersionsExit, unmarshalErr)
	}

	versionsDeleteArgs := []string{"app", "versions", "delete"}

	for _, av := range appVersions {
		if av.TrafficSplit == 0 {
			versionsDeleteArgs = append(versionsDeleteArgs, av.ID)
		}
	}

	versionsDeleteArgs = append(versionsDeleteArgs, "--quiet")

	if len(versionsDeleteArgs) > 4 {
		return sh.Run("gcloud", versionsDeleteArgs...)
	}

	return nil
}

// Clean will delete various bits of cruft.
func (WebApp) Clean() error {
	return sh.RunV("rm", "-rfv", "hugo/public")
}
