package xparams

import "testing"

func TestExtractor_Extract(t *testing.T) {
	query1 := "a=1"
	query2 := "a="
	query3 := "a=1&b=2"
	query4 := "a=1&b="
	query5 := "a=1&b=2&c=3"
	query6 := "a=1&b=2&c=&d="
	query7 := "a={\"a\":\"1\"}"
	query8 := "a={\"a\":\"1\",\"b\":2,\"c\":\"\"}"

	params := DefaultQueryFormExtractor.Extract(query1, LocationQuery)
	if len(params) != 1 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=test" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query2, LocationQuery)
	if len(params) != 1 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=test" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query3, LocationQuery)
	if len(params) != 2 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=1&b=test" && query != "a=test&b=2" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query4, LocationQuery)
	if len(params) != 2 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=1&b=test" && query != "a=test&b=" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query5, LocationQuery)
	if len(params) != 3 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=1&b=2&c=test" && query != "a=1&b=test&c=3" && query != "a=test&b=2&c=3" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query6, LocationQuery)
	if len(params) != 4 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=1&b=2&c=&d=test" && query != "a=1&b=2&c=test&d=" &&
			query != "a=1&b=test&c=&d=" && query != "a=test&b=2&c=&d=" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query7, LocationQuery)
	if len(params) != 1 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=test" {
			t.Error("failed")
		}
	}

	params = DefaultQueryFormExtractor.Extract(query8, LocationQuery)
	if len(params) != 1 {
		t.Error("failed")
	}
	for _, param := range params {
		query := param.Replace("test")
		if query != "a=test" {
			t.Error("failed")
		}
	}
}
