package textview

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadJSON(file string) (map[string]string, error) {
	fs, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fs.Close()

	var result map[string]string
	if err := json.NewDecoder(fs).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func FileToTextViewMarkdown(file string) (string, error) {
	result, err := ReadJSON(file)
	if err != nil {
		return "", err
	}

	arr := []string{}
	for k, v := range result {
		arr = append(arr, fmt.Sprintf("* %s: %s", k, v))
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	data, err := json.Marshal(map[string]string{"response": strings.Join(arr, "\n")})
	return string(data), err
}
