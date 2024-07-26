package entities

type Input struct {
	Nums   []int    `json:"nums"`
	Target int      `json:"target"`
	S      string   `json:"s"`
	Words  []string `json:"words"`
}

type TwoSums struct {
	Output []int `json:"output"`
}

type ThreeSums struct {
	Output [][]int `json:"output"`
}

type ExpectedOutput struct {
	Output []int `json:"expected_output"`
}