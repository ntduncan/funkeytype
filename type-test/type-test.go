package typetest

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/lipgloss"
	"ntduncan.com/typer/styles"
	"ntduncan.com/typer/utils"
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
	StartTime  time.Time
	EndTime    time.Time
	Size       int
	Mode       utils.TestMode
	TestTimer  timer.Model
}

func New(size int, mode utils.TestMode) TypeTest {
	testString := strings.Builder{}

	testLen := size

	if mode == utils.TimeTest {
		testLen = 150
	}

	for i := 0; i < testLen; i++ {
		err, word := utils.GetWordFromList(rand.Intn(1000))
		if err != nil {
			i--
			continue
		}

		if i == (testLen - 1) {
			testString.WriteString(word)
		} else {
			testString.WriteString(word + " ")
		}

	}

	timeout := time.Second * time.Duration(size)
	theTimer := timer.NewWithInterval(timeout, time.Second)

	return TypeTest{
		TestString: testString.String(),
		Params:     initNewTest(testString.String()),
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		Size:       size,
		Mode:       mode,
		TestTimer:  theTimer,
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
	if tt.StartTime.IsZero() {
		tt.StartTest()
	}

	tt.Params[index].Input = key
	if key == tt.Params[index].Char {
		tt.Params[index].IsValid = true
		if index == len(tt.Params)-1 {
			tt.EndTest()
		}

	}

}

func (tt *TypeTest) StartTest() {
	tt.StartTime = time.Now()
}

func (tt *TypeTest) EndTest() {
	tt.EndTime = time.Now()
}

func (tt *TypeTest) GetWPM() string {
	if tt.EndTime.IsZero() {
		return ""
	}

	numOfErrs := 0
	for _, i := range tt.Params {
		if tt.Mode == utils.TimeTest && i.Input == "" {
			break
		}

		if !i.IsValid {
			numOfErrs++
		}
	}

	var wpm string
	errDeduction := float64(numOfErrs / 60.0)
	length := 0

	if tt.Mode == utils.TimeTest {
		for _, param := range tt.Params {
			if param.Input != "" {
				length++
			} else {
				break
			}
		}
		length = length / 5

	} else {
		length = len(tt.Params) / 5
	}

	timeDelta := tt.EndTime.Sub(tt.StartTime).Seconds()

	// WPM = (number of words / time in minutes)
	wpm = fmt.Sprintf("%.2f", ((float64(length) - errDeduction) / (timeDelta / 60.0)))
	return lipgloss.NewStyle().Foreground(styles.Colors.Orange).Render(wpm)

}

func (tt *TypeTest) GetTestSize() string {
	styled := strings.Builder{}
	testModeOptions := tt.GetTestModeSizeOptions()

	if tt.Mode == utils.TimeTest {
		styled.WriteString("Test Duration: ")
	} else {
		styled.WriteString("Test Length: ")
	}

	for _, size := range testModeOptions {
		s := strconv.Itoa(size)
		var color lipgloss.Color
		if tt.Size == size {
			color = styles.Colors.Orange
		} else {
			color = styles.Colors.White
		}
		styled.WriteString(lipgloss.
			NewStyle().
			Foreground(color).
			Render(s + "  "))

	}

	return styled.String()
}

func (tt *TypeTest) GetTestModeSizeOptions() utils.TestModesType {
	switch tt.Mode {
	case utils.WordsTest:
		return utils.WordTestSizes
	case utils.TimeTest:
		return utils.TimeTestSizes
	default:
		return utils.WordTestSizes
	}
}
