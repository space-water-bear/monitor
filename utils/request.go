package utils

import (
	"bytes"
	"clients/pkg/errno"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"sort"
)

func sortMapKey(m map[string]interface{}) []string {

	var newMap = make([]string, 0)
	for k, _ := range m {
		newMap = append(newMap, k)
	}
	sort.Strings(newMap)
	return newMap
}

func GetMD5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodeMD5AndBase64(str string) string {
	return GetMD5Encode(base64.StdEncoding.EncodeToString([]byte(str)))
}

func pushData(data map[string]interface{}, url string) error {
	var sign string

	data["ip"] = viper.GetString("public_ip")

	mp := sortMapKey(data)
	for _, v := range mp {
		sign += fmt.Sprintf(`&%v=%v`, v, data[v])
		//fmt.Println(v, data[v])
	}
	key := EncodeMD5AndBase64(sign[1:])
	jdata, err := json.Marshal(data)
	if err != nil {
		return errno.ErrEncodeError
	}
	baseUrl := viper.GetString(`center`)
	req, _ := http.NewRequest("POST", baseUrl+url, bytes.NewBuffer(jdata))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("keys", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errno.InternalServerError
	}

	defer resp.Body.Close()

	//fmt.Println("status", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body: ", string(body))
	log.Infof(`result data: `, string(body))

	return nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	jdt, _ := json.Marshal(obj)
	json.Unmarshal(jdt, &data)
	return data
}
