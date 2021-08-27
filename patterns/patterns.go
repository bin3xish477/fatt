package patterns

var Patterns = map[string]string{
	`HttpUrl`:                `\bhttps?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)\b$`,
	`AwsArn`:                 `\barn:([^:\n]*):([^:\n]*):([^:\n]*):([^:\n]*):(([^:\/\n]*)[:\/])?(.*)\b$`,
	`AwsS3Bucket`:            `([a-z0-9.-]+\\.s3\\.amazonaws\\.com|//s3\\.amazonaws\\.com/[a-z0-9._-]+)`,
	`AwsIamAccessId`:         `AKIA[0-9A-Z]{16}`,
	`AwsSecretKey`:           `(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]`,
	`GitRepo`:                `\b(https?://)?github\.com(?:\/[^\s\/]+){2}\b$`,
	`Date`:                   `\b\d{2}[/-]\d{2}[/-]\d{2,4}\b$`,
	`Jwt`:                    `\beyJ[A-Za-z0-9-_=]+\.eyJ[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*\b$`,
	`UsAddress`:              `\b(\d+ [^,]+, [A-Z]{2} \d{5})\b$`,
	`IpAddrV4`:               `\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b$`,
	`IpAddrV6`:               `\b(?<![:.\w])(?:[A-F0-9]{1,4}:){7}[A-F0-9]{1,4}(?![:.\w])\b$`,
	`MacAddress`:             `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`,
	`Sid`:                    `\bS-\d-\d+-(\d+-){1,14}\d+\b$`,
	`Md5`:                    `\b[a-f0-9]{32}\b$`,
	`Sha256`:                 `\b[A-Fa-f0-9]{64}\b$`,
	`Sha512`:                 `\b[A-Fa-f0-9]{128}\b$`,
	`EmailAddr`:              `^\S+@\S+\.\S+$`,
	`PhoneNumber`:            `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$`,
	`Ssn`:                    `^(?!666|000|9\d{2})\d{3}-(?!00)\d{2}-(?!0{4})\d{4}$`,
	`SlackApiKey`:            `xox[baprs]-[0-9]{12}-[0-9]{12}-[0-9a-zA-Z]{24}`,
	`SlackAccessToken`:       `T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	`SlackWebHook`:           `https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`,
	`MailGunApiKey`:          `[0-9a-f]{32}-us[0-9]{1,2}`,
	`MailChimpKey`:           `[0-9a-f]{32}-us[0-9]{1,2}`,
	`GoogleApiKey`:           `AIza[0-9A-Za-z\\-_]{35}`,
	`FacebookToken`:          `EAACEdEose0cBA[0-9A-Za-z]+`,
	`Version`:                `\b(v(ersion)*\s*)([0-9]+)\.([0-9]+)\.([0-9]+)(?:\.([0-9]+))?\b$`,
	`LinuxPath`:              `^(/[^/ ]*)+/?$`,
	`WindowsPath`:            `^(?:[a-zA-Z]\:|\\\\[\w\.]+\\[\w.$]+)\\(?:[\w]+\\)*\w([\w.])+$`,
	`TwilioApiKey`:           `^SK[0-9a-fA-F]{32}$`,
	`Guid`:                   `^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$`,
	`Coordinates`:            `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`,
	`SquareAccessToken`:      `sqOatp-[0-9A-Za-z\\-_]{22}`,
	`SquareOauthSecret`:      `sq0csp-[ 0-9A-Za-z\\-_]{43}`,
	`AzureSubscriptionId`:    `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`,
	`AzureStorageKey`:        `AccountKey=[a-zA-Z0-9+\/=]{88}`,
	`WapBssid`:               `([0-9A-F]{2}([:-]|$)){6}`,
	`WindowsVersion`:         `(Windows (7|8(\.1)?|3(\.1|\.0)|10|ME|XP|Vista|95|98|2000|1\.0[1-4]|2\.(03|1[0|1])))$`,
	`EnvironmentVariable`:    `^[A-Z_]+=[\w\d]+\s?$`,
	`NetworkShare`:           `^(\\)(\\[\w\.-_]+){2,}(\\?)$`,
	`GitHubSecret`:           `github.*['|\"][0-9a-zA-Z]{35,40}['|\"]`,
	`HerokuKeys`:             `heroku.*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`,
	`PayPalToken`:            `access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}`,
	`AsymmetricKeyHeader`:    `\\-\\-\\-\\-\\-BEGIN ((EC|PGP|DSA|RSA|OPENSSH) )?PRIVATE KEY( BLOCK)?\\-\\-\\-\\-\\-`,
	`AmsAuthToken`:           `amzn.mws.[0-9a-f]{8}-[0-9a-f]{4}-10-9a-f1{4}-[0-9a,]{4}-[0-9a-f]{12}`,
	`PicaticApiKey`:          `sk_live_[0-9a-z]{32}`,
	`TwitterAccessToken`:     `[1-9][ 0-9]+-[0-9a-zA-Z]{40}`,
	`AuthorizationBasic`:     `(?i)Authorization: Basic [a-zA-Z0-9_\-.=]+\s?`,
	`AuthorizationBearer`:    `(?i)Authorization: Bearer [a-zA-Z0-9_\-.=]+\s?`,
	`GenericApiKey`:          `[a|A][p|P][i|I][_]?[k|K][e|E][y|Y].*['|"][0-9a-zA-Z]{32,45}['|"]$`,
	`WindowsRegistryKey`:     `(?i)^(HKEY_LOCAL_MACHINE|HKLM)([a-zA-Z0-9\s_@-\^!#.:/$%&+={}\[\]*])+$`,
	`VariableDeclaration`:    `\s?[\w_-]+\s?:?[=]{1}\s?["']?[\d\w\s]+["']?\s?`,
	`HTMLComment`:            `<!--(.*?)-->`,
	`AwsS3AccessLog`:         `(\S+) (\S+) \[(.*?)\] (\S+) (\S+) (\S+) (\S+) (\S+) "([^"]+)" (\S+) (\S+) (\S+) (\S+) (\S+) (\S+) "([^"]+)" "([^"]+)"`,
	`ApacheLog`:              `^(\S+) (\S+) (\S+) \[([\w:/]+\s[+\-]\d{4})\] "(\S+)\s?(\S+)?\s?(\S+)?" (\d{3}|-) (\d+|-)\s?"?([^"]*)"?\s?"?([^"]*)?"?$`,
	`WindowsIISLog`:          `^(\d{4}\-\d{2}\-\d{2}\s+)(\d{2}\:\d{2}\:\d{2}\s+)(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\s+)(.+?\s+)(.+?\s+)(.+?\s+)(\d{1,3}\s+)(.+?\s+)(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\s+)(.+?\s+)(\d{1,3}\s+)(\d{1,3}\s+)(\d+\s+)(\d+\s+)(\d+\s+)(\d+)?$`,
	`SyslogRfc3164`:          `^\<([0-9]{1,3})\>([A-Za-z]{3} [0-9]{1,2} \d{2}:\d{2}:\d{2}) ([\S]+) ([\S\s]+)`,
	`SyslogRfc5424`:          `^<(\d|\d{2}|1[1-8]\d|19[01])>(\d{1,2})\s(-|([12]\d{3})-(0\d|[1][012])-([012]\d|3[01])T([01]\d|2[0-4]):([0-5]\d):([0-5]\d|60)(?:\.(\d{1,6}))?(Z|[+-]\d{2}:\d{2}))\s([\S]{1,255})\s([\S]{1,48})\s([\S]{1,128})\s([\S]{1,32})\s(-|(?:\[.+?(?<!\\)\])+)(?:\s(.+))?$`,
	`NginxLog`:               `^([^ ]*) - ([^ ]*) \[(.*)\] \"([^ ]*) ([^ ]*) ([^ ]*)\" (-|[0-9]*) (-|[0-9]*) \"(.+?|-)\" ([^ ]*|-) ([^ ]*|-) ([^ ]*|-) \"(.+?|-)\" \"(.+?|-)\" \"(.+?|-)\"$`,
	`CVE`:                    `^CVE-\d{4}-\d{4,7}$`,
	`CVSSv2`:                 `(AV:([LAN])\/AC:([HML])\/Au:([NSM])\/C:([NPC])\/I:([NPC])\/A:([NPC])(?:\/E:(ND|U|POC|F|H)\/RL:(ND|OF|T|W|U)\/RC:(ND|UC|UR|C)(?:\/CDP:(N|L|LM|MH|H|ND)\/TD:(N|L|M|H|ND)\/CR:(L|M|H|ND)\/IR:(L|M|H|ND)\/AR:(L|M|H|ND))?)?)`,
	`CVSSv3`:                 `/AV:[NALP]/AC:[LH]/PR:[NLH]/UI:[NR]/S:[UC]/C:[NLH]/I:[NLH]/A:[NLH]`,
	`CVSSv3.1`:               `/AV:[NALP]/AC:[LH]/PR:[NLH]/UI:[NR]/S:[UC]/C:[NLH]/I:[NLH]/A:[NLH]/E:[F|H|U|P|X]/RL:[W|U|T|O|X]/RC:[C|R|U|X]/CR:[X|L|M|H]/IR:[X|L|M|H]/AR:[X|L|M|H]/MAV:[X|N|A|L|P]/MAC:[X|L|H]/MPR:[X|N|L|H]/MUI:[X|N|R]/MS:[X|U|C]/MC:[X|N|L|H]/MI:[X|N|L|H]/MA:[X|N|L|H]`,
	`PowerShellbase64Hidden`: `(?i)powershell.exe.*Hidden.*Enc\s[\w]+\s?`,
	`CmdExec`:                `(?i)^cmd\.exe \\c\s[\w\d"'*&\^%$#@!(),.\/+=|\\{}<>:; ]+\s?`,
	`TwitterHandle`:          `@([A-Za-z0-9_]{1,15})$`,
	`FacebookPageUrl`:        `(?:http:\/\/)?(?:www\.)?facebook\.com\/(?:(?:\w)*#!\/)?(?:pages\/)?(?:[\w\-]*\/)*([\w\-]*)`,
	`HTTPRequest`:            `GET\s+([^?\s]+)((?:[?&][^&\s]+)*)\s+(HTTP/.*)`,
	`HostHeader`:             `Host: (.+)\r\n`,
	`HTML`:                   `^<.*>.*</.*>\s?`,
	`YoutubeVideo`:           `((https(?:s)?:\/\/)?(www\.)?)?(?:youtu\.be\/|youtube\.com\/(?:embed\/|v\/|watch\?v=|watch\?&v=))((?:\w|-){11})((?:\&|\?)\S*)?`,
	`Log4jLog`:               `(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2},\d{3}) (\S*) (\S*) (\S*):(\d*) - (.*)`,
	`AwsElbAccessLog`:        `^(\S+)\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s+([^"]+)\s+\S+\s+\S+\s+"([^"]+)\s+HTTP\/\d.\d"\s+"([^"]+)"\s+-\s+-$`,
	`AkamaiLog`:              `([\d\.]+)\s([\w\-]+)\s([\w\-]+)\s\[(\d+\/\w+\/\d+\:\d+\:\d+\:\d+\s\+\d+)\]\s"(.+?)"\s([\d\-]+)\s([\d\-]+)\s"(.+?)"\s"(.+?)"\s"(.+?)"\s"(.+?)"\s"(.+?)"`,
	`EthereumAddress`:        `^0x([a-zA-Z0-9]{40})$`,
	`PgpFingerPrint`:         `^(([a-fA-F0-9]{4}\ ){9}[a-fA-F0-9]{4})$`,
}
