package strings

var Patterns = map[string]string{
	`httpUrl`:       `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`,
	`awsArn`:        `arn:([^:\n]*):([^:\n]*):([^:\n]*):([^:\n]*):(([^:\/\n]*)[:\/])?(.*)`,
	`gitRepo`:       `(https?://)?github\.com(?:\/[^\s\/]+){2}`,
	`date`:          `\d{2}[/.-]\d{2}[/.-]\d{2,4}`,
	`jwt`:           `[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*`,
	`usLocation`:    `(\d+ [^,]+, [A-Z]{2} \d{5})`,
	`privSSHKey`:    `-----BEGIN PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END PRIVATE KEY——`,
	`privateRSAKey`: `-----BEGIN RSA PRIVATE KEY-----[a-zA-Z0-9\\S]{100,}-----END RSA PRIVATE KEY-----`,
	`password`:      `((?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[-^+!&@#$%]).{8,40})`,
	`ipAddrV4`:      `\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}\b`,
	`ipAddrV6`:      `(?<![:.\w])(?:[A-F0-9]{1,4}:){7}[A-F0-9]{1,4}(?![:.\w])`,
	`sid`:           `S-\d-\d+-(\d+-){1,14}\d+`,
}
