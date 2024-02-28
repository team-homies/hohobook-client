package constant

// 메인 메뉴
const (
	MainPage       int = 1000
	BookManagePage int = 1100
	BookRentPage   int = 1200
	BookReturnPage int = 1300
)

// 책 관리
const (
	BookManageRegistPage int = 1110
	BookManageSearchPage int = 1120
	BookManageUpdatePage int = 1130
	BookManageRemovePage int = 1140
)

//책 등록
const (
	BookManageInfoRegistPage   int = 1111
	BookManageDetailRegistPage int = 1112
)

// 책 검색
const (
	BookManageAllSearchPage          int = 1121
	BookManageTitleSearchPage        int = 1122
	BookManageDetailSearchPage       int = 1123
	BookManageRentUserInfoSearchPage int = 1124
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
