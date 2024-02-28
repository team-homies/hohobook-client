package constant

// 메인 메뉴
const (
	MainPage       string = "1000"
	BookManagePage string = "1100"
	BookRentPage   string = "1200"
	BookReturnPage string = "1300"
)

// 책 관리
const (
	BookManageRegistPage string = "1110"
	BookManageSearchPage string = "1120"
	BookManageUpdatePage string = "1130"
	BookManageRemovePage string = "1140"
)

//책 등록
const (
	BookManageInfoRegistPage   string = "1111"
	BookManageDetailRegistPage string = "1112"
)

// 책 검색
const (
	BookManageAllSearchPage          string = "1121"
	BookManageTitleSearchPage        string = "1122"
	BookManageDetailSearchPage       string = "1123"
	BookManageRentUserInfoSearchPage string = "1124"
)

// 메뉴 마지막
const (
	LastPage = 9999
)

// // 대출
// const (
// 	BookRentRegistPage = 1210
// )

// // 반납
// const (
// 	BookReturnRegistPage = 1310
// )

// 사용자 오퍼레이션
const (
	OperationBook   string = "1"
	OperationRent   string = "2"
	OperationReturn string = "3"

	OperationBack string = "b"
	OperationMenu string = "m"
	OperationExit string = "e"

	OperationRegist string = "1"
	OperationSearch string = "2"
	OperationUpdate string = "3"
	OperationRemove string = "4"

	OperationInfoRegist   string = "1"
	OperationDetailRegist string = "2"

	OperationAllSearch      string = "1"
	OperationTitleSearch    string = "2"
	OperationDetailSearch   string = "3"
	OperationUserInfoSearch string = "4"
)
