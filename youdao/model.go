package youdao

type TransResult struct {
	Translation []string `json:"translation"`
	Basic       struct {
		Explains []string `json:"explains"`
	}
	Web []struct {
		Value []string `json:"value"`
		Key   string   `json:"key"`
	}
}
