package strings

var awsStrings = map[string]string{
	"iamUniqId": `(ABIA|ACCA|AGPA|AIDA|AIPA|AKIA|ANPA|ANVA|APKA|AROA|ASCA|ASIA)[A-Z0-9]{16,128}`,
	"arn":       `arn:([^:\n]*):([^:\n]*):([^:\n]*):([^:\n]*):(([^:\/\n]*)[:\/])?[^ \n]+`,
}

var piiStrings = map[string]string{
	"fullName": "[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*",
}

var http_strings = map[string]string{
	"httpUrl":    `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`,
	"nonHttpUrl": `[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`,
}
