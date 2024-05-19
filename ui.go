package main

import (
	"fmt"
)

func printMain() {
	fmt.Printf(`
메인메뉴입니다.
[1. 도서관리, 2. 대출, 3. 반납]
[b.뒤로가기, m.메인메뉴, e.종료]
원하는 키워드를 입력하고 Enter를 누르세요.
`)
}

func printBookManage() {
	fmt.Printf(`
메인 -> 도서관리 메뉴입니다.
[1.책 등록, 2.책 조회, 3.책 수정, 4.책 삭제]
[b.뒤로가기, m.메인메뉴, e.종료]
원하는 키워드를 입력하고 Enter를 누르세요.
`)
}

func printBookSearch() {
	fmt.Printf(`
메인 -> 도서관리 -> 책 조회 메뉴입니다.
[1.책 전체 조회, 2.책 제목 조회, 3.책 상세 조회, 4.대출자 조회]
[b.뒤로가기, m.메인메뉴, e.종료]
원하는 키워드를 입력하고 Enter를 누르세요.
`)
}

// case "3": //수정
// 			var title string
// 			var putChoice string
// 			var column string
// 			var str string
// 			var value string

// 			fmt.Println()
// 			fmt.Print("책 제목 입력> ") //책 제목 입력창
// 			fmt.Scanln(&title)

// 			for {
// 				PutBookMenu() //수정목록 ui
// 				fmt.Print("입력: ")
// 				fmt.Scanln(&putChoice) //수정할 번호 선택
// 				switch putChoice { //컬럼을 반환한다.
// 				case "1": column = "title"; str = "제목"
// 				case "2": column = "author"; str = "저자"
// 				case "3": column = "publisher"; str = "출판사"
// 				case "4": column = "theme"; str = "장르"
// 				case "5": column = "isbn"; str = "ISBN"
// 				case "b", "m", "e":
// 					Common(putChoice)
// 				default:
// 					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
// 					main()
// 				}
// 				fmt.Print(str, " 입력> ") //입력값을 반환한다.
// 				fmt.Scanln(&value)
// 				PutBook(title, column, value) //값 받아와서 JSON형식으로 서버에 넘기는 함수
// 			}
