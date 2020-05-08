package token

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func Create() error {

	tk, err := Sign(&gin.Context{}, Context{Keys:"game_ops"}, "")
	if err != nil {
		return err
	}

	url := fmt.Sprintf(`%v/api/auth`, viper.GetString(`center`))
	data := fmt.Sprintf(`{"ip":"%v", "data": "%v"}`, viper.GetString("public_ip"), tk)
	jsonStr := []byte(data)

	buffer := bytes.NewBuffer(jsonStr)
	request, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")  //添加请求头
	request.Header.Set("ServerAuth", "game_ops")  //添加请求头

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBytes))

	return nil
}
