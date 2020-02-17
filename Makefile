build:
	env GOOS=linux go build -ldflags="-s -w" -o handler/postSurvey main.go

sls_deploy: build
	serverless deploy

deploy: build sls_deploy
