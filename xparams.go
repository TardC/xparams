package xparams

import (
	"regexp"
)

type Extractor []*regexp.Regexp

func (e Extractor) Extract(s string, location string) []*Parameter {
	var ps []*Parameter

	for _, regExp := range e {
		result := regExp.FindAllStringSubmatchIndex(s, -1)
		for _, a := range result {
			p := &Parameter{
				location:        location,
				rawStr:          s,
				key:             s[a[2]:a[3]],
				value:           s[a[4]:a[5]],
				keyStartIndex:   a[2],
				keyEndIndex:     a[3],
				valueStartIndex: a[4],
				valueEndIndex:   a[5],
			}
			ps = append(ps, p)
		}
	}

	return ps
}

var (
	DefaultQueryFormExtractor Extractor
	DefaultCookieExtractor    Extractor
	DefaultJSONExtractor      Extractor
)

func init() {
	const (
		queryFormRegexpStr  = `(?U)([^&\s]+)=([^&\s]*)&`
		queryFormRegexpStr2 = `(?U)([^&\s]+)=([^&\s]*)$`

		cookieRegexpStr = `(\S+)=(\S*)\b`

		jsonStrRegexpStr     = `"([^{\[,]+)":"([^{\[,]+)"`
		jsonNumRegexpStr     = `"([^{\[,]+)":(-?[\d.]+)`
		jsonKeywordRegexpStr = `"([^{\[,]+)":(true|false|null)`
	)

	queryFormRegexp := regexp.MustCompile(queryFormRegexpStr)
	queryFormRegexp2 := regexp.MustCompile(queryFormRegexpStr2)
	cookieRegexp := regexp.MustCompile(cookieRegexpStr)
	jsonStrRegexp := regexp.MustCompile(jsonStrRegexpStr)
	jsonNumRegexp := regexp.MustCompile(jsonNumRegexpStr)
	jsonKeywordRegexp := regexp.MustCompile(jsonKeywordRegexpStr)

	DefaultQueryFormExtractor = Extractor{queryFormRegexp, queryFormRegexp2}
	DefaultCookieExtractor = Extractor{cookieRegexp}
	DefaultJSONExtractor = Extractor{jsonStrRegexp, jsonNumRegexp, jsonKeywordRegexp}
}

const (
	LocationQuery  = "query"
	LocationCookie = "cookie"
	LocationBody   = "body"
)

type Parameter struct {
	location        string // query, body, cookie. etc.
	rawStr          string
	key             string
	value           string
	keyStartIndex   int
	keyEndIndex     int
	valueStartIndex int
	valueEndIndex   int
}

func (p *Parameter) Replace(value string) string {
	return p.rawStr[:p.valueStartIndex] + value + p.rawStr[p.valueEndIndex:]
}

func (p *Parameter) Append(value string) string {
	return p.rawStr[:p.valueEndIndex] + value + p.rawStr[p.valueEndIndex:]
}
