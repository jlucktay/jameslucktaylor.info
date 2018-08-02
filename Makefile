default: deploy test

full: deploy test validate

deploy:
	gcloud app deploy

dev:
	dev_appserver.py app.yaml

test:
	hey -z 3s http://jameslucktaylor.info

validate:
	open "https://validator.w3.org/unicorn/check?ucn_uri=jameslucktaylor.info"
	open "https://realfavicongenerator.net/favicon_checker?protocol=https&site=jameslucktaylor.info"
