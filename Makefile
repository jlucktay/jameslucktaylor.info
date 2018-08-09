.POSIX:

SHELL := /usr/local/bin/bash

default: deploy test prune-old-versions

deploy:
	gcloud app deploy --quiet

test:
	hey -z 3s http://jameslucktaylor.info

prune-old-versions:
	gcloud app versions delete $(shell gcloud app versions list --format="json" | jq -r '[ .[] | select(.traffic_split == 0) | .id ] | join(" ")') --quiet

clean:
	rm -fv jameslucktaylor.info_*.report.html

full: deploy validate-web validate-lighthouse test clean

validate: validate-web validate-lighthouse test clean

kitchen-sink: deploy validate-web validate-lighthouse test zap clean

dev:
	dev_appserver.py app.yaml --support_datastore_emulator=true

validate-data:
	curl --silent --header 'Accept: application/json' http://linter.structured-data.org/?url=https://jameslucktaylor.info | jq '.messages'

validate-lighthouse: lighthouse-install
	lighthouse https://jameslucktaylor.info --view

validate-web:
	open "https://validator.w3.org/unicorn/check?ucn_uri=jameslucktaylor.info"
	open "https://ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info&clearCache=on"
	open "https://developers.google.com/speed/pagespeed/insights/?url=jameslucktaylor.info"
	open "https://search.google.com/test/mobile-friendly?url=jameslucktaylor.info"
	open "https://developers.facebook.com/tools/debug/og/object/?q=jameslucktaylor.info"
	open "https://developers.facebook.com/tools/debug/sharing/?q=jameslucktaylor.info"
	open "https://realfavicongenerator.net/favicon_checker?protocol=https&site=jameslucktaylor.info"

zap:
	/Applications/OWASP\ ZAP.app/Contents/Java/zap.sh -cmd -quickurl http://jameslucktaylor.info

lighthouse-install:
	npm update -g lighthouse
