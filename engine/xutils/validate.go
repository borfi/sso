package xutils

import (
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	urlRegexString      = "^/[a-zA-Z0-9-_/]{0,100}$" // 总长度限100个字符
	emailRegexString    = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:\\(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22)))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	telRegexString      = "^1[3456789][0-9]{9}$"
	zipcodeRegexString  = "^\\d{6}$"
	digitRegexString    = "\\d"
	certCodeRegexString = `^(\d{18,18}|\d{15,15}|\d{17,17}(x|X))$`
)

var (
	urlRegex      = regexp.MustCompile(urlRegexString)
	emailRegex    = regexp.MustCompile(emailRegexString)
	telRegex      = regexp.MustCompile(telRegexString)
	zipcodeRegex  = regexp.MustCompile(zipcodeRegexString)
	digitRegex    = regexp.MustCompile(digitRegexString)
	certCodeRegex = regexp.MustCompile(certCodeRegexString)
)

// IsEmail 是否为email
func IsEmail(str string) bool {
	return emailRegex.MatchString(str)
}

// IsTel 是否为手机号
func IsTel(str string) bool {
	return telRegex.MatchString(str)
}

// IsZipcode 是否为邮编
func IsZipcode(str string) bool {
	return zipcodeRegex.MatchString(str)
}

// IsIP 是否为ipv6
func IsIP(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil
}

// IsIPV4 是否为ipv4
func IsIPV4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() != nil
}

// IsURL 是否为合法url
func IsURL(str string) bool {
	if i := strings.Index(str, "#"); i > -1 {
		str = str[:i]
	}
	if len(str) == 0 {
		return false
	}

	parseURL, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	if parseURL.Scheme != "http" && parseURL.Scheme != "https" {
		return false
	}
	return err == nil
}

// IsURLPath 是否为url path(例如 /test/test、 /、 /test)
func IsURLPath(str string) bool {
	if str == "" {
		return false
	}

	if str[0] != '/' {
		return false
	}

	if !urlRegex.MatchString(str) {
		return false
	}

	return true
}

// IsPersonName 是否为合法人名(不能包含数字)
func IsPersonName(str string) bool {
	return !digitRegex.MatchString(str)
}

// IsCertCode 是否为身份证号
func IsCertCode(str string) bool {
	return certCodeRegex.MatchString(str)
}

// HasSpecialChar 是否包含特殊字符
func HasSpecialChar(str string) bool {
	if !utf8.Valid([]byte(str)) {
		return true
	}
	for _, ru := range str {
		if utf8.RuneLen(ru) > 3 {
			return true
		}
	}
	return false
}
