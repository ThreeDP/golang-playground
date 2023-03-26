package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckWeb(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebs(t *testing.T) {
	websites := []string{
        "http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	expected := map[string]bool {
		"http://google.com": true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}
	result := CheckWebs(mockCheckWeb, websites)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, result %v", expected, result)
	}
}

func slowStubCheckWeb(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebs(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebs(slowStubCheckWeb, urls)
	}
}