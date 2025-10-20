package ytbdtc

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestStruct(t *testing.T) {
	// もこうについて検索！！
	var result SearchResponse

	prm := CreateNewSearchParam("もこう", "channle")
	prm.SetKey("AIzaSyDs63GAVAIW5_yGt7vba_6KHh4nXs7HQE4")
	res, _ := http.Get(prm.ToURL())
	text, _ := io.ReadAll(res.Body)
	t.Log(listToString(result.Items))

	_ = json.Unmarshal(text, &result)
	prm.SetNextPageToken(result.NextPageToken)
	res, _ = http.Get(prm.ToURL())
	text, _ = io.ReadAll(res.Body)

	t.Log(listToString(result.Items))
	res.Body.Close()
}
