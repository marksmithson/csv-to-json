package generators_test

import (
	"github.com/marksmithson/csv-to-json/internal/pkg/generators"
	"strings"
	"testing"
)

func TestNilReader(t *testing.T) {
	json, error := generators.CSVToJSON(nil)
	t.Run("returns error", func(t *testing.T) {
		if error == nil {
			t.Errorf("no error when a nil reader")
		}
	})
	t.Run("returns empty JSON", func(t *testing.T) {
		assertEmptyJSON(json, t)
	})
}

func TestEmptyReader(t *testing.T) {
	json, error := generators.CSVToJSON(strings.NewReader(""))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns empty JSON", func(t *testing.T) {
		assertEmptyJSON(json, t)
	})
}

func TestInvalidCSV(t *testing.T) {
	json, error := generators.CSVToJSON(strings.NewReader("\",,"))
	t.Run("returns error", func(t *testing.T) {
		if error == nil {
			t.Errorf("no error when invalid CSV")
		}
	})
	t.Run("returns empty JSON", func(t *testing.T) {
		assertEmptyJSON(json, t)
	})
}

func TestOnlyHeader(t *testing.T) {
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns empty JSON", func(t *testing.T) {
		assertEmptyJSON(json, t)
	})
}

func TestSimpleCSV(t *testing.T) {
	expected := `[{"col1":"cell1-1","col2":"cell1-2"},{"col1":"cell2-1","col2":"cell2-2"}]`
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2\ncell1-1,cell1-2\ncell2-1,cell2-2"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns expected JSON", func(t *testing.T) {
		assertJSON(expected, json, t)
	})
}

func TestEmptyLinesAtEnd(t *testing.T) {
	expected := `[{"col1":"cell1-1","col2":"cell1-2"}]`
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2\ncell1-1,cell1-2\n\n\n"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns JSON ignoring empty lines", func(t *testing.T) {
		assertJSON(expected, json, t)
	})
}

func TestEmptyLinesInMiddle(t *testing.T) {
	expected := `[{"col1":"cell1-1","col2":"cell1-2"},{"col1":"cell2-1","col2":"cell2-2"}]`
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2\ncell1-1,cell1-2\n\n\ncell2-1,cell2-2"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns JSON ignoring empty lines", func(t *testing.T) {
		assertJSON(expected, json, t)
	})
}

func TestSparseCSV(t *testing.T) {
	expected := `[{"col1":"cell1-1"},{"col1":"cell2-1","col2":"cell2-2"}]`
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2\ncell1-1\ncell2-1,cell2-2"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns expected JSON", func(t *testing.T) {
		assertJSON(expected, json, t)
	})
}

func TestAdditionalField(t *testing.T) {
	expected := `[{"_3_":"cell1-3","col1":"cell1-1","col2":"cell1-2"},{"col1":"cell2-1","col2":"cell2-2"}]`
	json, error := generators.CSVToJSON(strings.NewReader("col1,col2\ncell1-1,cell1-2,cell1-3\ncell2-1,cell2-2"))
	t.Run("does not error", func(t *testing.T) {
		assertNoError(error, t)
	})
	t.Run("returns JSON with colIndex as key", func(t *testing.T) {
		assertJSON(expected, json, t)
	})
}

func assertNoError(error error, t *testing.T) {
	if error != nil {
		t.Errorf("unexpected error %v", error)
	}
}

func assertEmptyJSON(json []byte, t *testing.T) {
	if string(json) != "[]" {
		t.Errorf("expected empty JSON to be returned, was %s", json)
	}
}

func assertJSON(expected string, json []byte, t *testing.T) {
	if string(json) != expected {
		t.Errorf("expected %s JSON to be returned, was %s", expected, string(json))
	}
}
