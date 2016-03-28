package main

import "fmt"
import "gocurl/curl"

func myHeaderFunc(data []byte, userdata interface{}) int {
	fmt.Printf("Recv headers: %v", string(data))
	return len(data)
}

func main() {
	var url string = "http://www.google.com"
	var ret int = 0

	var easy curl.CURL

	ret = easy.EasyInit()
	fmt.Printf("EasyInit return %d\n", ret)
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.CURLOPT_URL, url)
	easy.EasySetopt(curl.CURLOPT_VERBOSE, 1)
	easy.EasySetopt(curl.CURLOPT_HEADERFUNCTION, myHeaderFunc)

	ret = easy.EasyPerform()
	fmt.Printf("EasyPerform return %d\n", ret)
}
