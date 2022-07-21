package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

// hash function

func Hash_func(s string) string{
	md5 := MD5(s)
	tmp_int64 := MD5_to_decimal(md5)
	res := decimal_to_hash(tmp_int64)
	return res
}

func decimal_to_hash(tmp int64) string{
	var res string
	var tmp_mod int64
	codeDict := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")
	len := len(codeDict)
	for i:=0; i<10; i++{
		tmp_mod = tmp % int64(len)
		tmp = tmp / int64(len)
		res += string(codeDict[tmp_mod])
	}
	return res
}

func MD5_to_decimal(hexstr string)int64{
	var tmp int64

	len_str := len(hexstr) / 2 - 1
	big_int := big.NewInt(0)
	big_tmp,_ := big_int.SetString(hexstr[:len_str], 16)
	tmp = int64(big_tmp.Uint64())
	return tmp
}

func MD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}