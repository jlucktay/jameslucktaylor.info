default: deploy test

full: deploy validate test

deploy:
	gcloud app deploy

dev:
	dev_appserver.py app.yaml

test:
	hey -z 3s http://jameslucktaylor.info

validate:
	open "https://validator.w3.org/unicorn/check?ucn_uri=jameslucktaylor.info"
	open "https://www.ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info&latest"
	open "https://realfavicongenerator.net/favicon_checker?protocol=https&site=jameslucktaylor.info"
