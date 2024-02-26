package util

import "main/config"

// GET, POST, PUT, DELETE 방식에 따른 URL 대응 로직 추가
func GenerateURL(path string) (url string) {
	host := config.GetInstance().HOST
	port := config.GetInstance().PORT
	return host + ":" + port + path
}
