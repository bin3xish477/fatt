windows:
	GOOS=windows GOARCH=amd64 go build -o ./fatt.exe ./cmd/fatt/*
darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
linux:
	GOOS=linux GOARCH=amd64 go build -o ./fatt ./cmd/fatt/*
