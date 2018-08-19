.POSIX:

### From: https://stackoverflow.com/questions/714100/os-detecting-makefile
# ifeq ($(OS),Windows_NT)
#     CCFLAGS += -D WIN32
#     ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
#         CCFLAGS += -D AMD64
#     else
#         ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
#             CCFLAGS += -D AMD64
#         endif
#         ifeq ($(PROCESSOR_ARCHITECTURE),x86)
#             CCFLAGS += -D IA32
#         endif
#     endif
# else
#     UNAME_S := $(shell uname -s)
#     ifeq ($(UNAME_S),Linux)
#         CCFLAGS += -D LINUX
#     endif
#     ifeq ($(UNAME_S),Darwin)
#         CCFLAGS += -D OSX
#     endif
#     UNAME_P := $(shell uname -p)
#     ifeq ($(UNAME_P),x86_64)
#         CCFLAGS += -D AMD64
#     endif
#     ifneq ($(filter %86,$(UNAME_P)),)
#         CCFLAGS += -D IA32
#     endif
#     ifneq ($(filter arm%,$(UNAME_P)),)
#         CCFLAGS += -D ARM
#     endif
# endif

SHELL := /usr/local/bin/bash

default: deploy test prune-old-versions

deploy:
	gcloud app deploy --quiet

test:
	hey -z 3s http://jameslucktaylor.info

prune-old-versions:
	gcloud app versions delete $(shell gcloud app versions list --format="json" | jq -r '[ .[] | select(.traffic_split == 0) | .id ] | join(" ")') --quiet

clean: prune-old-versions
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
