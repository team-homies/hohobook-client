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

func printBooksign() {
	fmt.Printf(`
메인 -> 도서관리 메뉴입니다.
[1.책 정보 등록, 2.책 상세정보 등록]
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
