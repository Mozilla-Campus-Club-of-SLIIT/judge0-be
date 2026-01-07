package submission

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func TestSubmission() {
	data := map[string]string{
		"language_id":     "60",
		"source_code":     "cGFja2FnZSBtYWluOwppbXBvcnQgImZtdCI7IApmdW5jIG1haW4oKXsKZm10LlByaW50bG4oIkhlbGxvIikgCn0=",
		"expected_output": "SGVsbG8=",
	}

	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(
		"https://heshanthenura.myaddr.io/api/submissions?base64_encoded=true&wait=true",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
