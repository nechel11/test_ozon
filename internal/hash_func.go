package internal

import (
	"crypto/md5"
	"encoding/hex"
)

// hash function
func hash_func(s string) string{
	md5 := MD5_cutter_len10(MD5(s))
	specified_string := string_to_spec(md5)
	return specified_string
}

//MD5 string to specified string
func string_to_spec(s string) string{
	codeDict := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")
	var nearest_prime rune  = 67 // nearest prime to len(codeDict)
	len := len(codeDict)
	var res string
	for _,v:=range(s){
		tmp := (v * nearest_prime) % rune(len)
		res = res + string(codeDict[tmp])
	}
	return res
}

// MD5 string to MD5 len10 string
func MD5_cutter_len10(s string) string {	
	var tmp string
	for i,v := range(s[2:]){
		if i % 3 == 0{
			tmp = tmp + string(v)
		}
		i++
	}
	return tmp
}

// short url to MD5 string
func MD5(text string) string {
    algorithm := md5.New()
    algorithm.Write([]byte(text))
    return hex.EncodeToString(algorithm.Sum(nil))
}