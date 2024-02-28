package config

import (
	"sync"
)

type CallInfo struct {
	HOST string
	PORT string
	PATH path
}

type path struct {
	GET    get
	POST   post
	PUT    put
	DELETE delete
}

type get struct {
	METHOD             string
	BookList           string
	BookTitleSearch    string
	BookDetailSearch   string
	RentUserInfoSearch string
	BookSearch_Update  string
}

type post struct {
	METHOD             string
	BookRegist         string
	BookRegistDetail   string
	RentUserInfoRegist string
}

type put struct {
	METHOD     string
	BookUpdate string
	BookRent   string
	BookReturn string
}

type delete struct {
	METHOD               string
	BookRemove           string
	ReturnUserInfoRemove string
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
					METHOD:   "GET",
					BookList: "/book/list",
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

func INPUTInstance(INPUT string) *CallInfo {
	if instance == nil {
		instance = &CallInfo{
			HOST: "http://localhost",
			PORT: "8090",
			PATH: path{
				GET: get{
					METHOD:             "GET",
					BookTitleSearch:    "/book/search/" + INPUT,
					BookDetailSearch:   "/book/search/" + INPUT + "/detail",
					RentUserInfoSearch: "/book/search/" + INPUT + "/user",
					BookSearch_Update:  "/book/search/" + INPUT + "/change",
				},
				POST: post{
					METHOD:             "POST",
					BookRegistDetail:   "/book/registration/" + INPUT + "/detail",
					RentUserInfoRegist: "/book/rent/" + INPUT + "/user",
				},
				PUT: put{
					METHOD:     "PUT",
					BookUpdate: "/book/search/" + INPUT + "/change",
					BookRent:   "/book/rent/" + INPUT,
					BookReturn: "/book/return/" + INPUT,
				},
				DELETE: delete{
					METHOD:               "DELETE",
					BookRemove:           "/book/search/" + INPUT + "/del",
					ReturnUserInfoRemove: "/book/return/" + INPUT + "/user",
				},
			},
		}
	}
	return instance
}
