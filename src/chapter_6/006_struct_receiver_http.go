/*
第四节，Go语言接收器（http）
Go语言提供的http包里大量使用了类型的方法。Go语言使用http包进行HTTP的请求，
使用http包的NewRequest()方法可以创建HTTP请求，填充请求中的http头(req.Header)，
再调用http.Client的Do方法，将传入的HTTP请求发送出去。
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}

	/*创建一个http请求*/
	req, err := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("key=value"))

	/*如果出错*/
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	/*为头添加消息*/
	req.Header.Add("User-Agent", "myClient")

	/*开始请求*/
	resp, err := client.Do(req)
	/*处理请求的错误*/
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()
}
