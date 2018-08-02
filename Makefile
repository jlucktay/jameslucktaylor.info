default: deploy test

clean:
	rm -f jameslucktaylor.info_*.report.html

full: deploy validate test

deploy:
	gcloud app deploy --quiet

dev:
	dev_appserver.py app.yaml

test:
	hey -z 3s http://jameslucktaylor.info

validate: lighthouse-install
	open "https://validator.w3.org/unicorn/check?ucn_uri=jameslucktaylor.info"
	open "https://www.ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info&clearCache=on"
	open "https://realfavicongenerator.net/favicon_checker?protocol=https&site=jameslucktaylor.info"
	lighthouse https://jameslucktaylor.info --view

zap:
	/Applications/OWASP\ ZAP.app/Contents/Java/zap.sh -cmd -quickurl http://jameslucktaylor.info

lighthouse-install:
	npm update -g lighthouse
