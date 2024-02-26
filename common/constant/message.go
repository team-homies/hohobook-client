package constant

import "fmt"

// 공통 메세지
const (
	ErrInvalidJsonEncoding = "JSON 인코딩 실패"
	ErrInvalidJsonDecoding = "JSON 디코딩 실패"
	ErrInvalidRequest      = "요청 생성 실패"
	InfoSuccessCall        = "잘못된 키워드입니다. 다시 시도하세요."
	ErrFailCall            = "잘못된 키워드입니다. 다시 시도하세요."
)

func PrintMessage(msg string, sub ...any) {
	fmt.Println(msg, sub)
}
