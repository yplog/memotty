package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yplog/memotty/internal/csv"
	"github.com/yplog/memotty/internal/models"
	"github.com/yplog/memotty/internal/ui"
	"github.com/yplog/memotty/internal/update"
)

type App struct {
	model models.Model
}

var (
	version   = "dev"
	commit    = "unknown"
	buildTime = "unknown"
)

func (a App) Init() tea.Cmd {
	return a.model.Init()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch a.model.State {
		case models.StateFileSelection:
			updatedModel, cmd := a.model.HandleFileSelectionUpdate(msg)
			a.model = updatedModel
			return a, cmd
		case models.StateMenu:
			updatedModel, cmd := a.model.HandleMenuUpdate(msg)
			a.model = updatedModel

			if a.model.State == models.StateQuestion {
				questions, err := csv.LoadQuestionsFromCSV(a.model.SelectedFile, a.model.SelectedMode)
				if err != nil {
					log.Printf("Error loading questions: %v", err)
					a.model.Questions = []models.Question{}
				} else {
					a.model.Questions = questions
				}
				a.model.UserAnswers = make([]int, len(a.model.Questions))
				a.model.UserTexts = make([]string, len(a.model.Questions))
			}

			return a, cmd
		case models.StateQuestion:
			updatedModel, cmd := a.model.HandleQuestionUpdate(msg)
			a.model = updatedModel
			return a, cmd
		case models.StateResult:
			updatedModel, cmd := a.model.HandleResultUpdate(msg)
			a.model = updatedModel

			if a.model.State == models.StateFileSelection {
				csvFiles, err := csv.GetCSVFiles()
				if err != nil {
					log.Printf("Error loading CSV files: %v", err)
					return a, tea.Quit
				}
				a.model.CSVFiles = csvFiles
			}

			return a, cmd
		}
	}
	return a, nil
}

func (a App) View() string {
	switch a.model.State {
	case models.StateFileSelection:
		return ui.RenderFileSelection(a.model)
	case models.StateMenu:
		return ui.RenderMenu(a.model)
	case models.StateQuestion:
		if a.model.SelectedMode == models.ModeMultipleChoice {
			return ui.RenderMultipleChoiceQuestion(a.model)
		} else {
			return ui.RenderWrittenQuestion(a.model)
		}
	case models.StateResult:
		return ui.RenderResults(a.model)
	default:
		return "Unknown state"
	}
}

func main() {
	ui.SetVersionInfo(version, commit, buildTime)

	csvFiles, err := csv.GetCSVFiles()
	if err != nil {
		fmt.Printf("Error initializing CSV system: %v\n", err)
		os.Exit(1)
	}

	model := models.NewModel()
	model.CSVFiles = csvFiles

	app := App{model: model}

	p := tea.NewProgram(app, tea.WithAltScreen())
	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}

	if fApp, ok := finalModel.(App); ok {
		if fApp.model.RequestUpdate {
			fmt.Println("\nUpdating memotty to latest release...")
			if err := update.Run(); err != nil {
				fmt.Printf("Update failed: %v\n", err)
			}
		}
	}
}
