package domain

type FileData struct {
	Name     string
	FileType string
}

func NewFileData(name string, fileType string) FileData {
	return FileData{Name: name, FileType: fileType}
}
