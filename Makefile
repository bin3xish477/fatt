windows:
	GOOS=windows GOARCH=amd64 go build -o ./fatt.exe ./cmd/fatt/*
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./fatt.exe ./cmd/fatt/*
	tar -czvf fatt-windows-amd64.tar.gz ./fatt.exe

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
darwin-release:
	GOOS=windows GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
	tar -czvf fatt-darwin-amd64.tar.gz ./fatt

linux:
	GOOS=linux GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
linux-release:
	GOOS=windows GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
	tar -czvf fatt-linux-amd64.tar.gz ./fatt
