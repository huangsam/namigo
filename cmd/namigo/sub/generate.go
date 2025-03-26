package sub

import (
	"fmt"
	"time"

	"github.com/huangsam/namigo/internal/util"
	"github.com/huangsam/namigo/pkg/generate"
	"github.com/urfave/cli/v2"
)

const outputDelay = 500 * time.Millisecond

// GeneratePromptAction generates a prompt for AI chatbots.
func GeneratePromptAction(c *cli.Context) error {
	purpose, err := util.GetString(c, "purpose", "👋 Enter project purpose")
	if err != nil {
		return err
	}
	theme, err := util.GetString(c, "theme", "👋 Enter project theme")
	if err != nil {
		return err
	}
	demographics, err := util.GetString(c, "demographics", "👋 Enter target demographics")
	if err != nil {
		return err
	}
	interests, err := util.GetString(c, "interests", "👋 Enter target interests")
	if err != nil {
		return err
	}
	maxSize := c.Int("size")
	maxLength := c.Int("length")

	prompt, err := generate.GeneratePrompt(purpose, theme, demographics, interests, maxSize, maxLength)
	if err != nil {
		return err
	}

	fmt.Println("🍺 Prepare prompt")
	fmt.Println()
	time.Sleep(outputDelay)
	fmt.Println(prompt)
	time.Sleep(outputDelay)
	fmt.Println()
	fmt.Println("🎉 Copy into the AI of your choice!")

	return nil
}
