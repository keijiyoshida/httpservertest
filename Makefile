GOOSS := darwin linux
GOARCHES := 386 amd64
PACKAGES := cmd/httpservertest httpservertest

compile:
	@VERSION=`go run cmd/httpservertest/main.go -v | awk '{split($$0,a," "); print a[2]}'`; \
	if [ -e ./out ]; then \
		rm -rf ./out; \
	fi; \
	mkdir ./out; \
	for GOOS in ${GOOSS}; do \
		for GOARCH in ${GOARCHES}; do \
			mkdir out/httpservertest-$$VERSION-$$GOOS-$$GOARCH; \
			docker run --rm -it -v "$$PWD":/usr/local/go/src/github.com/keijiyoshida/httpservertest -w /usr/local/go/src/github.com/keijiyoshida/httpservertest golang:1.8 /bin/bash -c "GOOS=$$GOOS GOARCH=$$GOARC go build -o out/httpservertest-$$VERSION-$$GOOS-$$GOARCH/httpservertest ./cmd/httpservertest"; \
			tar -zcf out/httpservertest-$$VERSION-$$GOOS-$$GOARCH.tar.gz -C out httpservertest-$$VERSION-$$GOOS-$$GOARCH; \
			rm -rf out/httpservertest-$$VERSION-$$GOOS-$$GOARCH; \
		done \
	done

test:
	@for package in ${PACKAGES}; do \
		go test -cover -race ./$$package; \
	done
