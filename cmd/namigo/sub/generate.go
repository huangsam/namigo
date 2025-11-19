package sub

import (
	"context"
	"fmt"
	"time"

	"github.com/huangsam/namigo/v2/internal/core"
	"github.com/huangsam/namigo/v2/pkg/generate"
	"github.com/urfave/cli/v3"
)

const outputDelay = 500 * time.Millisecond

// GeneratePromptAction generates a prompt for AI chatbots.
func GeneratePromptAction(_ context.Context, cmd *cli.Command) error {
	purpose, err := core.GetString(cmd, "purpose", "👋 Enter project purpose")
	if err != nil {
		return err
	}
	theme, err := core.GetString(cmd, "theme", "👋 Enter project theme")
	if err != nil {
		return err
	}
	demographics, err := core.GetString(cmd, "demographics", "👋 Enter target demographics")
	if err != nil {
		return err
	}
	interests, err := core.GetString(cmd, "interests", "👋 Enter target interests")
	if err != nil {
		return err
	}
	maxSize := cmd.Int("size")
	maxLength := cmd.Int("length")

	prompt, err := generate.Prompt(purpose, theme, demographics, interests, maxSize, maxLength)
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
