c:
	go build -o ./bin/matra ./cmd/cli/main.go
	GOOS=js GOARCH=wasm go build -o ./web/wasm/matra.wasm ./cmd/wasm/main.go
	cp "/usr/local/go/lib/wasm/wasm_exec.js" ./web/js/

s:
	python3 -m http.server 8000 -d ./web

r:
	./bin/matra -i assets/input.txt -o assets/output.txt
