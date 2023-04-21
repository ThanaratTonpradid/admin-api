package helper

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var (
	regExpSpace *regexp.Regexp
)

func init() {
	regExpSpace = regexp.MustCompile(`[\s\xa0]+`)
}

func RemoveDuplicateSpaces(v string) string {
	return strings.TrimSpace(regExpSpace.ReplaceAllString(v, " "))
}

func RemoveAllSpaces(v string) string {
	return strings.TrimSpace(regExpSpace.ReplaceAllString(v, ""))
}

func ToString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func ToFloat64(v string) (float64, error) {
	v = strings.ReplaceAll(v, ",", "")
	return strconv.ParseFloat(v, 64)
}

func UUID() string {
	return uuid.New().String()
}

func UUIDWithoutHyphen() string {
	return strings.Replace(UUID(), "-", "", -1)
}

func GetHostDomain(host string) string {
	rawUrl := fmt.Sprintf("https://%s", host)
	u, err := url.Parse(rawUrl)
	if err != nil {
		return host
	}
	parts := strings.Split(u.Hostname(), ".")
	count := len(parts)
	if count < 3 {
		return u.Hostname()
	}
	return fmt.Sprintf("%s.%s", parts[count-2], parts[count-1])
}

func GetHostURL(v string) string {
	u, err := url.Parse(v)
	if err != nil {
		return v
	}
	return strings.Replace(v, u.Path, "", 1)
}
