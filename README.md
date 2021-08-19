# fatt
fatt (Find All The Things) is a tool written in Go that'll find common strings in a specified file or HTTP response. If you have any regular expressions that aren't included, please make a pull request containing the regular expression and I'll include it into the tools current collection. I've already began working on **fatt v1.1.0** which will include additional patterns, options, and capabilities.

https://user-images.githubusercontent.com/44281620/130001712-ebf962c6-fd7a-4fe0-94b8-b7cb354ada21.mp4

### Usage
```bash
➤ cat /path/to/file.txt | fatt --workers 25 --outfile results.txt

➤ fatt --file /path/to/file.txt --workers 30 --outfile

➤ fatt -u http://someurl.com -o results.txt -n
```

### Help Menu
```
Usage: main [--file FILE] [--url URL] [--outfile OUTFILE] [--workers WORKERS] [--nocolor]

Options:
  --file FILE, -f FILE   file to scan
  --url URL, -u URL      url to scan
  --outfile OUTFILE, -o OUTFILE
                         name of directory to save results to
  --workers WORKERS, -w WORKERS
                         number of threads for scanning [default: 20]
  --nocolor, -n          turn off color output
  --help, -h             display this help and exit
```

### Installation
#### With Go Installed (Go v1.17 or newer)
```
➤ go get -u -v github.com/binexisHATT/fatt/cmd/fatt
➤ fatt --help
```
#### From Source (Go Must Be Installed)
```
➤ git clone https://github.com/binexisHATT/fatt.git
➤ cd fatt
➤ go build -o fatt cmd/fatt/*
```

#### From Release
```
1. Go to fatt's release page: https://github.com/binexisHATT/fatt/releases
2. Download the .tar.gz archive for the target OS/Architecture and the .sha256sum file for verification
3. Run sha256sum --check to verify download hash
4. Extract the archive with the tar command to obtain binary

# Example for Linux (amd64):
➤ wget https://github.com/binexisHATT/fatt/releases/download/1.0.0/fatt-linux-amd64.tar.gz
➤ wget https://github.com/binexisHATT/fatt/releases/download/1.0.0/fatt-linux-amd64.tar.gz.sha256sum

# Verify SHA256 Hash
➤ sha256sum --check fatt-linux-amd64.tar.gz.sha256sum
# Output if hashes matched:
fatt-linux-amd64.tar.gz: OK

# Extract tar.gz Archive
➤ tar -xvzf fatt-linux-amd64.tar.gz
# Install fatt to PATH
➤ sudo install ./fatt /usr/local/bin/fatt
# OR
➤ sudo mv ./fatt /usr/local/bin/
```

