package typetest

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
}

func New(length int) TypeTest {
	testString := strings.Builder{}
	for i := 0; i < length; i++ {
		err, word := utils.GetWordFromList(rand.Intn(1000))
		if err != nil {
			i--
			continue
		}

		if i == (length - 1) {
			testString.WriteString(word)
		} else {
			testString.WriteString(word + " ")
		}

	}

	return TypeTest{
		TestString: testString.String(),
		Params:     initNewTest(testString.String()),
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		Size:       length,
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
		if !i.IsValid {
			numOfErrs++
		}
	}

	length := len(tt.Params) / 5
	timeDelta := tt.EndTime.Sub(tt.StartTime).Seconds()

	errDeduction := float64(numOfErrs / 60.0)

	// WPM = (number of words / time in minutes)
	wpm := fmt.Sprintf("%.2f", ((float64(length) - errDeduction) / (timeDelta / 60.0)))

	return lipgloss.NewStyle().Foreground(styles.Colors.Orange).Render(wpm)
}

func (tt *TypeTest) GetTestSize() string {
	styled := strings.Builder{}

	styled.WriteString("Test Length: ")

	for _, size := range utils.TestSizes {
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
