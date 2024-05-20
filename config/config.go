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
	METHOD                string
	FindBookList          string
	FindBookInfoByTitle   string
	FindBookDetailByTitle string
	FindRentUser          string
	FindBookChange        string
}

type post struct {
	METHOD                string
	CreateBookInfo        string
	CreateBookDetail      string
	CreateUserInfoForRent string
}

type put struct {
	METHOD             string
	UpdateBookInfoById string
	UpdateIsused       string
	ReturnBook         string
}

type delete struct {
	METHOD                  string
	RemoveBook              string
	RemoveUserInfoForReturn string
}

var (
	instance *CallInfo
	once     sync.Once
)

func Instance() *CallInfo {
	once.Do(func() {
		if instance == nil {
			instance = &CallInfo{
				HOST: "http://localhost",
				PORT: "8090",
				PATH: path{
					GET: get{
						METHOD:       "GET",
						FindBookList: "/book/list",
					},
					POST: post{
						METHOD:         "POST",
						CreateBookInfo: "/book/registration",
					},
				},
			}
		}
	})
	return instance
}

func InputInstance(query string) *CallInfo {
	// once.Do(func() {
	// 	if instance == nil {
	instance = &CallInfo{
		HOST: "http://localhost",
		PORT: "8090",
		PATH: path{
			GET: get{
				METHOD:                "GET",
				FindBookInfoByTitle:   "/book/search/" + query,
				FindBookDetailByTitle: "/book/search/" + query + "/detail",
				FindRentUser:          "/book/search/" + query + "/user",
				FindBookChange:        "/book/search/" + query + "/change",
			},
			POST: post{
				METHOD:                "POST",
				CreateBookDetail:      "/book/registration/" + query + "/detail",
				CreateUserInfoForRent: "/book/rent/" + query + "/user",
			},
			PUT: put{
				METHOD:             "PUT",
				UpdateBookInfoById: "/book/search/" + query + "/change",
				UpdateIsused:       "/book/rent/" + query,
				ReturnBook:         "/book/return/" + query,
			},
			DELETE: delete{
				METHOD:                  "DELETE",
				RemoveBook:              "/book/search/" + query + "/del",
				RemoveUserInfoForReturn: "/book/return/" + query + "/user",
			},
		},
	}
	// }
	// })
	return instance
}
