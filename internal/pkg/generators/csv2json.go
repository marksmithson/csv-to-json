package generators

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func CSVToJSON(csvReader io.Reader) (jsonBytes []byte, error error) {
	if csvReader == nil {
		return emptyJson(), errors.New("csvReader is nil")
	}
	var header []string
	jsonData := make([]map[string]string, 0)
	reader := csv.NewReader(csvReader)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil && !isFieldCountError(error) {
			return emptyJson(), error
		}

		// capture the header
		if header == nil {
			header = line
			continue
		}

		var data = make(map[string]string)

		for i, heading := range header {
			if len(line) > i {
				data[heading] = line[i]
			}
		}
		if len(line) > len(header) {
			for i, val := range line[len(header):] {
				data[fmt.Sprintf("_%v_", i+1+len(header))] = val
			}
		}
		jsonData = append(jsonData, data)
	}
	jsonBytes, _ = json.Marshal(jsonData)
	return jsonBytes, nil
}

func isFieldCountError(error error) bool {
	if parseError, ok := error.(*csv.ParseError); ok {
		return parseError.Err == csv.ErrFieldCount
	}
	return false
}

func emptyJson() []byte {
	return []byte("[]")
}
