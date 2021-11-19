package onebox

import (
	"encoding/json"
	"fmt"
	"testing"

	"git.onespace.co.th/osgolib/oneid"
	"github.com/stretchr/testify/assert"
)

var (
	ONE_BOX_HOST  = "https://uatbox.one.th"
	ONE_BOX_TOKEN = "5d948b748c7be83c197c6db182b26ee0b90bf82c3510da1b8272f4243d62304edef1407aa1d57c303b17299ae4756524c105d88715aefd41b93a6e3d53a4a0c7"
	OneIdCTX      = oneid.Init(
		"https://testoneid.inet.co.th/",
		"430",
		"cbPUgyNQ7Wn68es0RBK8QcAm7g29hoiZurEYN2BX",
		"hEDnyc",
		``,
	)
	username = "devonespace01"
	password = "0neSpace"
)

var onebox = Init(ONE_BOX_HOST, ONE_BOX_TOKEN)

func TestLoginMobile(t *testing.T) {
	var resLogin oneid.ResponseLoginPwdOne
	pwd := oneid.PWD{
		Username: username,
		Password: password,
	}
	if err := OneIdCTX.LoginPWD(pwd, &resLogin).Error(); err != nil {
		fmt.Println(err)
	}
	var res ResponseAccount
	if err := onebox.GetAccount(resLogin.AccessToken, &res); err != nil {
		fmt.Println(err)
	}
	var accountID string
	for _, s := range res.Result {
		if s.AccountName == "ผู้ใช้ทั่วไป" {
			accountID = s.AccountID
		}
	}
	var resCreate ResponseCreateFolder
	if err := onebox.CreateFolder(accountID, "test-folder", &resCreate); err != nil {
		fmt.Print(err)
	}
	filePayload := savePDFBody{
		AccountID:     accountID,
		FolderName:    "test",
		FileExtension: "pdf",
		Base64:        BASE64,
		FolderID:      resCreate.Data.FolderID,
	}
	var resSaveBase64 ResponseSaveFileBase64
	if err := onebox.SavePDFbase64(filePayload, &resSaveBase64); err != nil {
		fmt.Print(err)
	}

	d, err := onebox.DownloadFile(resSaveBase64.Data.ID)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, d, ONE_BOX_HOST+"/onebox_downloads/api/v2/download_file?file_id="+resSaveBase64.Data.ID)
}

func Print(target interface{}) {
	fooByte, _ := json.MarshalIndent(&target, "", "\t")
	fmt.Println(string(fooByte))
}