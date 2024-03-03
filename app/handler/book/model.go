package book

// 책 정보
type Book struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Theme       string `json:"theme"`
	Publisher   string `json:"publisher"`
	Rent_status string `json:"rent_status"`
	ISBN        string `json:"isbn"`
}

// 책 상세 정보
type BookDetail struct {
	ID          uint64 `json:"id"`
	Description string `json:"description"`
	Image       []byte `json:"image"`
	Index       string `json:"index"`
}