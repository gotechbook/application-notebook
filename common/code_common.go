package common

const (
	OK = 200 // Success
)

// 根据code 获取信息
func GetCodeMessage(code uint32, msg string) (rst string) {
	codeMap := map[uint32]string{
		OK: "Success",
	}
	if msg == "" {
		if value, ok := codeMap[code]; ok {
			rst = value
		} else {
			rst = "未定义错误类型!"
		}
	}
	return
}
