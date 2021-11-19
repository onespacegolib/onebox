package onebox

type Response struct {
	Result string `json:"result"`
	Data   struct {
		Download string `json:"download"`
	} `json:"data"`
	ErrorMessage interface{} `json:"errorMessage"`
	Code         int         `json:"code"`
}
type ResponseAccount struct {
	Result          []Result `json:"result"`
	Status          string   `json:"status"`
	TransactionCode string   `json:"transaction_code"`
	ErrorCode       string   `json:"errorCode"`
	ErrorMessage    string   `json:"errorMessage"`
}
type Result struct {
	AccountID   string      `json:"account_id"`
	AccountName string      `json:"account_name"`
	BranchNo    interface{} `json:"branch_no"`
	BusinessID  interface{} `json:"business_id"`
	Email       string      `json:"email"`
	Taxid       interface{} `json:"taxid"`
}

type ResponseCreateFolder struct {
	Data struct {
		CreDtm         string `json:"cre_dtm"`
		FolderID       string `json:"folder_id"`
		FolderName     string `json:"folder_name"`
		ParentFolderID string `json:"parent_folder_id"`
	} `json:"data"`
	Message      string `json:"message"`
	Status       string `json:"status"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type savePDFBody struct {
	AccountID     string `json:"account_id"`
	FolderName    string `json:"fileName"`
	FileExtension string `json:"fileExtension"`
	Base64        string `json:"base64"`
	FolderID      string `json:"folder_id"`
}

type ResponseSaveFileBase64 struct {
	Data struct {
		SizeFile   string `json:"size_file"`
		ID         string `json:"id"`
		FolderName string `json:"folder_name"`
		StatusFile string `json:"status_file"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
