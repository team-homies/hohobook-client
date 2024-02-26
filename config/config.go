package config

import "sync"

type CallInfo struct {
	HOST string
	PORT string
	PATH path
}

type path struct {
	GET  get
	POST post
}

type get struct {
	METHOD     string
	BookSearch string
}

type post struct {
	METHOD     string
	BookRegist string
}

var (
	instance *CallInfo
	once     sync.Once
)

func GetInstance() *CallInfo {
	if instance == nil {
		instance = &CallInfo{
			HOST: "http://localhost",
			PORT: "8090",
			PATH: path{
				GET: get{
					METHOD:     "GET",
					BookSearch: "/book/list",
				},
				POST: post{
					METHOD:     "POST",
					BookRegist: "/book/registration",
				},
			},
		}
	}
	return instance
}
