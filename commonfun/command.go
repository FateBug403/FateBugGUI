package command

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Spiltdoamins 分解多级域名(比如将b.a.baidu分解为b.a.baidu和a.baidu,可以得到更多子域名)
func Spiltdoamins(subdomain string) []string {
	var subdomains []string
	tmp := subdomain //保存传递进来的域名，后面要用
	for {
		v := strings.SplitAfterN(subdomain, ".", 2)
		if strings.Count(v[1], ".") < 2 {
			break
		}
		subdomain = v[1]
		subdomains = append(subdomains, v[1])
	}
	//加上源处理的域名
	subdomains = append(subdomains, tmp)
	return subdomains
}

// 字符串切片去重
func RemoveDuplicateElement(value []string) []string {
	result := make([]string, 0, len(value))
	temp := map[string]struct{}{}
	for _, item := range value {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 正则匹配判断字符串是否为ip地址
func RegexIP(str string) bool {
	partIp := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIp + "\\." + partIp + "\\." + partIp + "\\." + partIp
	match, _ := regexp.MatchString(grammer, str)
	return match
}


//In 二分查找判断某个字符串是否在某个切片中
func In(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

//InInt 二分查找判断某个整形是否在某个切片中
func InInt(target int, int_array []int) bool {
	sort.Ints(int_array)
	index := sort.SearchInts(int_array,target)
	//index的取值：[0,len(str_array)]
	if index < len(int_array) && int_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}

//ReadFile 从指定文件中读取行
func ReadFile(path string) []string{
	file,err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	// 创建一个读取器
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// InfoContains 判断一个字符串是否包含子串（模糊匹配）
func InfoContains(infoValue, filterValue string) bool {
	return strings.Contains(strings.ToLower(infoValue), strings.ToLower(filterValue))
}

//RSA_Encrypt RSA加密
func RSA_Encrypt(plainText []byte, PublicKeyBase64 string) string {
	// 解码 Base64 编码的公钥
	publicKeyBytes, err := base64.StdEncoding.DecodeString(PublicKeyBase64)
	if err != nil {
		log.Println("解密错误")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		log.Println(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		log.Println(err)
	}
	//返回密文的base64加密
	return base64.StdEncoding.EncodeToString(cipherText)
}