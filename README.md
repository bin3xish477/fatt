# fatt

fatt (Find All The Things) is a tool written in Go that'll find common strings in a specified file or HTTP response. If you have any regular expressions that aren't included, please make a pull request containing the regular expression and I'll include it into the tools current collection. I've already began working on **fatt v1.1.0** which will include additional patterns, options, and capabilities.

https://user-images.githubusercontent.com/44281620/131411335-21117104-b2cf-4792-8102-132956411630.mp4

### Usage Examples
```bash
➤ cat /path/to/file.txt | fatt --workers 25 --outfile results.txt

➤ fatt --file /path/to/file.txt --workers 30 --outfile

➤ fatt -u http://someurl.com -o results.txt -n

➤ fatt -f /path/to/file.txt -x 'VariableDeclaration,HTML' -q
```
### Help Menu
```
Usage: main [--file FILE] [--url URL] [--outfile OUTFILE] [--workers WORKERS] [--nocolor] [--quiet] [--list] [--exclude EXCLUDE]

Options:
  --file FILE, -f FILE   file to scan
  --url URL, -u URL      url to scan
  --outfile OUTFILE, -o OUTFILE
                         name of directory to save results to
  --workers WORKERS, -w WORKERS
                         number of threads for scanning [default: 20]
  --nocolor, -n          turn off color output
  --quiet, -q            make output less verbose
  --list, -l             list all pattern names
  --exclude EXCLUDE, -x EXCLUDE
                         exclude a specific pattern from search (could specify a file or as comma-seperated values
  --help, -h             display this help and exit
```

### Installation

#### Run with Docker
```
➤ git clone https://github.com/binexisHATT/fatt.git
➤ docker build -t fatt .
➤ docker run --rm --name fatt fatt -h
```

#### With Go Installed (Go v1.16 or newer)
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
➤ wget https://github.com/binexisHATT/fatt/releases/download/1.1.0/fatt-linux-amd64.tar.gz
➤ wget https://github.com/binexisHATT/fatt/releases/download/1.1.0/fatt-linux-amd64.tar.gz.sha256sum

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

### Add Autotab Completion for fatt
```
source fatt-Completion.sh
```

### Change Log
```
fatt v1.1.0
---------------------------------------------------------------------------------------
Patterns Added:
	- AwsS3AccessLog: matches AWS S3 log entry
	- ApacheLog: matches Apache log entry
	- WindowsIIS: matches Windows IIS log entry 
	- SyslogRfc3164: matches syslog log entry for RFC3164
	- SyslogRfc5424: matches syslog log entry for RFC5425
	- NginxLog: matches an Nginx log entry
	- CVE: matches a CVE Id
	- CVSSv2: matches a CVSS v2 vector
	- CVSSv3: matches a CVSS v3 vector
	- CVSSv3.1: matches a CVSS v3.1 vector
	- PowerShellbase64Hidden: matches PowerShell base64 encoded and hidden command
	- CmdExec: matches a cmd.exe non-interactive command-line
	- HTMLComment: matches HTML comments
	- TwitterHandle: matches a twitter handle
	- FacebookPageUrl: matches a facebook page URL
	- HostHeader: matches HTTP Host header
	- AzureStorageKey: matches Azure storage access keys for a storage account
	- YoutubeVideo: matches a Youtube video URL
	- Log4jLog: matches a log4j log entry
	- AwsElbAccessLog: matches an AWS ELB access log entry
	- AkamaiLog: matches an Akamai log entry
	- NameSpace: matches a namespace definition
	- SePrivilege: matches an Windows SE privilege definition
	- UserAgent: matches generic User Agent string
	- SQLQuery: matches a SQL query of the form (SELECT ... FROM ... WHERE ...)
	- WordPressVersion: matches WordPress version
	- DrupalVersion: matches Drupal version
	- JoomlaVersion: matches Joomla version
	- MagentoVersion: matches Magento version
	- UmbracoVersion: matches Umbraco version
	- ImageFile: matches valid image file name
	- SQLiVars: matches potential SQL Injection variables
	- RCEVars: matches potential Remote Code Execution variables
	- IDORVars: matches potential Insecure Direct Object Reference variables

Options Added:
	- (-q | --quiet): less verbose output
	- (-l | --list): list all available patterns and count	
	- (-x | --exclude): exclude a specified pattern from search

Misc:
	- Added autotab Bash completion script
	- Added a Dockerfile
```
