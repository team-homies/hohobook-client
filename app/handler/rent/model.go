package rent

import "time"

// 유저가 대출할 때 등록해야할 정보
type UserData struct {
	UserName  string `json:"user_name"`
	UserPhone string `json:"user_phone"`
	UserAge   string `json:"user_age"`
}

// 대여자 정보 데이터 응답 구조체
type GetUserResponse struct {
	ID         uint64    `json:"id"`
	BookId     uint64    `json:"book_id"`
	RentDate   time.Time `json:"rent_date"`
	ReturnDate time.Time `json:"return_date"`
	UserName   string    `json:"user_name"`
	UserPhone  string    `json:"user_phone"`
	UserAge    string    `json:"user_age"`
}

type Book struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Theme       string `json:"theme"`
	Publisher   string `json:"publisher"`
	Rent_status string `json:"rent_status"`
	ISBN        string `json:"isbn"`
}
