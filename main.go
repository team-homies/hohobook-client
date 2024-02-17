package main

import (
	"fmt"
	"main/app/handler/book"
	"main/app/handler/rent"
	"os"
	"time"
)

func main() {
	var choice string
	var manageBook string
	var searchBook string
	var menu string

Label1:
	for {
		printMain()
		fmt.Print("입력: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1": //1. 도서관리
		Label2:
			printBookManage()
			fmt.Print("입력: ")
			fmt.Scanln(&manageBook)

			switch manageBook {
			case "1": //1) 책 등록
			Label3_1:
				newBook := book.GetBookInput()
				book.PostBook(newBook)
				fmt.Println("[1.책 재등록, b.뒤로가기, m.메인메뉴, e.종료]")

				fmt.Print("입력: ")
				fmt.Scanln(&menu)
				switch menu {
				case "1":
					goto Label3_1
				case "b": // 뒤로가기
					goto Label2
				case "m": // 메인메뉴
					goto Label1
				case "e": // 종료
					fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
					time.Sleep(10 * time.Second)
					os.Exit(0)
				default:
					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
					goto Label3_1
				}

			case "2": //2) 책 조회
			Label3_2:
				printBookSearch()
				fmt.Print("입력: ")
				fmt.Scanln(&searchBook)
				switch searchBook {
				case "1": // 2-1) 책 전체 조회
					book.GetBookList()
					fmt.Println("[b.뒤로가기, m.메인메뉴, e.종료]")

					fmt.Print("입력: ")
					fmt.Scanln(&menu)
					switch menu {
					case "b": // 뒤로가기
						goto Label3_2
					case "m": // 메인메뉴
						goto Label1
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto Label3_2
					}

				case "2": // 2-2) 책 제목 조회
					book.GetBookByTitle()
					fmt.Println("[b.뒤로가기, m.메인메뉴, e.종료]")

					fmt.Print("입력: ")
					fmt.Scanln(&menu)

					switch menu {
					case "b": // 뒤로가기
						goto Label3_2
					case "m": // 메인메뉴
						goto Label1
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto Label3_2
					}
				case "3": // 2-3) 책 상세 조회
					book.GetBookDetails()
					fmt.Println("[b.뒤로가기, m.메인메뉴, e.종료]")

					fmt.Print("입력: ")
					fmt.Scanln(&menu)
					switch menu {
					case "b": // 뒤로가기
						goto Label2
					case "m": // 메인메뉴
						goto Label1
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto Label3_2
					}
					switch menu {
					case "b": // 뒤로가기
						goto Label2
					case "m": // 메인메뉴
						goto Label1
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto Label3_2
					}
				}
			case "3": // 3) 책 수정
				fmt.Println("책 수정을 선택하셨습니다.")
			case "4": // 4) 책 삭제
				fmt.Println("책 삭제를 선택하셨습니다.")
			case "b": // 뒤로가기
				goto Label1
			case "m": // 메인메뉴
				goto Label1
			case "e": // 종료
				fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
				time.Sleep(10 * time.Second)
				os.Exit(0)
			default:
				fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
				goto Label2
			}
		case "2": // 2. 대출
			fmt.Println("대출 기능을 선택하셨습니다.")
			rent.RentBook(rent.UserData{})
		case "3": // 3. 반납
			fmt.Println("반납 기능을 선택하셨습니다.")
		case "b": // 뒤로가기
			goto Label1
		case "m": // 메인메뉴
			goto Label1
		case "e": // 종료
			fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
			time.Sleep(10 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("잘못된 입력입니다. 다시 시도하세요.")

		}

	}
}
