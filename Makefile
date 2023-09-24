build:
	go build -o bin/wcgo cmd/wcgo/main.go

clean:
	rm -rf bin/

test:
	cd cmd/wcgo/ && pwd && go test -v