package data

import (
	"net/url"
)

var (
	queryMap = map[string]string{"a": "Response A", "b": "Response B", "c": "Response C"}
)
func GetBasicData(q url.Values) string {
	retData := "Response Default"
	for k, v := range queryMap {
		if vv := q.Get(k); vv != "" {
			retData = v + ": " + vv
		}
	}
	return retData
}
