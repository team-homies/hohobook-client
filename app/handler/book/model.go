package book

// 책 정보
type Book struct {
	ID          uint64
	Title       string
	Author      string
	Theme       string
	Publisher   string
	Rent_status string
	ISBN        string
}

// 책 상세 정보
type BookDetail struct {
	ID          uint64
	Description string
	Image       []byte
	Index       string
}
