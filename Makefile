default: deploy test

deploy:
	gcloud app deploy

dev:
	dev_appserver.py app.yaml

test:
	hey -z 3s http://jameslucktaylor.info
