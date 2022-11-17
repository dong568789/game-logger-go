package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
	"io"
	"sort"
)

func MakeSign(params map[string]string) string {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var str string
	for _, item := range keys {
		str += params[item] + "&"
	}
	return generatorMD5(str + viper.GetString("secret"))
}

func generatorMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
