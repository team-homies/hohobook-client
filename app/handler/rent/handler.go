package rent

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// 2) 책 조회
// 2-4) 대출자 정보 조회 기능 함수
func GetUser() {
	// 검색어 입력
	fmt.Println("대출자 정보 조회를 선택하셨습니다.")
	fmt.Print("제목 키워드 검색: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// GET요청 -> 조회 요청 생성 : 입력값이 포함된 제목을 가진 책의 상세 정보
	getUrl := "http://localhost:8090/book/search/" + INPUT + "/user"
	res, err := http.Get(getUrl)
	if err != nil {
		fmt.Println("요청 생성 실패:")
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []GetUserResponse
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON 디코딩 실패")
	}

	// 조회 결과 출력
	for _, rent := range data {
		fmt.Printf("| 대출자 명 : %s | 대출자 전화번호 : %s | 대출자 나이 : %s | 대출일시 : %v |\n", rent.UserName, rent.UserPhone, rent.UserAge, rent.RentDate)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		fmt.Println("대출자 정보 조회 성공")
	} else {
		fmt.Println("대출자 정보 조회 실패")
	}

}

// 2. 대출 기능 함수
func RentBook(userData UserData) {
	fmt.Print("대출할 책 제목 입력 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// PUT요청 -> 대출가능 상태 변경 : 가능 -> 불가능
	putUrl := "http://localhost:8090/book/search/" + INPUT + "/rent"
	req, err := http.NewRequest("PUT", putUrl, nil) //요청 생성
	if err != nil {
		fmt.Println("요청 생성 실패")
	}

	// 대출 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("요청 실행 실패")
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		fmt.Println("입력한 책이 존재하지 않습니다.")
		return
	}

	// 대출자 등록
	if res.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("본인 정보를 입력하세요")

		fmt.Print("성명 입력 : ")
		scanner.Scan()
		userData.UserName = scanner.Text()
		fmt.Print("전화번호 입력 : ")
		scanner.Scan()
		userData.UserPhone = scanner.Text()
		fmt.Print("나이 입력 : ")
		scanner.Scan()
		userData.UserAge = scanner.Text()

		// JSON으로 인코딩 -> io.Reader로 변환
		var body bytes.Buffer
		err := json.NewEncoder(&body).Encode(userData)
		if err != nil {
			fmt.Println("JSON 인코딩 실패", err)
		}
		// POST요청 -> 등록 요청 생성 : 대출자 정보
		postUrl := "http://localhost:8090/book/search/" + INPUT + "/rent/user"
		res, err := http.Post(postUrl, "application/json", &body)
		if err != nil {
			fmt.Println("요청 생성 실패")
		}
		defer res.Body.Close()

		// 대출 상태 확인
		if res.StatusCode == http.StatusOK {
			fmt.Println("대출 완료")
		} else {
			fmt.Println("대출 실패", res.StatusCode)
		}
	} else {
		fmt.Println("대출 불가능 상태입니다.")
	}
}

// 3. 반납 기능 함수
func ReturnBook() {
	fmt.Print("반납할 책 제목 입력 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// PUT요청 -> 대출가능 상태 변경 : 불가능 -> 가능
	putUrl := "http://localhost:8090/book/search/" + INPUT + "/return"
	req, err := http.NewRequest("PUT", putUrl, nil) //요청 생성
	if err != nil {
		fmt.Println("반납 실패")
	}

	// 반납 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("반납 실패")
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		fmt.Println("반납 불가능 상태입니다")
		return
	}

	// 대출자 정보 삭제
	if res.StatusCode == http.StatusOK {
		// DELETE요청 -> 삭제 요청 생성 : 반납자 정보
		deleteUrl := "http://localhost:8090/book/search/" + INPUT + "/return/user"
		req, err := http.NewRequest("DELETE", deleteUrl, nil) //요청 생성
		if err != nil {
			fmt.Println("요청 생성 실패")
		}

		// 반납 요청 실행
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("요청 실행 실패")
		}
		defer res.Body.Close()

		// 대출 상태 확인
		if res.StatusCode == http.StatusOK {
			fmt.Println("반납 완료")
		} else {
			fmt.Println("반납 실패")
		}
	} else {
		fmt.Println("반납 불가능 상태입니다.")
	}
}
