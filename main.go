package main

import (
	"fmt"
	"main/app/handler/book"
	"main/app/handler/rent"
	"os"
	"time"
)

var operationMain string
var operationManageBook string
var operationSignBook string
var operationSearchBook string
var operationLast string

func main() {

	for {
	operationMain:
		printMain()
		fmt.Print("입력: ")
		fmt.Scanln(&operationMain)

		switch operationMain {
		case "1": //1. 도서관리
		operationManageBook:
			printBookManage()
			fmt.Print("입력: ")
			fmt.Scanln(&operationManageBook)

			switch operationManageBook {
			case "1": //1) 책 등록
			operationSign:
				printBooksign()
				fmt.Print("입력: ")
				fmt.Scanln(&operationSignBook)

				switch operationSignBook {
				case "1": // 1-1) 책 정보 등록
				reregist:
					books, _ := book.GetBookInput()

					book.PostBook(books)

				operationPOST:
					fmt.Println("[1.책 정보 재등록, b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)

					switch operationLast {
					case "1":
						goto reregist
					case "b": // 뒤로가기
						goto operationSign
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationPOST
					}
				case "2":
				reregitdetail:
					book.PostBookDetail()

				operationPOSTdetail:
					fmt.Println("[1.책 상세정보 재등록, b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)

					switch operationLast {
					case "1":
						goto reregitdetail
					case "b": // 뒤로가기
						goto operationSign
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationPOSTdetail

					}
				case "b": // 뒤로가기
					goto operationManageBook
				case "m": // 메인메뉴
					break
				case "e": // 종료
					fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
					time.Sleep(10 * time.Second)
					os.Exit(0)
				default:
					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
					goto operationSign
				}

			case "2": //2) 책 조회
			operationSearchBook:
				printBookSearch()
				fmt.Print("입력: ")
				fmt.Scanln(&operationSearchBook)

				switch operationSearchBook {
				case "1": // 2-1) 책 전체 조회
					book.GetBookList()

				operationGET1:
					fmt.Println("[b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)
					switch operationLast {
					case "b": // 뒤로가기
						goto operationSearchBook
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationGET1
					}

				case "2": // 2-2) 책 제목 조회
				research1:
					book.GetBookByTitle()
				operationGET2:
					fmt.Println("[1.재검색, b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)
					switch operationLast {
					case "1": // 재검색
						goto research1
					case "b": // 뒤로가기
						goto operationSearchBook
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationGET2
					}
				case "3": // 2-3) 책 상세 조회
				research2:
					book.GetBookDetails()
				operationGET3:
					fmt.Println("[1.재검색, b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)
					switch operationLast {
					case "1": // 재검색
						goto research2
					case "b": // 뒤로가기
						goto operationSearchBook
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationGET3
					}
				case "4": // 2-4) 대출자 정보 조회
				searchUserAgain:
					rent.GetUser()
				operationUSER:
					fmt.Println("[1.재검색, b.뒤로가기, m.메인메뉴, e.종료]")
					fmt.Print("입력: ")
					fmt.Scanln(&operationLast)
					switch operationLast {
					case "1": // 재검색
						goto searchUserAgain
					case "b": // 뒤로가기
						goto operationSearchBook
					case "m": // 메인메뉴
						break
					case "e": // 종료
						fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
						time.Sleep(10 * time.Second)
						os.Exit(0)
					default:
						fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
						goto operationUSER
					}
				case "b": // 뒤로가기
					goto operationManageBook
				case "m": // 메인메뉴
					break
				case "e": // 종료
					fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
					time.Sleep(10 * time.Second)
					os.Exit(0)
				default:
					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
					goto operationSearchBook
				}
			case "3": // 3) 책 수정
			updateAgain:
				fmt.Println("책 수정을 선택하셨습니다.")
				book.PutBook()
			operationUPDATE:
				fmt.Println("[1.재삭제, b.뒤로가기, m.메인메뉴, e.종료]")
				fmt.Print("입력: ")
				fmt.Scanln(&operationLast)

				switch operationLast {
				case "1": // 재삭제
					goto updateAgain
				case "b": // 뒤로가기
					goto operationManageBook
				case "m": // 메인메뉴
					break
				case "e": // 종료
					fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
					time.Sleep(10 * time.Second)
					os.Exit(0)
				default:
					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
					goto operationUPDATE
				}

			case "4": // 4) 책 삭제
			deleteAgain:
				fmt.Println("책 삭제를 선택하셨습니다.")
				book.DeleteBook()
			operationDELETE:
				fmt.Println("[1.재삭제, b.뒤로가기, m.메인메뉴, e.종료]")
				fmt.Print("입력: ")
				fmt.Scanln(&operationLast)

				switch operationLast {
				case "1": // 재삭제
					goto deleteAgain
				case "b": // 뒤로가기
					goto operationManageBook
				case "m": // 메인메뉴
					break
				case "e": // 종료
					fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
					time.Sleep(10 * time.Second)
					os.Exit(0)
				default:
					fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
					goto operationDELETE
				}

			case "b": // 뒤로가기
				goto operationMain
			case "m": // 메인메뉴
				break
			case "e": // 종료
				fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
				time.Sleep(10 * time.Second)
				os.Exit(0)
			default:
				fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
				goto operationManageBook
			}
		case "2": // 2. 대출
		rentAgain:
			fmt.Println("대출 기능을 선택하셨습니다.")
			rent.RentBook(rent.UserData{})
		operationRENT:
			fmt.Println("[1.재대출, b.뒤로가기, m.메인메뉴, e.종료]")
			fmt.Print("입력: ")
			fmt.Scanln(&operationLast)
			switch operationLast {
			case "1": // 재검색
				goto rentAgain
			case "b": // 뒤로가기
				break
			case "m": // 메인메뉴
				break
			case "e": // 종료
				fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
				time.Sleep(10 * time.Second)
				os.Exit(0)
			default:
				fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
				goto operationRENT
			}
		case "3": // 3. 반납
		returnAgain:
			fmt.Println("반납 기능을 선택하셨습니다.")
			rent.ReturnBook()
		operationRETURN:
			fmt.Println("[1.재대출, b.뒤로가기, m.메인메뉴, e.종료]")
			fmt.Print("입력: ")
			fmt.Scanln(&operationLast)
			switch operationLast {
			case "1": // 재대출
				goto returnAgain
			case "b": // 뒤로가기
				break
			case "m": // 메인메뉴
				break
			case "e": // 종료
				fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
				time.Sleep(10 * time.Second)
				os.Exit(0)
			default:
				fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
				goto operationRETURN
			}

		case "b": // 뒤로가기
			goto operationMain
		case "m": // 메인메뉴
			goto operationMain
		case "e": // 종료
			fmt.Println("프로그램을 종료합니다.\n10초 뒤 종료됩니다.")
			time.Sleep(10 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("잘못된 입력입니다. 다시 시도하세요.")
			goto operationMain

		}

	}
}
