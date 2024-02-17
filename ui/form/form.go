package UIForm

import (
	"fmt"
	"log"
	"taskmanager/internal/tasks"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func Create() tasks.Task {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	return task
}

type (
	errMsg error
)

const (
	name = iota
	description
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	task          = tasks.Task{}
)

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func nameValidator(value string) error {
	task.Name = value

	return nil

}

func descriptionValidator(value string) error {
	task.Description = value

	return nil
}

func initialModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Nombre"
	inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[description] = textinput.New()
	inputs[description].Placeholder = "Descripción"
	inputs[description].Width = 50
	inputs[description].Prompt = ""
	inputs[description].Validate = descriptionValidator

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		` Añade una tarea:

 %s
 %s

 %s
 %s

 %s
`,
		inputStyle.Width(30).Render("Nombre de la tarea:"),
		m.inputs[name].View(),
		inputStyle.Width(30).Render("Descripción:"),
		m.inputs[description].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
