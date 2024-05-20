package util

import "main/config"

// GET, POST, PUT, DELETE 방식에 따른 URL 대응 로직 추가
func GenerateURL(path string) (url string) {
	host := config.Instance().HOST
	port := config.Instance().PORT
	return host + ":" + port + path
}
