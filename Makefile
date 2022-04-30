.PHONY: assets cli

assets:
	(cd ui && yarn build)
	go-bindata -o ui/assets.go -pkg ui -prefix ui/build ui/build

cli:
	GOOS=darwin go build -o build/sbitportal-mac github.com/SBit-Project/sbit-portal/cli/sbitportal
	GOOS=linux go build -o build/sbitportal-linux github.com/SBit-Project/sbit-portal/cli/sbitportal