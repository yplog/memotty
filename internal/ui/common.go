package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderEmptyQuestionsError() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	var s string
	s += titleStyle.Render("ERROR: No questions loaded")
	s += "\n\n"
	s += "No questions were found in the selected CSV file.\n"
	s += "Please check that your CSV file is properly formatted:\n"
	s += "- Each line should have: question,answer\n"
	s += "- Questions and answers should not be empty\n\n"
	s += lipgloss.NewStyle().Faint(true).Render("q: Back to menu")

	return s
}

func HasValidQuestions(m models.Model) bool {
	return len(m.Questions) > 0 && m.CurrentQ < len(m.Questions)
}
