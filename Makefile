windows:
	GOOS=windows GOARCH=amd64 go build -o ./fatt.exe ./cmd/fatt/*
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./fatt.exe ./cmd/fatt/*
	tar -czvf fatt-windows-amd64.tar.gz ./fatt.exe ./README.md
	sha256sum fatt-windows-amd64.tar.gz > fatt-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./fatt.exe ./cmd/fatt/*
	tar -czvf fatt-windows-386.tar.gz ./fatt.exe ./README.md
	sha256sum fatt-windows-386.tar.gz > fatt-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
	tar -czvf fatt-darwin-amd64.tar.gz ./fatt ./README.md
	sha256sum fatt-darwin-amd64.tar.gz > fatt-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
linux-release:
	GOOS=windows GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
	tar -czvf fatt-linux-amd64.tar.gz ./fatt ./README.md
	sha256sum fatt-linux-amd64.tar.gz > fatt-linux-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./fatt ./cmd/fatt/*
	tar -czvf fatt-linux-386.tar.gz ./fatt ./README.md
	sha256sum fatt-linux-386.tar.gz > fatt-linux-386.tar.gz.sha256sum

clean:
	@if [ -f fatt ]; then rm fatt; fi
	@if [ -f fatt.exe ]; then rm fatt.exe; fi
