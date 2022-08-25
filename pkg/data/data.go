package data

import "net/http"

var (
	queryMap = map[string]string{"a": "Response A", "b": "Response B", "c": "Response C"}
)
func GetBasicData(r *http.Request) string {
	q := r.URL.Query()
	retData := "Response Default"
	for k, v := range queryMap {
		if vv := q.Get(k); vv != "" {
			retData = v + ": " + vv
		}
	}
	return retData
}
