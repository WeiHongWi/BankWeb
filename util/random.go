package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = `abcdefghijklmnopqrstuvwxyz`

// Random Seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random Generate int64
func Random_int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random Generate Alpha
func Random_string() string {
	var random_len int64 = Random_int64(3, 9)
	var tmp strings.Builder

	var i int64
	for i = 0; i < random_len; i++ {
		var pos int64 = Random_int64(0, 25)
		char := alphabet[pos]
		tmp.WriteByte(char)
	}

	return tmp.String()
}

func Random_owner() string {
	return Random_string()
}

func Random_money() int64 {
	return Random_int64(0, 1000)
}

func Random_currency() string {
	currency_list := []string{"NT", "USD", "EUR"}

	var tmp int64 = Random_int64(0, 2)
	return currency_list[tmp]
}
