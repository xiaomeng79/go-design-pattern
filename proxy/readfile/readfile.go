package readfile

type IReadFile interface {
	ReadFile(string) string
}

type ReadFile struct {

}

func (r ReadFile) ReadFile(filename string) string{
	return "文件内容为：保密内容"
}
