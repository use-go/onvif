package goonvif

import (
	"encoding/xml"
	"time"
	"encoding/base64"
	"crypto/sha1"
	"github.com/elgs/gostrgen"
)

/*************************
	WS-Security types
*************************/
const (passwordType = "https://www.oasis-open.org/committees/download.php/13392/wss-v1.1-spec-pr-UsernameTokenProfile-01.htm#PasswordDigest")

type security struct {
	XMLName xml.Name  `xml:"wsse:Security"`
	Auth wsAuth
}

type password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type wsAuth struct {
	XMLName xml.Name  `xml:"wsse:UsernameToken"`
	Username string   `xml:"wsse:Username"`
	Password password `xml:"wsse:Password"`
	Nonce string      `xml:"wsse:Nonce"`
	Created string    `xml:"wsse:Created"`
}

func newSecurity(username, passwd string) security {
	/** Generating Nonce sequence **/
	charsToGenerate := 16
	charSet := gostrgen.Lower | gostrgen.Digit

	nonce, _ := gostrgen.RandGen(charsToGenerate, charSet, "", "")

	auth := security{
		Auth:wsAuth{
			Username:username,
			Password:password {
				Type:passwordType,
				Password:generateToken(username, nonce, time.Now(), passwd),
			},
			Nonce: nonce,
			Created: time.Now().Format(time.RFC3339),
		},
	}

	return auth
}

//Digest = B64ENCODE( SHA1( B64DECODE( Nonce ) + Date + Password ) )
func generateToken(Username string, Nonce string, Created time.Time, Password string) string {

	sDec, _ := base64.StdEncoding.DecodeString(Nonce)


	hasher := sha1.New()
	//hasher.Write([]byte((base64.StdEncoding.EncodeToString([]byte(Nonce)) + Created.Format(time.RFC3339) + Password)))
	hasher.Write([]byte(string(sDec) + Created.Format(time.RFC3339) + Password))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

