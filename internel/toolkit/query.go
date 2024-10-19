package toolkit

import (
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func toPinyin(s string) (string, string) {
	pys := pinyin.LazyConvert(s, nil)
	py := strings.Join(pys, "")

	heads := strings.Builder{}
	for _, v := range pys {
		heads.WriteString(string(v[0]))
	}
	return py, heads.String()
}

func PinyinMatchQuery(src string, match string) bool {
	if strings.Contains(src, match) {
		return true
	}

	py, heads := toPinyin(src)
	if strings.Contains(py, match) {
		return true
	}
	if strings.Contains(heads, match) {
		return true
	}
	return false
}
