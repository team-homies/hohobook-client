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
	BookManageRegistPage       int = 1110
	BookManageSearchPage       int = 1120
	BookManageSearchUpdatePage int = 1130
	BookManageUpdatePage1      int = 1131
	BookManageRemovePage       int = 1140
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

// Case
const (
	CaseExit    int = 9991
	CaseDefault int = 9992
	CaseReplay  int = 9993
)

// switch 목록

var OperationMain string                            //main
var OprationManageBook string                       //1100
var OprationBookManageRegistPage string             //1110
var OprationBookManageInfoRegistPage string         //1111
var OprationBookManageDetailRegistPage string       //1112
var OprationBookManageSearchPage string             //1120
var OprationBookManageAllSearchPage string          //1121
var OprationBookManageTitleSearchPage string        //1122
var OprationBookManageDetailSearchPage string       //1123
var OprationBookManageRentUserInfoSearchPage string //1124
var OprationBookManageUpdatePage string             //1130
var OprationBookManageRemovePage string             //1140
var OprationBookRentPage string                     //1200
var OprationBookReturnPage string                   //1300
var OperationSearchBook string                      //1400
var OperationLast string

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

	OperationReplay string = "1"
)
