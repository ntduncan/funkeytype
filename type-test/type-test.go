package typetest

import (
	"time"
)

type TestElem struct {
	Char    string
	Input   string
	IsValid bool
}

type TestParams []TestElem

type TypeTest struct {
	TestString string
	Params     TestParams
	Wpm        int
	startTime  time.Time
	endTime    time.Time
}

func New(testString string) TypeTest {
	return TypeTest{
		TestString: testString,
		Params:     initNewTest(testString),
		Wpm:        0,
		startTime:  time.Time{},
		endTime:    time.Time{},
	}

}

func initNewTest(testString string) TestParams {

	var params TestParams

	for _, chr := range testString {
		elem := TestElem{
			Char:    string(chr),
			Input:   "",
			IsValid: false,
		}

		params = append(params, elem)
	}

	return params
}

func (tt *TypeTest) HandleKeyPress(key string, index int) {
	/*nilTime := time.Time{}
	if tt.startTime == nilTime {
		tt.startTime = time.Now()
	}*/

	tt.Params[index].Input = key
	if key == tt.Params[index].Char {
		tt.Params[index].IsValid = true
	}

}

/*
func (tt *TypeTest) startTest() {
	tt.startTime = time.Now()

}

func (tt *TypeTest) endTest() {
	tt.endTime = time.Now()

	//@TODO: Calculate WPM

}*/
