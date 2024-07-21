package main

import (
	"dtcp/utils"
)

func Genbuf(seed string, n int) string {
	var buf string
	for i := 0; i < n; i++ {
		buf = buf + seed
	}
	return buf
}

func Genrandbuf(n int) string {
	return utils.RandStringRunes(n)
}
