package onebox

const (
	// APIEndpointGetAccount ใช้ค้นหาบัชญีผู้ใช้ด้วย User id เพื่อดูว่าผู้ใช้นั้นๆอยู่ใน Business อะไร
	APIEndpointGetAccount = `/onebox_uploads/api/v2/get_account_byuserid`
	// APIEndpointGetAccount ใช้สร้าง folder ใน onebox
	APIEndpointCreateFolder = `/onebox_uploads/api/create_folder`
	// APIEndpointGetAccount ใช้ ีupload base64 ขึ้นไปยัง onebox
	APIEndpointSaveBase64 = `/onebox_uploads/api/save_base64`
	// APIEndpointDownload ใช้ download เอกสาร
	APIEndpointDownload = `/onebox_downloads/api/v2/download_file`
	// APIEndpointDownload ใช้ download เอกสารเป็น base64
	APIEndpointDownloadBase64 = `/onebox_uploads/api/downloads_files_as_base64`
)
