package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"../guid"
	"fmt"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

func GenRandKey(keylen int) string {
	//return "0123456789ABCDEFFEDCBA9876543210"
	
	uuid ,err := guid.NewV4()
	if err !=nil{
		return ""
	}
	return fmt.Sprintf("%x", uuid[0:keylen])
}

