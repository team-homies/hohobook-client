package main

import (
	"fmt"
	"main/app"
	"main/app/handler/book"
	"main/app/handler/rent"
	"main/common/constant"
	"os"
	"time"
)

// 과제
// 생성된 util, config, constant, navigater 패키지를 완성하여 모든 로직에 적용

func main() {
	for {
	MainLabel:
		// 1000
		fmt.Print(app.Navigate(1000))
		fmt.Scanln(&constant.OperationMain)
		switch constant.OperationMain {
		case constant.OperationBook:
		BookManageLabel:
			// 1100
			fmt.Print(app.Navigate(1100))
			fmt.Scanln(&constant.OprationManageBook)
			switch constant.OprationManageBook {
			case constant.OperationRegist:
			BookManageRegistPageLabel:
				// 1110
				fmt.Print(app.Navigate(1110))
				fmt.Scanln(&constant.OprationBookManageRegistPage)
				switch constant.OprationBookManageRegistPage {
				case constant.OperationInfoRegist:
				BookManageInfoRegistPageLabel:
					// 1111
					fmt.Print(app.Navigate(1111))
					books, _ := book.GetBookInput()
					book.PostBook(books)
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageInfoRegistPage)
					switch constant.OprationBookManageInfoRegistPage {
					case constant.OperationReplay:
						goto BookManageInfoRegistPageLabel
					case constant.OperationBack:
						goto BookManageRegistPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationDetailRegist:
				BookManageDetailRegistPageLabel:
					// 1112
					fmt.Print(app.Navigate(1112))
					book.PostBookDetail()
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageDetailRegistPage)
					switch constant.OprationBookManageDetailRegistPage {
					case constant.OperationReplay:
						goto BookManageDetailRegistPageLabel
					case constant.OperationBack:
						goto BookManageRegistPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationBack:
					goto BookManageLabel
				case constant.OperationMenu:
					goto MainLabel
				case constant.OperationExit:
					fmt.Print(app.Navigate(9991))
					time.Sleep(5 * time.Second)
					os.Exit(0)
				default:
					fmt.Print(app.Navigate(9992))
				}
			case constant.OperationSearch:
			BookManageSearchPageLabel:
				// 1120
				fmt.Print(app.Navigate(1120))
				fmt.Scanln(&constant.OprationBookManageSearchPage)
				switch constant.OprationBookManageSearchPage {
				case constant.OperationAllSearch:
				BookManageAllSearchPageLabel:
					// 1121
					fmt.Print(app.Navigate(1121))
					book.GetBookList()
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageAllSearchPage)
					switch constant.OprationBookManageAllSearchPage {
					case constant.OperationReplay:
						goto BookManageAllSearchPageLabel
					case constant.OperationBack:
						goto BookManageSearchPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationTitleSearch:
				BookManageTitleSearchPageLabel:
					// 1122
					fmt.Print(app.Navigate(1122))
					book.GetBookByTitle()
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageTitleSearchPage)
					switch constant.OprationBookManageTitleSearchPage {
					case constant.OperationReplay:
						goto BookManageTitleSearchPageLabel
					case constant.OperationBack:
						goto BookManageSearchPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationDetailSearch:
				BookManageDetailSearchPageLabel:
					// 1123
					fmt.Print(app.Navigate(1123))
					book.GetBookDetails()
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageDetailSearchPage)
					switch constant.OprationBookManageDetailSearchPage {
					case constant.OperationReplay:
						goto BookManageDetailSearchPageLabel
					case constant.OperationBack:
						goto BookManageSearchPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationUserInfoSearch:
				BookManageRentUserInfoSearchPageLabel:
					// 1124
					fmt.Print(app.Navigate(1124))
					rent.GetUser()
					fmt.Print(app.Navigate(9993))
					fmt.Scanln(&constant.OprationBookManageRentUserInfoSearchPage)
					switch constant.OprationBookManageRentUserInfoSearchPage {
					case constant.OperationReplay:
						goto BookManageRentUserInfoSearchPageLabel
					case constant.OperationBack:
						goto BookManageSearchPageLabel
					case constant.OperationMenu:
						goto MainLabel
					case constant.OperationExit:
						fmt.Print(app.Navigate(9991))
						time.Sleep(5 * time.Second)
						os.Exit(0)
					default:
						fmt.Print(app.Navigate(9992))
					}
				case constant.OperationBack:
					goto BookManageLabel
				case constant.OperationMenu:
					goto MainLabel
				case constant.OperationExit:
					fmt.Print(app.Navigate(9991))
					time.Sleep(5 * time.Second)
					os.Exit(0)
				default:
					fmt.Print(app.Navigate(9992))
				}
			case constant.OperationUpdate:
			BookManageUpdatePageLabel:
				// 1130
				fmt.Print(app.Navigate(1130))
				book.GetBookChange()
				fmt.Print(app.Navigate(9993))
				fmt.Scanln(&constant.OprationBookManageUpdatePage)
				switch constant.OprationBookManageUpdatePage {
				case constant.OperationReplay:
					goto BookManageUpdatePageLabel
				case constant.OperationBack:
					goto BookManageLabel
				case constant.OperationMenu:
					goto MainLabel
				case constant.OperationExit:
					fmt.Print(app.Navigate(9991))
					time.Sleep(5 * time.Second)
					os.Exit(0)
				default:
					fmt.Print(app.Navigate(9992))
				}
			case constant.OperationRemove:
			BookManageRemovePageLabel:
				// 1140
				fmt.Print(app.Navigate(1140))
				book.DeleteBook()
				fmt.Print(app.Navigate(9993))
				fmt.Scanln(&constant.OprationBookManageRemovePage)
				switch constant.OprationBookManageRemovePage {
				case constant.OperationReplay:
					goto BookManageRemovePageLabel
				case constant.OperationBack:
					goto BookManageLabel
				case constant.OperationMenu:
					goto MainLabel
				case constant.OperationExit:
					fmt.Print(app.Navigate(9991))
					time.Sleep(5 * time.Second)
					os.Exit(0)
				default:
					fmt.Print(app.Navigate(9992))
				}
			case constant.OperationBack:
				goto BookManageLabel
			case constant.OperationMenu:
				goto MainLabel
			case constant.OperationExit:
				fmt.Print(app.Navigate(9991))
				time.Sleep(5 * time.Second)
				os.Exit(0)
			default:
				fmt.Print(app.Navigate(9992))
			}
		case constant.OperationRent:
		BookRentPageLabel:
			// 1200
			fmt.Print(app.Navigate(1200))
			rent.RentBook(rent.UserData{})
			fmt.Print(app.Navigate(9993))
			fmt.Scanln(&constant.OprationBookRentPage)
			switch constant.OprationBookRentPage {
			case constant.OperationReplay:
				goto BookRentPageLabel
			case constant.OperationBack:
				goto MainLabel
			case constant.OperationMenu:
				goto MainLabel
			case constant.OperationExit:
				fmt.Print(app.Navigate(9991))
				time.Sleep(5 * time.Second)
				os.Exit(0)
			default:
				fmt.Print(app.Navigate(9992))
			}
		case constant.OperationReturn:
		BookReturnPageLabel:
			// 1300
			fmt.Print(app.Navigate(1300))
			rent.ReturnBook()
			fmt.Print(app.Navigate(9993))
			fmt.Scanln(&constant.OprationBookReturnPage)
			switch constant.OprationBookReturnPage {
			case constant.OperationReplay:
				goto BookReturnPageLabel
			case constant.OperationBack:
				goto MainLabel
			case constant.OperationMenu:
				goto MainLabel
			case constant.OperationExit:
				fmt.Print(app.Navigate(9991))
				time.Sleep(5 * time.Second)
				os.Exit(0)
			default:
				fmt.Print(app.Navigate(9992))
			}
		case constant.OperationExit:
			fmt.Print(app.Navigate(9991))
			time.Sleep(5 * time.Second)
			os.Exit(0)
		default:
			fmt.Print(app.Navigate(9992))

		}
	}

}
