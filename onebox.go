package onebox

import (
	"bytes"
	"encoding/json"
	"fmt"

	requests "git.onespace.co.th/osgolib/http-requests"
	"github.com/labstack/echo/v4"
)

var (
	resRequest requests.Response
)

type (
	Context interface {
		GetAccount(oneIDToken string, res *ResponseAccount) error
		CreateFolder(accountID string, folderName string, res *ResponseCreateFolder) error
		SavePDFbase64(data savePDFBody, res *ResponseSaveFileBase64) error
		DownloadFile(string) (string, error)
	}

	context struct {
		host   string
		err    error
		bearer string
	}
)

func (c *context) apiEndpoint(endpoint string) string {
	return c.host + endpoint
}

func (c *context) GetAccount(oneIDToken string, res *ResponseAccount) error {
	var resRequest requests.Response
	headers := map[string]string{
		echo.HeaderContentType:   "application/json",
		echo.HeaderAuthorization: "Bearer " + c.bearer,
	}
	payload := map[string]string{
		`accesstoken`: oneIDToken,
	}
	bP, _ := json.Marshal(payload)
	if err := requests.Call().Post(requests.Params{
		URL:     c.apiEndpoint(APIEndpointGetAccount),
		HEADERS: headers,
		TIMEOUT: 30,
		BODY:    bytes.NewBuffer(bP),
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c.err
	}

	if err := json.Unmarshal(resRequest.Result, &res); err != nil {
		c.err = err
		return c.err
	}

	return nil
}

func (c *context) CreateFolder(accountID string, folderName string, res *ResponseCreateFolder) error {
	var resRequest requests.Response
	headers := map[string]string{
		echo.HeaderContentType:   "application/json",
		echo.HeaderAuthorization: "Bearer " + c.bearer,
	}
	payload := map[string]string{
		"account_id":       accountID,
		"parent_folder_id": "",
		"folder_name":      folderName,
	}
	bP, _ := json.Marshal(payload)
	if err := requests.Call().Post(requests.Params{
		URL:     c.apiEndpoint(APIEndpointCreateFolder),
		HEADERS: headers,
		TIMEOUT: 30,
		BODY:    bytes.NewBuffer(bP),
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c.err
	}

	if err := json.Unmarshal(resRequest.Result, &res); err != nil {
		c.err = err
		return c.err
	}

	return nil
}

func (c *context) SavePDFbase64(data savePDFBody, res *ResponseSaveFileBase64) error {
	var resRequest requests.Response
	headers := map[string]string{
		echo.HeaderContentType:   "application/json",
		echo.HeaderAuthorization: "Bearer " + c.bearer,
	}
	bP, _ := json.Marshal(data)
	if err := requests.Call().Post(requests.Params{
		URL:     c.apiEndpoint(APIEndpointSaveBase64),
		HEADERS: headers,
		TIMEOUT: 30,
		BODY:    bytes.NewBuffer(bP),
	}, &resRequest).Error(); err != nil {
		c.err = err
		return c.err
	}

	if err := json.Unmarshal(resRequest.Result, &res); err != nil {
		c.err = err
		return c.err
	}

	return nil
}

func (c *context) DownloadFile(fileID string) (string, error) {
	if fileID == "" {
		return "", fmt.Errorf("file not found")
	}
	return c.apiEndpoint(APIEndpointDownload) + "?file_id=" + fileID, nil
}
