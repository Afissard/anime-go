package app

import (
	"github.com/charmbracelet/bubbles/textinput"
)

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		textInput: ti,
		err:       nil,
	}
}
