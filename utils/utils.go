package utils

import "net/url"

func EncodeMessage(message string) string {
	return url.QueryEscape(message)
}
