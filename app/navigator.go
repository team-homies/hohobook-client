package app

import "main/common/constant"

func Navigate(operation int) (page []string, err error) {
	switch operation {
	case constant.MainPage:
		{
			page = append(page, "메인메뉴입니다.")
			page = append(page, "[1.책 재등록, b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	case constant.BookManagePage:
		{
			page = append(page, "메인 -> 도서관리 메뉴입니다.")
			page = append(page, "[1.책 등록, 2.책 조회, 3.책 수정, 4.책 삭제]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}

	case constant.BookManageRegistPage:
		{
			page = append(page, "책 등록 메뉴입니다")
			page = append(page, "[1.책 재등록, b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	case constant.BookManageSearchPage:
		{
			page = append(page, "책 조회 메뉴입니다.")
			page = append(page, "[1.책 전체 조회, 2.책 제목 조회, 3.책 상세 조회, 4.대출자 조회]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	case constant.BookRentPage:
		{
			page = append(page, "메인 -> 대출 메뉴입니다.")
			page = append(page, "[1.대출 등록, 2.대출 조회]")
			page = append(page, "[b.뒤로가기, m.메인메뉴, e.종료]")
			page = append(page, "원하는 키워드를 입력하고 Enter를 누르세요.")
		}
	}

	return
}
