package interviewed

import (
	"encoding/json"
	"strings"
)

// 定义一个日志解析函数，接收一个字符串类型的参数，能够把字符串按如下格式返回。
// 输入:2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone
// 输出:{"device":"iPhone","location":"Beijing","name":"Charles"}
// time O(N) space O(N)
func ParseLog(s string) (string, error) {
	kvs := make(map[string]string, 0)
	need := strings.Split(s, "|")[2]
	userLogin := strings.Split(need, "&")
	for _, loginInfo := range userLogin {
		before, after, found := strings.Cut(loginInfo, "=")
		if !found {
			continue
		}
		kvs[before] = after
	}
	res, err := json.Marshal(kvs)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
