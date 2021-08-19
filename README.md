# fatt
fatt (Find All The Things) is a tool written in Go that'll find common strings in a specified file or HTTP response. If you have any regular expressions that aren't included, please make a pull request containing the regular expression and I'll include it into the tools.

https://user-images.githubusercontent.com/44281620/130001712-ebf962c6-fd7a-4fe0-94b8-b7cb354ada21.mp4

### Usage
```bash
cat /path/to/file.txt | fatt --workers 25 --outfile results.txt
fatt --file /path/to/file.txt --workers 30 --outfile
# fatt can also find strings in HTTP responses
fatt --url http://someurl.com --outfile results.txt --nocolor
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
#### From Source (Go Must Be Installed)
```
git clone https://github.com/binexisHATT/fatt.git
cd fatt
go build -o fatt cmd/fatt/*
```

#### From Release
```
1. Go to fatt's release page: https://github.com/binexisHATT/fatt/releases
2. Download the .tar.gz archive for the target OS/Architecture and the .sha256sum file for verification
3. Run sha256sum --check to verify download hash
4. Extract the archive with the tar command to obtain binary

# Example for Linux (amd64):
wget https://github.com/binexisHATT/fatt/releases/download/1.0.0/fatt-linux-amd64.tar.gz
wget https://github.com/binexisHATT/fatt/releases/download/1.0.0/fatt-linux-amd64.tar.gz.sha256sum

# Verify SHA256 Hash
sha256sum --check fatt-linux-amd64.tar.gz.sha256sum
# Output if hashes matched:
fatt-linux-amd64.tar.gz: OK

# Extract tar.gz Archive
tar -xvzf fatt-linux-amd64.tar.gz
# Install fatt to PATH
sudo install ./fatt /usr/local/bin/fatt

# OR sudo mv ./fatt /usr/local/bin/
```

### Current Regex Patterns Supported
```json
package patterns

var Patterns = map[string]string{
	`HttpUrl`:             `\bhttps?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)\b$`,
	`AwsArn`:              `\barn:([^:\n]*):([^:\n]*):([^:\n]*):([^:\n]*):(([^:\/\n]*)[:\/])?(.*)\b$`,
	`AwsS3Bucket`:         `([a-z0-9.-]+\\.s3\\.amazonaws\\.com|//s3\\.amazonaws\\.com/[a-z0-9._-]+)`,
	`AwsIamAccessId`:      `AKIA[0-9A-Z]{16}`,
	`AwsSecretKey`:        `(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]`,
	`GitRepo`:             `\b(https?://)?github\.com(?:\/[^\s\/]+){2}\b$`,
	`Date`:                `\b\d{2}[/-]\d{2}[/-]\d{2,4}\b$`,
	`Jwt`:                 `\beyJ[A-Za-z0-9-_=]+\.eyJ[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*\b$`,
	`UsAddress`:           `\b(\d+ [^,]+, [A-Z]{2} \d{5})\b$`,
	`IpAddrV4`:            `\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b$`,
	`IpAddrV6`:            `\b(?<![:.\w])(?:[A-F0-9]{1,4}:){7}[A-F0-9]{1,4}(?![:.\w])\b$`,
	`MacAddress`:          `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`,
	`Sid`:                 `\bS-\d-\d+-(\d+-){1,14}\d+\b$`,
	`Md5`:                 `\b[a-f0-9]{32}\b$`,
	`Sha256`:              `\b[A-Fa-f0-9]{64}\b$`,
	`Sha512`:              `\b[A-Fa-f0-9]{128}\b$`,
	`Name`:                `\b([a-zA-Z]{2,}\s[a-zA-Z]{1,}'?-?[a-zA-Z]{2,}\s?([a-zA-Z]{1,})?)\b$`,
	`EmailAddr`:           `^\S+@\S+\.\S+$`,
	`PhoneNumber`:         `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`,
	`Ssn`:                 `^(?!666|000|9\d{2})\d{3}-(?!00)\d{2}-(?!0{4})\d{4}$`,
	`SlackApiKey`:         `xox[baprs]-[0-9]{12}-[0-9]{12}-[0-9a-zA-Z]{24}`,
	`SlackAccessToken`:    `T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	`SlackWebHook`:        `https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	`MailGunApiKey`:       `[0-9a-f]{32}-us[0-9]{1,2}`,
	`MailChimpKey`:        `[0-9a-f]{32}-us[0-9]{1,2}`,
	`GoogleApiKey`:        `AIza[0-9A-Za-z\\-_]{35}`,
	`FacebookToken`:       `EAACEdEose0cBA[0-9A-Za-z]+`,
	`Version`:             `\b(v(ersion)*\s*)([0-9]+)\.([0-9]+)\.([0-9]+)(?:\.([0-9]+))?\b$`,
	`LinuxPath`:           `^(/[^/ ]*)+/?$`,
	`WindowsPath`:         `^(?:[a-zA-Z]\:|\\\\[\w\.]+\\[\w.$]+)\\(?:[\w]+\\)*\w([\w.])+$`,
	`TwilioApiKey`:        `^SK[0-9a-fA-F]{32}$`,
	`Password`:            `(pass(word|phrase)?|secret): \S+`,
	`UserName`:            `(user( ?name)?|login): \S+`,
	`Guid`:                `^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$`,
	`BtcAddress`:          `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`,
	`CreditCard`:          `(?:(?:(?:\d{4}[- ]?){3}\d{4}|\d{15,16}))`,
	`McCreditCard`:        `5[1-5]\d{2}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`,
	`VisaCreditCard`:      `4\d{3}[\s-]?\d{4}[\s-]?\d{4}[\s-]?\d{4}`,
	`Coordinates`:         `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`,
	`SquareAccessToken`:   `sqOatp-[0-9A-Za-z\\-_]{22}`,
	`SquareOauthSecret`:   `sq0csp-[ 0-9A-Za-z\\-_]{43}`,
	`AzureSubscriptionId`: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`,
	`WapBssid`:            `([0-9A-F]{2}([:-]|$)){6}`,
	`WindowsVersion`:      `(Windows (7|8(\.1)?|3(\.1|\.0)|10|ME|XP|Vista|95|98|2000|1\.0[1-4]|2\.(03|1[0|1])))$`,
	`EnvironmentVariable`: `^([A-Z]+=\w+)$`,
	`NetworkShare`:        `^(\\)(\\[\w\.-_]+){2,}(\\?)$`,
	`GitHubSecret`:        `github.*['|\"][0-9a-zA-Z]{35,40}['|\"]`,
	`HerokuKeys`:          `heroku.*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`,
	`PayPalToken`:         `access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}`,
	`AsymmetricKeyHeader`: `\\-\\-\\-\\-\\-BEGIN ((EC|PGP|DSA|RSA|OPENSSH) )?PRIVATE KEY( BLOCK)?\\-\\-\\-\\-\\-`,
	`AmsAuthToken`:        `amzn.mws.[0-9a-f]{8}-[0-9a-f]{4}-10-9a-f1{4}-[0-9a,]{4}-[0-9a-f]{12}`,
	`PicaticApiKey`:       `sk_live_[0-9a-z]{32}`,
	`TwitterAccessToken`:  `[1-9][ 0-9]+-[0-9a-zA-Z]{40}`,
	`AuthorizationBasic`:  `basic [a-zA-Z0-9_\-\:\.=]+`,
	`AuthorizationBearer`: `bearer [a-zA-Z0-9_\\-\\.=]+`,
	`GenericApiKey`:       `[a|A][p|P][i|I][_]?[k|K][e|E][y|Y].*['|\"][0-9a-zA-Z]{32,45}['|\"]`,
	`WindowsRegistryKey`:  `(?i)^(HKEY_LOCAL_MACHINE|HKLM)([a-zA-Z0-9\s_@-\^!#.:/$%&+={}\[\]*])+$`,
	`VariableDeclaration`: `[\w_-]+\s?:?=\s?["'][\d\w\s]+["']`,
	`Comment`:             `((/\*([^*]|(\*+[^*/]))*\*+/)|(//.*)|^#\s[\w\s]+\n)`,
}
```
