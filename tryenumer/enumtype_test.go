package tryenumer

import (
	"encoding/json"
	"testing"
)

type mystruct struct {
	Test1 EnumType
	Test2 EnumType2
}

func TestGreen(t *testing.T) {
	result := mystruct{}

	if err := json.Unmarshal([]byte(`{"Test1": "Green"}`), &result); err != nil {
		t.Fatal(err)
	}
	if result.Test1 != Green {
		t.Fatalf("Expecting 'Green' got '%s'", result.Test1.String())
	}
}

func TestGreenInt(t *testing.T) {
	result := mystruct{}

	if err := json.Unmarshal([]byte(`{"Test1": 2}`), &result); err == nil { // FAIL if no error
		t.Fatalf("Bypassed enumer and allowed to parse interger")
	}
}

func TestEnumType2Int(t *testing.T) {
	result := mystruct{}

	// without the enumer code generated it parses as int
	if err := json.Unmarshal([]byte(`{"Test2": 2}`), &result); err != nil {
		t.Fatal(err)
	}

	if EnumType(result.Test2) != Green {
		t.Fatalf("Expecting 2 got '%d'", result.Test2)
	}
}
