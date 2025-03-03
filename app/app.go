package app

import "github.com/charmbracelet/bubbles/textinput"

type (
	errMsg error
)

type Model struct {
	textInput textinput.Model
	err       error
}
