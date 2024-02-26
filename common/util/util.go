package util

import "main/config"

func GenerateURL(path string) (url string) {
	host := config.GetInstance().HOST
	port := config.GetInstance().PORT
	return host + ":" + port + path
}
