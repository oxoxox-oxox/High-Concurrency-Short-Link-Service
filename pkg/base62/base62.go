package base62

import "strings"

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(id uint64) string {
	if id == 0 {
		return string(chars[0])
	}

	var sb strings.Builder
	for id > 0 {
		remainder := id % 62
		sb.WriteByte((chars[remainder]))
		id = id / 62
	}

	runes := []rune(sb.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}
