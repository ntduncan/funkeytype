package utils

type TestModesType = []int
type TestModesStyled = []string

var WordTestSizes = TestModesType{10, 25, 50, 100}
var TimeTestSizes = TestModesType{15, 30, 60, 120}

type TestMode int

const (
	WordsTest TestMode = iota
	TimeTest
)
