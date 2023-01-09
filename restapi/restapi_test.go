package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// io.WriteString(w, "{\"status\":\"good\"}")
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Status Created"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 200 != resp.StatusCode {
		t.Fatal("Status Code Not OK")
	}
}
