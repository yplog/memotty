package models

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type QuizMode int

const (
	ModeMultipleChoice QuizMode = iota
	ModeWrittenAnswer
)

type QuizState int

const (
	StateFileSelection QuizState = iota
	StateMenu
	StateQuestion
	StateResult
)

type Question struct {
	Question    string
	Options     []string
	Correct     int
	CorrectText string
	Mode        QuizMode
}

type Model struct {
	Questions    []Question
	CurrentQ     int
	Cursor       int
	UserAnswers  []int
	UserTexts    []string
	CorrectCount int
	State        QuizState
	SelectedMode QuizMode
	InputText    string
	CSVFiles     []string
	SelectedFile string
}

func NewModel() Model {
	return Model{
		State:       StateFileSelection,
		Cursor:      0,
		UserAnswers: make([]int, 5),
		UserTexts:   make([]string, 5),
		CSVFiles:    []string{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// This is a placeholder - actual implementation will be in cmd/quiz/main.go
func (m Model) View() string {
	return ""
}

func (m Model) HandleFileSelectionUpdate(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "down", "j":
		if m.Cursor < len(m.CSVFiles)-1 {
			m.Cursor++
		}
	case "enter", " ":
		if len(m.CSVFiles) > 0 {
			m.SelectedFile = m.CSVFiles[m.Cursor]
			m.State = StateMenu
			m.Cursor = 0
		}
	}
	return m, nil
}

func (m Model) HandleMenuUpdate(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "down", "j":
		if m.Cursor < 1 {
			m.Cursor++
		}
	case "enter", " ":
		m.SelectedMode = QuizMode(m.Cursor)
		m.CurrentQ = 0
		m.Cursor = 0
		m.InputText = ""
		m.CorrectCount = 0
		m.State = StateQuestion
	case "b":
		// Back to file selection
		m.State = StateFileSelection
		m.Cursor = 0
	}
	return m, nil
}

func (m Model) HandleQuestionUpdate(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up":
		if m.SelectedMode == ModeMultipleChoice && m.Cursor > 0 {
			m.Cursor--
		}
	case "down":
		if m.SelectedMode == ModeMultipleChoice && m.Cursor < len(m.Questions[m.CurrentQ].Options)-1 {
			m.Cursor++
		}
	case "k":
		if m.SelectedMode == ModeMultipleChoice && m.Cursor > 0 {
			m.Cursor--
		} else if m.SelectedMode == ModeWrittenAnswer {
			m.InputText += "k"
		}
	case "j":
		if m.SelectedMode == ModeMultipleChoice && m.Cursor < len(m.Questions[m.CurrentQ].Options)-1 {
			m.Cursor++
		} else if m.SelectedMode == ModeWrittenAnswer {
			m.InputText += "j"
		}
	case "enter":
		if m.SelectedMode == ModeMultipleChoice {
			m.UserAnswers[m.CurrentQ] = m.Cursor
			if m.Cursor == m.Questions[m.CurrentQ].Correct {
				m.CorrectCount++
			}
		} else {
			m.UserTexts[m.CurrentQ] = strings.TrimSpace(m.InputText)
			if strings.ToLower(m.UserTexts[m.CurrentQ]) == strings.ToLower(m.Questions[m.CurrentQ].CorrectText) {
				m.CorrectCount++
			}
		}

		m.CurrentQ++
		if m.CurrentQ >= len(m.Questions) {
			m.State = StateResult
		} else {
			m.Cursor = 0
			m.InputText = ""
		}
	case "backspace":
		if m.SelectedMode == ModeWrittenAnswer && len(m.InputText) > 0 {
			m.InputText = m.InputText[:len(m.InputText)-1]
		}
	default:
		if m.SelectedMode == ModeWrittenAnswer && len(msg.String()) == 1 {
			m.InputText += msg.String()
		}
	}
	return m, nil
}

func (m Model) HandleResultUpdate(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "enter", " ":
		return m, tea.Quit
	case "r":
		return NewModel(), nil
	case "m":
		m.State = StateMenu
		m.Cursor = 0
		m.CurrentQ = 0
		m.CorrectCount = 0
		m.InputText = ""
	case "f":
		m.State = StateFileSelection
		m.Cursor = 0
		m.CurrentQ = 0
		m.CorrectCount = 0
		m.InputText = ""
	}
	return m, nil
}
