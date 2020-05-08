package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

// 1.初始化路由
func init() {
	// 初始化
	router = gin.Default()
	router.GET("/health", HealthCheck)
	router.GET("/disk", DiskCheck)
	router.GET("/cpu", CPUCheck)
	router.GET("/ram", RAMCheck)
}

func GetTest(url string, router *gin.Engine) []byte {
	// 构造get
	request := httptest.NewRequest(http.MethodGet, url, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, request)

	// 提取响应
	resp := w.Result()
	defer resp.Body.Close()

	// 读取响应
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}

func PostJsonTest(url string, param map[string]interface{}, router *gin.Engine) []byte {
	// 将参数转换为json比特流
	jsonByte, _ := json.Marshal(param)

	// 构造post请求，json数据已请求body的形式传递
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(jsonByte))

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取相应
	resp := w.Result()
	defer resp.Body.Close()

	// 读取响应Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}

func TestHealthCheck(t *testing.T) {
	// 初始化请求地址
	url := "/health"

	// 发起Get请求
	respBody := GetTest(url, router)
	fmt.Printf("response: %v\n", string(respBody))

}

func TestCPUCheck(t *testing.T) {
	// 初始化请求地址
	url := "/cpu"

	// 发起Get请求
	respBody := GetTest(url, router)
	fmt.Printf("response: %v \n", string(respBody))
}

func TestDiskCheck(t *testing.T) {
	// 初始化请求地址
	url := "/disk"

	// 发起Get请求
	respBody := GetTest(url, router)
	fmt.Printf("response: %v \n", string(respBody))

}

func TestRAMCheck(t *testing.T) {
	// 初始化请求地址
	url := "/ram"

	// 发起Get请求
	respBody := GetTest(url, router)
	fmt.Printf("response: %v \n", string(respBody))

}
