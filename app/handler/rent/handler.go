package rent

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"main/common/constant"
	"main/common/util"
	"main/config"
	"net/http"
	"os"
	"strconv"
)

// 2) 책 조회
// 2-4) 대출자 정보 조회 기능 함수
func FindRentUser() {
	// 검색어 입력
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// GET요청 -> 조회 요청 생성 : 입력값이 포함된 제목을 가진 책의 상세 정보
	// getUrl := "http://localhost:8090/book/search/" + INPUT + "/user"
	getUrl := util.GenerateURL(config.InputInstance(INPUT).PATH.GET.FindRentUser)
	res, err := http.Get(getUrl)
	if err != nil {
		constant.PrintMessage(constant.ErrInvalidRequest)
	}
	defer res.Body.Close()

	// JSON 디코딩
	var data []GetUserResponse
	json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		constant.PrintMessage(constant.ErrInvalidJsonDecoding)
	}

	// 조회 결과 출력
	for _, rent := range data {
		fmt.Printf("| 대출자 명 : %s | 대출자 전화번호 : %s | 대출자 나이 : %d | 대출일시 : %v |\n", rent.UserName, rent.UserPhone, rent.UserAge, rent.RentDate)
	}

	// 조회 여부 확인
	if res.StatusCode == http.StatusOK {
		constant.PrintMessage(constant.InfoSuccessCall)
	} else {
		constant.PrintMessage(constant.ErrFailCall)
	}

}

// 2. 대출 기능 함수
func RentBook(userData UserData) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// PUT요청 -> 대출가능 상태 변경 : 가능 -> 불가능
	// putUrl := "http://localhost:8090/book/rent/" + INPUT
	putUrl := util.GenerateURL(config.InputInstance(INPUT).PATH.PUT.UpdateIsused)
	req, err := http.NewRequest("PUT", putUrl, nil) //요청 생성
	if err != nil {
		constant.PrintMessage(constant.ErrInvalidRequest)
	}

	// 대출 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		constant.PrintMessage(constant.ErrInvalidJsonDecoding)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		constant.PrintMessage(constant.ErrFailCall)
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
		UserAgeStr := scanner.Text()

		age, _ := strconv.Atoi(UserAgeStr)
		userData.UserAge = age

		// JSON으로 인코딩 -> io.Reader로 변환
		var body bytes.Buffer
		err := json.NewEncoder(&body).Encode(userData)
		if err != nil {
			constant.PrintMessage(constant.ErrInvalidJsonEncoding)
		}
		// POST요청 -> 등록 요청 생성 : 대출자 정보
		// postUrl := "http://localhost:8090/book/rent/" + INPUT + "/user"
		postUrl := util.GenerateURL(config.InputInstance(INPUT).PATH.POST.CreateUserInfoForRent)
		res, err := http.Post(postUrl, "application/json", &body)
		if err != nil {
			constant.PrintMessage(constant.ErrInvalidRequest)
		}
		defer res.Body.Close()

		// JSON 디코딩
		var data map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			constant.PrintMessage(constant.ErrInvalidJsonDecoding)
		}

		// 대출 상태 확인
		if res.StatusCode == http.StatusOK {
			constant.PrintMessage(constant.InfoSuccessCall)
		} else {
			constant.PrintMessage(constant.ErrFailCall)
		}
	} else {
		constant.PrintMessage(constant.ErrFailCall)
	}
}

// 3. 반납 기능 함수
func ReturnBook() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	INPUT := scanner.Text()

	// PUT요청 -> 대출가능 상태 변경 : 불가능 -> 가능
	// putUrl := "http://localhost:8090/book/return/" + INPUT
	putUrl := util.GenerateURL(config.InputInstance(INPUT).PATH.PUT.ReturnBook)
	req, err := http.NewRequest("PUT", putUrl, nil) //요청 생성
	if err != nil {
		constant.PrintMessage(constant.ErrInvalidRequest)
	}

	// 반납 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		constant.PrintMessage(constant.ClientDoFailCall)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		constant.PrintMessage(constant.ErrFailCall)
		return
	}

	// 대출자 정보 삭제
	if res.StatusCode == http.StatusOK {
		// DELETE요청 -> 삭제 요청 생성 : 반납자 정보
		// deleteUrl := "http://localhost:8090/book/return/" + INPUT + "/user"
		deleteUrl := util.GenerateURL(config.InputInstance(INPUT).PATH.DELETE.RemoveUserInfoForReturn)
		req, err := http.NewRequest("DELETE", deleteUrl, nil) //요청 생성
		if err != nil {
			constant.PrintMessage(constant.ErrInvalidRequest)
		}

		// 반납 요청 실행
		res, err := client.Do(req)
		if err != nil {
			constant.PrintMessage(constant.ClientDoFailCall)
		}
		defer res.Body.Close()

		// 대출 상태 확인
		if res.StatusCode == http.StatusOK {
			constant.PrintMessage(constant.InfoSuccessCall)
		} else {
			constant.PrintMessage(constant.ErrFailCall)
		}
	} else {
		constant.PrintMessage(constant.ErrFailCall)
	}
}
