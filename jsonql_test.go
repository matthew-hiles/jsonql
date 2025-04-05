package jsonql

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	jsonString := `
[
  {
    "name": "elgs",
    "gender": "m",
    "skills": [
      "Golang",
      "Java",
      "C"
    ]
  },
  {
    "name": "enny",
    "gender": "f",
    "skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  },
  {
    "name": "sam",
    "gender": "m",
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ]
  }
]
`
	parser, err := NewStringQuery(jsonString)
	if err != nil {
		t.Error(err)
	}

	var pass = []struct {
		in string
	}{
		{"name='elgs'"},
		{"gender='f'"},
		{"skills.[1]='Sleeping'"},
		{"skills contains 'Verification'"},
	}
	var fail = []struct {
		in string
		ex interface{}
	}{}
	for _, v := range pass {
		result, err := parser.Query(v.in)
		if err != nil {
			t.Error(v.in, err)
		}
		fmt.Println(v.in, result)
		//		if v.ex != result {
		//			t.Error("Expected:", v.ex, "actual:", result)
		//		}
	}
	for range fail {

	}
}

func TestPrepare(t *testing.T) {
	jsStr := `[{"name":"dogman","job":"police"},{"name":"chief","job":"police"}]`

	query := "job = 'police' && name != 'dogman'"

	pq, err := Prepare(query)
	if err != nil {
		t.Errorf("failed to prepare query: %s", err)
		return
	}

	var obj []any
	err = json.Unmarshal([]byte(jsStr), &obj)
	if err != nil {
		t.Errorf("failed to unmarshal json: %s", err)
		return
	}

	rez, err := pq.Query(obj)
	if err != nil {
		t.Errorf("failed to run query: %s", err)
		return
	}
	fmt.Printf("%+v\n", rez)
	return
	/*
		if m, ok := rez.(map[string]string); ok {
			if m["name"] != "chief" {
				t.Errorf("expected chief, got %s", m["name"])
			}
		} else {
			t.Errorf("failed to cast result")
		}
	*/
}
