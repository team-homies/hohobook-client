package app

import (
	"main/common/constant"
)

func Navigate(operation int) (page []string) {
	switch operation {
	// 1000
	case constant.MainPage:
		{
			page = append(page, "1.메인 메뉴입니다.]\n")
			page = append(page, "[1. 도서관리, 2. 대출, 3. 반납]\n")
			page = append(page, "[e.종료]\n")
			page = append(page, "[원하는 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1100
	case constant.BookManagePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 메뉴입니다.]\n")
			page = append(page, "[1.책 등록, 2.책 조회, 3.책 수정, 4.책 삭제]\n")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]\n")
			page = append(page, "[원하는 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1110
	case constant.BookManageRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 메뉴 입니다.]\n")
			page = append(page, "[1.책 정보 등록, 2.책 상세 정보 등록]\n")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]\n")
			page = append(page, "[원하는 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1111
	case constant.BookManageInfoRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 -> 1.책 정보 등록 메뉴 입니다.]\n")
			page = append(page, "[입력받은 책 정보를 등록하겠습니다.\n")
		}
	// 1112
	case constant.BookManageDetailRegistPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 1.책 등록 -> 2.책 상세 정보 등록 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 해당 책의 상세 정보를 등록하겠습니다.\n")
			page = append(page, "[책 제목을 정확히 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1120
	case constant.BookManageSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 메뉴 입니다.]\n")
			page = append(page, "[1.책 전체 조회, 2.책 제목 조회, 3.책 상세 조회, 4.대출자 조회]\n")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]\n")
			page = append(page, "[원하는 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1121
	case constant.BookManageAllSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 1.책 전체 조회 메뉴 입니다.]\n")
			page = append(page, "[등록된 책들의 정보를 조회하겠습니다.\n")
			page = append(page, "[입력 :")
		}
	// 1122
	case constant.BookManageTitleSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 2.책 제목 조회 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 책 정보를 조회하겠습니다.\n")
			page = append(page, "[책 제목 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1123
	case constant.BookManageDetailSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 3.책 상세 정보 조회 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 책 상세 정보를 조회하겠습니다.\n")
			page = append(page, "[책 제목 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1124
	case constant.BookManageRentUserInfoSearchPage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 2.책 조회 -> 4.대출자 정보 조회 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 대출자 정보를 조회하겠습니다.\n")
			page = append(page, "[책 제목을 정확히 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1130
	case constant.BookManageUpdatePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 3.책 수정 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 책 정보를 조회하겠습니다.\n")
			page = append(page, "[책 제목 키워드를 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1140
	case constant.BookManageRemovePage:
		{
			page = append(page, "1.메인 -> 1.도서관리 -> 4.책 삭제 메뉴 입니다.]\n")
			page = append(page, "책 제목으로 책을 삭제하겠습니다.\n")
			page = append(page, "[책 제목을 정확히 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1200
	case constant.BookRentPage:
		{
			page = append(page, "1.메인 -> 2.대출 메뉴입니다.]\n")
			page = append(page, "책 제목으로 책을 대출하겠습니다.\n")
			page = append(page, "[책 제목을 정확히 입력하고 Enter를 누르세요.\n")
			page = append(page, "[입력 :")
		}
	// 1300
	case constant.BookReturnPage:
		{
			page = append(page, "1.메인 -> 3.반납 메뉴입니다.]\n")
			page = append(page, "책 제목으로 책을 반납하겠습니다.\n")
			page = append(page, "[책 제목을 정확히 입력하고 Enter를 누르세요.]\n")
			page = append(page, "[입력 :")
		}

	// 9991
	case constant.CaseExit:
		{
			page = append(page, "프로그램을 종료합니다. 5초 뒤 종료됩니다.")
		}

	// 9992
	case constant.CaseDefault:
		{
			page = append(page, "잘못된 입력입니다. 다시 시도하세요.")
		}

	// 9993
	case constant.CaseReplay:
		{
			page = append(page, "[1.재실행, b.뒤로가기, m.메인메뉴, e.종료]\n")
			page = append(page, "[입력 :")
		}

	}

	return
}
