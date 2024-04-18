package prompt

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

type InputPrompt struct {
	Label   string
	Default string
}

type SelectPrompt struct {
	Label string
	Items []string
}

func GetInput(pc InputPrompt) string {
	validate := func(input string) error {
		if len(input) <= 0 && len(pc.Default) <= 0 {
			return errors.New("Input cannot be empty")
		}

		if strings.ContainsAny(input, " ") {
			return errors.New("Input cannot contain spaces")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} › ",
		Valid:   "{{ \"?\" | green }} {{ . | bold }} › ",
		Invalid: "{{ \"?\" | red }} {{ . | bold }} › ",
		Success: "{{ \"✔\" | green}} {{ . | bold }} › ",
	}

	var defaultLabel string = ""
	if len(pc.Default) > 0 {
		defaultLabel = fmt.Sprintf("(default %s)", pc.Default)
	}

	prompt := promptui.Prompt{
		Label:     pc.Label + defaultLabel,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		color.Red("[!] Prompt failed")
		fmt.Println(err)
		os.Exit(1)
	}

	if len(result) <= 0 {
		result = pc.Default
	}

	return result
}

func GetSelect(pc SelectPrompt) string {
	items := pc.Items
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.Label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}
