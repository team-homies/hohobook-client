package app

import (
	"main/app/handler/book"
	"main/common/constant"
)

func Navigate(operation int) (page []string, err error) {
	switch operation {
	// 1000
	case constant.MainPage:
		{
			page = append(page, "1.메인 메뉴입니다.")
			page = append(page, "[1. 도서관리, 2. 대출, 3. 반납]")
			page = append(page, "[e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1100
	case constant.BookManagePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 메뉴입니다.")
			page = append(page, "[1.책 정보 등록, 2.책 조회, 3.책 수정, 4.책 삭제]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1110
	case constant.BookManageRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 메뉴 입니다.")
			page = append(page, "[1.책 정보 등록, 2.책 상세 정보 등록]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1111
	case constant.BookManageInfoRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 -> 1.책 정보 등록 메뉴 입니다.")
			page = append(page, "입력받은 책 정보를 등록하겠습니다.")

		}
	// 1112
	case constant.BookManageDetailRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 -> 2.책 상세 정보 등록 메뉴 입니다.")
			page = append(page, "책 제목으로 해당 책의 상세 정보를 등록하겠습니다.")
			page = append(page, "책 제목을 정확히 입력하고 Enter를 누르세요.")
			book.PostBookDetail()
		}
	// 1120
	case constant.BookManageSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 메뉴 입니다.")
			page = append(page, "[1.책 전체 조회, 2.책 제목 조회, 3.책 상세 조회, 4.대출자 조회]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
		// 1121
	case constant.BookManageAllSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 1.책 전체 조회 메뉴 입니다.")
			page = append(page, "등록된 책들의 정보를 조회하겠습니다.")
			book.GetBookList()
		}
		// 1122
	case constant.BookManageTitleSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 2.책 제목 조회 메뉴 입니다.")
			page = append(page, "책 제목으로 책 정보를 조회하겠습니다.")
			page = append(page, "책 제목 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1123
	case constant.BookManageDetailSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 3.책 상세 정보 조회 메뉴 입니다.")
			page = append(page, "책 제목으로 책 상세 정보를 조회하겠습니다.")
			page = append(page, "책 제목 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1124
	case constant.BookManageRentUserInfoSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 4.대출자 정보 조회 메뉴 입니다.")
			page = append(page, "책 제목으로 대출자 정보를 조회하겠습니다.")
			page = append(page, "책 제목을 정확히 입력하고 Enter를 누르세요.")
		}
	// 1130
	case constant.BookManageUpdatePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 3.책 수정 메뉴 입니다.")
			page = append(page, "책 제목으로 책 정보를 조회하겠습니다.")
			page = append(page, "책 제목 키워드를 입력하고 Enter를 누르세요.")
		}
	// 1140
	case constant.BookManageRemovePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 4.책 삭제 메뉴 입니다.")
			page = append(page, "책 제목으로 책을 삭제하겠습니다.")
			page = append(page, "책 제목을 정확히 입력하고 Enter를 누르세요.")
		}
	// 1200
	case constant.BookRentPage:
		{
			page = append(page, "1.메인 -> 2.대출 메뉴입니다.")
			page = append(page, "책 제목으로 책을 대출하겠습니다.")
			page = append(page, "책 제목을 정확히 입력하고 Enter를 누르세요.")
		}
	// 1300
	case constant.BookReturnPage:
		{
			page = append(page, "1.메인 -> 3.반납 메뉴입니다.")
			page = append(page, "책 제목으로 책을 반납하겠습니다.")
			page = append(page, "책 제목을 정확히 입력하고 Enter를 누르세요.")
		}
		// 9999
		// case constant.LastPage:
		// 	{
		// 		page = append(page, "'b' : 뒤로가서 추가로 작업합니다.")
		// 		page = append(page, "'m' : 첫 화면으로 돌아갑니다.")
		// 		page = append(page, "'e' : 프로그램을 종료합니다.")
		// 	}
	}

	return
}
