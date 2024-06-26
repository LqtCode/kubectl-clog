/**
 * Author: Lqt
 * Date: 2023/12/28
 * Mail: liuqingtian@52tt.com
 */

package plugin

import (
	"strconv"
	"strings"
)

var _ ContentFilter = (*istioLogFilter)(nil)

type istioLogFilter struct{}

func (f *istioLogFilter) FilteringLine(s string, log Logger) (int, error) {
	const istioScanMax = 200

	ss := strings.Split(s[:min(istioScanMax, len(s))], " ")
	if len(ss) <= 4 {
		return 0, nil
	}

	if !strings.HasPrefix(ss[3], "HTTP") {
		return 0, nil
	}

	httpCode := ss[4]
	if _, err := strconv.ParseInt(httpCode, 10, 64); err != nil {
		return 0, err
	}

	idx := strings.Index(s, httpCode)

	levelTag := httpCode
	if strings.HasPrefix(httpCode, "2") || strings.HasPrefix(httpCode, "3") {
		levelTag = log.WrapBgInfo(httpCode)
	} else {
		levelTag = log.WrapBgError(httpCode)
	}

	return log.Println(s[:idx] + levelTag + s[idx+len(httpCode):])
}
