package system

import (
	"fmt"

	"ntduncan.com/typer/utils"
)

type Config struct {
	TopScores TopScoreMap    `json:"TopScores"`
	Mode      utils.TestMode `json:"Mode"`
	Size      int            `json:"Size"`
}

func (c *Config) UpdateTopScore(score float64) error {
	switch c.Mode {
	case utils.TimeTest:
		switch c.Size {
		case 15:
			c.TopScores.TimeScores[0] = score
			break
		case 30:
			c.TopScores.TimeScores[1] = score
			break
		case 60:
			c.TopScores.TimeScores[2] = score
			break
		case 120:
			c.TopScores.TimeScores[3] = score
			break
		default:
			return fmt.Errorf("Invalid Score Index for Time Test Mode. got=%v", c.Size)
		}

	case utils.WordsTest:
		switch c.Size {
		case 10:
			c.TopScores.WordScores[0] = score
			break
		case 25:
			c.TopScores.WordScores[1] = score
			break
		case 50:
			c.TopScores.WordScores[2] = score
			break
		case 100:
			c.TopScores.WordScores[3] = score
			break
		default:
			return fmt.Errorf("Invalid Score Index for Word Test Mode. got=%v", c.Size)
		}
	}

	return nil
}

func (c *Config) GetTopScore() (float64, error) {

	switch c.Mode {
	case utils.TimeTest:
		index := 0
		switch c.Size {
		case 15:
			index = 0
			break
		case 30:
			index = 1
			break
		case 60:
			index = 2
			break
		case 120:
			index = 3
			break
		default:
			return 0, fmt.Errorf("Invalid Score Index for Time Test Mode. got=%v", c.Size)
		}
		return c.TopScores.TimeScores[index], nil

	case utils.WordsTest:
		index := 0
		switch c.Size {
		case 10:
			index = 0
			break
		case 25:
			index = 1
			break
		case 50:
			index = 2
			break
		case 100:
			index = 3
			break
		default:
			return 0, fmt.Errorf("Invalid Score Index for Word Test Mode. got=%v", c.Size)
		}
		return c.TopScores.WordScores[index], nil
	}

	return 0, fmt.Errorf("Unable to get a score from config")
}
