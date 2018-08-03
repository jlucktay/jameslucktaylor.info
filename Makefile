default: deploy test

clean:
	rm -fv jameslucktaylor.info_*.report.html

full: deploy validate-web validate-lighthouse test clean

validate: validate-web validate-lighthouse test clean

kitchen-sink: deploy validate-web validate-lighthouse test zap clean

deploy:
	gcloud app deploy --quiet

dev:
	dev_appserver.py app.yaml

test:
	hey -z 3s http://jameslucktaylor.info

validate-lighthouse: lighthouse-install
	lighthouse https://jameslucktaylor.info --view

validate-web:
	open "https://validator.w3.org/unicorn/check?ucn_uri=jameslucktaylor.info"
	open "https://www.ssllabs.com/ssltest/analyze.html?d=jameslucktaylor.info&clearCache=on"
	open "https://realfavicongenerator.net/favicon_checker?protocol=https&site=jameslucktaylor.info"

zap:
	/Applications/OWASP\ ZAP.app/Contents/Java/zap.sh -cmd -quickurl http://jameslucktaylor.info
	# TODO: send output through parser (https://yq.readthedocs.io/en/latest/) and make it more succinct

lighthouse-install:
	npm update -g lighthouse
