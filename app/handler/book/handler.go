package book

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 1. 도서관리
// 1) 책 등록 기능 함수
func PostBook(book Book) {

	// Post함수의 매개변수는 func http.Post(url string, contentType string, body io.Reader)
	// io.Reader : 데이터를 읽을 수 있는 객체를 표현하는 인터페이스
	// 요청 바디에 데이터를 보내기 위해서는 io.Reader형식으로 보내야함

	// bytes.Buffer : book struct -> JSON으로 인코딩 -> io.Reader로 변환역할
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(book)
	if err != nil {
		fmt.Println("JSON 인코딩 실패", err)
	}

	// POST 요청 -> 등록 요청 생성 : 책 정보
	postUrl := "http://localhost:8090/book/registration"
	res, err := http.Post(postUrl, "application/json", &body)
	if err != nil {
		fmt.Println("요청 생성 실패", err)
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []Book
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패", err)
	}

	// 등록 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("등록 성공")
	} else {
		fmt.Println("등록 실패", res.Status)
	}

}

// 1-1) 책 등록 시 입력 기능 함수
func GetBookInput() (newBook Book) {
	//fmt.Scanln()을 사용하면 스페이스바를 입력하면 다음항목으로 넘어가버림
	//때문에 한 줄 전체를 읽어서 처리하는 방법인 bufio 패키지의 NewScanner 함수를 사용
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("책 등록을 선택하셨습니다.")
	fmt.Println("등록할 책의 정보를 입력하세요:\n①제목/②저자/③테마/④출판사/⑤ISBN")
	//1. Scan()을 사용하여 한 줄씩 읽어옴
	//2. Text()를 사용하여 사용자 입력을 문자열로 가져옴
	fmt.Print("①제목: ")
	scanner.Scan()
	newBook.Title = scanner.Text()
	fmt.Print("②저자: ")
	scanner.Scan()
	newBook.Author = scanner.Text()
	fmt.Print("③테마: ")
	scanner.Scan()
	newBook.Theme = scanner.Text()
	fmt.Print("④출판사: ")
	scanner.Scan()
	newBook.Publisher = scanner.Text()
	fmt.Print("⑤ISBN: ")
	scanner.Scan()
	newBook.ISBN = scanner.Text()
	return newBook
}

// 2) 책 조회
// 2-1) 책 전체 조회 기능 함수
func GetBookList() {
	fmt.Println("책 전체 조회를 선택하셨습니다.")

	// GET요청 -> 조회 요청 생성 : 책 정보 전체
	getUrl := "http://localhost:8090/book/list"
	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("요청 생성 실패", err)
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []Book
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패", err)
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 제목: %s | 저자: %s | 테마: %s | 출판사: %s | 대출가능여부: %s | ISBN: %s |\n", book.Title, book.Author, book.Theme, book.Publisher, book.Rent_status, book.ISBN)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("전체 조회 성공")
	} else {
		fmt.Println("전체 조회 실패", res.Status)
	}

}

// 2-2) 책 제목 검색 기능 함수
func GetBookByTitle() {
	// 검색어 입력
	fmt.Println("책 제목 조회를 선택하셨습니다.")
	fmt.Print("제목 키워드 검색: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// GET요청 -> 조회 요청 생성 : 입력값이 포함된 제목을 가진 책 정보
	getUrl := "http://localhost:8090/book/search/" + INPUT
	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("요청 생성 실패:", err)
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []Book
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패", err)
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 제목: %s | 저자: %s | 테마: %s | 출판사: %s | 대출가능여부: %s | isbn: %s |\n", book.Title, book.Author, book.Theme, book.Publisher, book.Rent_status, book.ISBN)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("제목 조회 성공")
	} else {
		fmt.Println("제목 조회 실패", res.Status)
	}

}

// 2-3) 책 상세 검색 기능 함수
func GetBookDetails() {
	// 검색어 입력
	fmt.Println("책 상세 조회를 선택하셨습니다.")
	fmt.Print("제목 키워드 검색: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// GET요청 -> 조회 요청 생성 : 입력값이 포함된 제목을 가진 책의 상세 정보
	getUrl := "http://localhost:8090/book/search/" + INPUT + "/detail"
	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("요청 생성 실패:", err)
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []BookDetail
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패", err)
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 책 소개글 : %s | 목차 : %s |\n", book.Description, book.Index)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("제목 조회 성공")
	} else {
		fmt.Println("제목 조회 실패", res.Status)
	}

}

// 3) 책 정보 수정 기능 함수
func PutBook() {

}

// 4) 책 정보 삭제 기능 함수
func DeleteBook() {
	fmt.Println("삭제할 책 제목 입력 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// DELETE요청 -> 삭제 요청 생성 : '입력값==제목'인 책 정보
	deleteUrl := "http://localhost:8090/book/search/" + INPUT + "/del"
	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		fmt.Println("요청 생성 실패", err)
		return
	}

	// 삭제 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("요청 실행 실패", err)
	}
	defer res.Body.Close()

	// 삭제 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("삭제 성공")
	} else {
		fmt.Println("삭제 실패", res.Status)
	}
}
