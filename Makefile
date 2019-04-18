examples: xmpl xmpl-feat

xmpl:
	go generate ./example/main.go
	mkdir build
	cpp -P example/main.pgo build/main.go
	go build -o xmpl ./build
	rm -r build

xmpl-feat:
	go generate ./example/main.go
	mkdir build
	cpp -DFEATURE -P example/main.pgo build/main.go
	cp example/feature.go build/
	go build -tags feature -o xmpl-feat ./build
	rm -r build

.PHONY: clean
clean:
	-rm -r build
	-rm xmpl
	-rm xmpl-feat