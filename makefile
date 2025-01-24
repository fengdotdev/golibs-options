# for updating the version of the project and pushing the tag to the repository
OPTIONS_VERSION = 0.0.4

updatev:
		git tag v${OPTIONS_VERSION} && git push origin v${OPTIONS_VERSION}


fix:
	go mod tidy

updatemod:
	go get -u ./...


test:
	go test -v ./...