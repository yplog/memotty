package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderResults(m models.Model) string {
	content := buildFullResults(m)

	lines := strings.Split(content, "\n")

	viewportSize := 20
	start := m.ScrollOffset
	end := start + viewportSize

	if start >= len(lines) {
		start = len(lines) - viewportSize
	}
	if start < 0 {
		start = 0
	}
	if end > len(lines) {
		end = len(lines)
	}

	var visibleLines []string
	for i := start; i < end && i < len(lines); i++ {
		visibleLines = append(visibleLines, lines[i])
	}

	result := strings.Join(visibleLines, "\n")
	result += "\n\n"
	result += lipgloss.NewStyle().Faint(true).Render("↑/↓: Scroll • m: Main menu • r: Restart • f: Back to files • q/Enter: Exit")

	return result
}

func buildFullResults(m models.Model) string {
	var s strings.Builder

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(2)

	successColor := lipgloss.Color("82")
	errorColor := lipgloss.Color("196")
	warningColor := lipgloss.Color("214")

	percentage := float64(m.CorrectCount) / float64(len(m.Questions)) * 100
	var resultColor lipgloss.Color
	var emoji string

	if percentage >= 80 {
		resultColor = successColor
		emoji = "🎉"
	} else if percentage >= 60 {
		resultColor = warningColor
		emoji = "👍"
	} else {
		resultColor = errorColor
		emoji = "📚"
	}

	s.WriteString(titleStyle.Render("📊 QUIZ RESULTS"))
	s.WriteString("\n\n")

	infoStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("240"))

	if m.SelectedFile != "" {
		s.WriteString(infoStyle.Render(fmt.Sprintf("📄 File: %s", m.SelectedFile)))
		s.WriteString("\n")
	}

	mode := "Multiple Choice"
	if m.SelectedMode == models.ModeWrittenAnswer {
		mode = "Written Answer"
	}
	s.WriteString(infoStyle.Render(fmt.Sprintf("🎯 Mode: %s", mode)))
	s.WriteString("\n\n")

	resultStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(resultColor).
		MarginBottom(2)

	s.WriteString(resultStyle.Render(fmt.Sprintf("%s Your Score: %d/%d (%.1f%%)",
		emoji, m.CorrectCount, len(m.Questions), percentage)))
	s.WriteString("\n\n")

	s.WriteString("📝 Detailed Results:\n\n")

	for i, question := range m.Questions {
		questionNum := fmt.Sprintf("Q%d:", i+1)
		s.WriteString(lipgloss.NewStyle().Bold(true).Render(questionNum))
		s.WriteString(" " + question.Question + "\n")

		if m.SelectedMode == models.ModeMultipleChoice {
			userAnswer := m.UserAnswers[i]
			correctAnswer := question.Correct

			if userAnswer == correctAnswer {
				s.WriteString(lipgloss.NewStyle().Foreground(successColor).Render("✅ Correct"))
				s.WriteString(fmt.Sprintf(" - %s) %s\n",
					string(rune('A'+userAnswer)), question.Options[userAnswer]))
			} else {
				s.WriteString(lipgloss.NewStyle().Foreground(errorColor).Render("❌ Incorrect"))
				s.WriteString(fmt.Sprintf(" - Your answer: %s) %s\n",
					string(rune('A'+userAnswer)), question.Options[userAnswer]))
				s.WriteString(lipgloss.NewStyle().Foreground(successColor).Render("   Correct answer"))
				s.WriteString(fmt.Sprintf(": %s) %s\n",
					string(rune('A'+correctAnswer)), question.Options[correctAnswer]))
			}
		} else {
			userAnswer := strings.TrimSpace(m.UserTexts[i])
			correctAnswer := question.CorrectText

			if strings.ToLower(userAnswer) == strings.ToLower(correctAnswer) {
				s.WriteString(lipgloss.NewStyle().Foreground(successColor).Render("✅ Correct"))
				s.WriteString(fmt.Sprintf(" - %s\n", userAnswer))
			} else {
				s.WriteString(lipgloss.NewStyle().Foreground(errorColor).Render("❌ Incorrect"))
				s.WriteString(fmt.Sprintf(" - Your answer: %s\n", userAnswer))
				s.WriteString(lipgloss.NewStyle().Foreground(successColor).Render("   Correct answer"))
				s.WriteString(fmt.Sprintf(": %s\n", correctAnswer))
			}
		}
		s.WriteString("\n")
	}

	s.WriteString(lipgloss.NewStyle().Faint(true).Render("m: Main menu • r: Restart • f: Back to files • q/Enter: Exit"))

	return s.String()
}
