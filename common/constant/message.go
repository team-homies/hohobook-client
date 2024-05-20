package constant

import "fmt"

// 공통 메세지
const (
	ErrInvalidJsonEncoding = "JSON 인코딩 실패"
	ErrInvalidJsonDecoding = "JSON 디코딩 실패"
	ErrInvalidRequest      = "요청 생성 실패"
	ClientDoFailCall       = "요청 실행 실패"
	InfoSuccessCall        = "요청 성공."
	ErrFailCall            = "요청 실패."
)

func PrintMessage(msg string, sub ...any) {
	fmt.Println(msg)
}
