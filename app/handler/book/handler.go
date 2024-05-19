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
func PostBook(book []Book) {

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
	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패", err)
	}

	// 등록 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("등록 성공")
	} else {
		fmt.Println("등록 실패", err)
	}

}

// 1-1) 책 등록 시 입력 기능 함수
func GetBookInput() (books []Book, err error) {
	//fmt.Scanln()을 사용하면 스페이스바를 입력하면 다음항목으로 넘어가버림
	//때문에 한 줄 전체를 읽어서 처리하는 방법인 bufio 패키지의 NewScanner 함수를 사용
	fmt.Println("책 등록을 선택하셨습니다.")
	fmt.Println("등록할 책의 정보를 입력하세요:\n①제목/②저자/③테마/④출판사/⑤ISBN")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("책 정보를 입력하려면 아무키를 누르세요\n책 내용을 그만 입력하려면 'q' 키를 눌러주세요")
		scanner.Scan()
		if scanner.Text() == "q" {
			break

		}
		var newBook Book

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

		books = append(books, newBook)

	}
	return books, nil
}

// 1-2) 책 상세정보 등록 기능 함수
func PostBookDetail() (bookDetails []BookDetail, err error) {

	// Post함수의 매개변수는 func http.Post(url string, contentType string, body io.Reader)
	// io.Reader : 데이터를 읽을 수 있는 객체를 표현하는 인터페이스
	// 요청 바디에 데이터를 보내기 위해서는 io.Reader형식으로 보내야함
	fmt.Print("제목 키워드 검색: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	fmt.Println("책 상세정보 등록을 선택하셨습니다.")
	fmt.Println("등록할 책의 정보를 입력하세요:\n①소개글/②목차")

	var newBookDetail BookDetail

	//1. Scan()을 사용하여 한 줄씩 읽어옴
	//2. Text()를 사용하여 사용자 입력을 문자열로 가져옴
	fmt.Print("①소개글: ")
	scanner.Scan()
	newBookDetail.Description = scanner.Text()
	fmt.Print("②목차: ")
	scanner.Scan()
	newBookDetail.Index = scanner.Text()

	bookDetails = append(bookDetails, newBookDetail)

	// bytes.Buffer : book struct -> JSON으로 인코딩 -> io.Reader로 변환역할
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(bookDetails)
	if err != nil {
		fmt.Println("JSON 인코딩 실패")
	}

	// POST 요청 -> 등록 요청 생성 : 책 상세정보
	postUrl := "http://localhost:8090/book/registration/" + INPUT + "/detail"
	res, err := http.Post(postUrl, "application/json", &body) //서버에 요청 보내고 결과값(등록이 완료되었습니다메세지)을 res에 담아 디코딩
	if err != nil {
		fmt.Println("요청 생성 실패")
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패")
	}

	// 등록 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("등록 성공")
	} else {
		fmt.Println("등록 실패")
	}
	return

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
		fmt.Println("전체 조회 실패")
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
		fmt.Println("요청 생성 실패:")
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []Book
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패")
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 제목: %s | 저자: %s | 테마: %s | 출판사: %s | 대출가능여부: %s | isbn: %s |\n", book.Title, book.Author, book.Theme, book.Publisher, book.Rent_status, book.ISBN)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("제목 조회 성공")
	} else {
		fmt.Println("제목 조회 실패")
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
		fmt.Println("요청 생성 실패:")
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []BookDetail
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패")
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 책 소개글 : %s | 목차 : %s |\n", book.Description, book.Index)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("제목 조회 성공")
	} else {
		fmt.Println("제목 조회 실패")
	}

}

// 3) 수정
// 3-1) 수정 전 책 정보 조회 기능 함수
func GetBookChange() {
	fmt.Print("수정할 책 제목 입력 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// GET 요청 -> 수정할 책 리스트 조회
	getUrl := "http://localhost:8090/book/search/" + INPUT + "/change"
	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("요청 생성 실패:")
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []Book
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패")
	}

	// 조회 결과 출력
	for _, book := range data {
		fmt.Printf("| 책번호: %d | 제목: %s | 저자: %s | 장르: %s | 출판사: %s | isbn: %s |\n", book.ID, book.Title, book.Author, book.Theme, book.Publisher, book.ISBN)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("제목 조회 성공")
	} else {
		fmt.Println("제목 조회 실패")
	}
}

// 3-2) 책 정보 수정 기능 함수
func PutBook() (books Book, err error) {
	// var books []Book

	fmt.Print("수정할 책 번호 입력 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	fmt.Println("수정할 책의 정보를 입력하세요:\n①제목/②저자/③테마/④출판사/⑤ISBN")

	// var newBook Book

	//1. Scan()을 사용하여 한 줄씩 읽어옴
	//2. Text()를 사용하여 사용자 입력을 문자열로 가져옴
	fmt.Print("①제목: ")
	scanner.Scan()
	books.Title = scanner.Text()
	fmt.Print("②저자: ")
	scanner.Scan()
	books.Author = scanner.Text()
	fmt.Print("③장르: ")
	scanner.Scan()
	books.Theme = scanner.Text()
	fmt.Print("④출판사: ")
	scanner.Scan()
	books.Publisher = scanner.Text()
	fmt.Print("⑤ISBN: ")
	scanner.Scan()
	books.ISBN = scanner.Text()

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(books)
	if err != nil {
		fmt.Println("JSON 인코딩 실패", err)
	}

	// PUT요청 -> 수정
	putUrl := "http://localhost:8090/book/search/" + INPUT + "/change"
	req, err := http.NewRequest("PUT", putUrl, &body)
	if err != nil {
		fmt.Println("요청 생성 실패", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("요청 실행 실패")
	}
	defer res.Body.Close()

	// 수정 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("수정 성공")
	} else {
		fmt.Println("수정 실패", err)
	}
	return
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
		fmt.Println("요청 생성 실패")
		return
	}

	// 삭제 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("요청 실행 실패")
	}
	defer res.Body.Close()

	// 삭제 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("삭제 성공")
	} else {
		fmt.Println("삭제 실패")
	}
}
