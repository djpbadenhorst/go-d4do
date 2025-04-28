package passwordPrompt

import (
	"log"
	"strings"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	Password     textinput.Model
}

func Run() string {
	model := initialModel()
	if _, err := tea.NewProgram(&model).Run(); err != nil {
		log.Fatal(err)
	}
	return model.Password.Value()
}

func initialModel() model {
	textInp := textinput.New()

	textInp.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	textInp.CharLimit = 32
	textInp.Width = 20
	textInp.Placeholder = "Password"
	textInp.Focus()
	textInp.EchoMode = textinput.EchoPassword
	textInp.EchoCharacter = '•'
	
	return model{
		Password: textInp,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			s := msg.String()

			if s == "enter" {
				return m, tea.Quit
			}

			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.Password, cmd = m.Password.Update(msg)

	return m, cmd
}


func (m model) View() string {
	var b strings.Builder

	b.WriteString(m.Password.View())
	b.WriteRune('\n')
	
	return b.String()
}
