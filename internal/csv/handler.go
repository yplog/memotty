package csv

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yplog/memotty/internal/models"
)

func EnsureMemottyDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %w", err)
	}

	memottyDir := filepath.Join(homeDir, ".memotty")

	if _, err := os.Stat(memottyDir); os.IsNotExist(err) {
		err = os.MkdirAll(memottyDir, 0755)
		if err != nil {
			return "", fmt.Errorf("could not create .memotty directory: %w", err)
		}

		err = createSampleCSV(memottyDir)
		if err != nil {
			return "", fmt.Errorf("could not create sample CSV: %w", err)
		}
	}

	return memottyDir, nil
}

func createSampleCSV(dir string) error {
	sampleFile := filepath.Join(dir, "sample_questions.csv")

	file, err := os.Create(sampleFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	records := [][]string{
		{"What is the synonym of happy?", "joyful"},
		{"What is the antonym of cold?", "hot"},
		{"What does the word run mean in the context of exercise?", "to jog or sprint"},
		{"What part of speech is the word quickly?", "adverb"},
		{"What is the plural form of mouse?", "mice"},
		{"Which language does the word fiancé originate from", "french"},
		{"What is the past tense of go?", "went"},
		{"What does the prefix un- mean?", "not or opposite"},
		{"What is the comparative form of good?", "better"},
		{"What is the root word of beautiful?", "beauty"},
	}

	return writer.WriteAll(records)
}

func GetCSVFiles() ([]string, error) {
	memottyDir, err := EnsureMemottyDir()
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(filepath.Join(memottyDir, "*.csv"))
	if err != nil {
		return nil, fmt.Errorf("could not list CSV files: %w", err)
	}

	var csvFiles []string
	for _, file := range files {
		csvFiles = append(csvFiles, filepath.Base(file))
	}

	return csvFiles, nil
}

func LoadQuestionsFromCSV(filename string, mode models.QuizMode) ([]models.Question, error) {
	memottyDir, err := EnsureMemottyDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(memottyDir, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV file: %w", err)
	}

	var questions []models.Question
	var allAnswers []string

	for _, record := range records {
		if len(record) < 2 {
			continue
		}

		question := strings.TrimSpace(record[0])
		correctAnswer := strings.TrimSpace(record[1])

		if question == "" || correctAnswer == "" {
			continue
		}

		allAnswers = append(allAnswers, correctAnswer)

		if mode == models.ModeWrittenAnswer {
			waQuestion := models.Question{
				Question:    question,
				CorrectText: correctAnswer,
				Mode:        models.ModeWrittenAnswer,
			}
			questions = append(questions, waQuestion)
		}
	}

	if mode == models.ModeMultipleChoice {
		for i, record := range records {
			if len(record) < 2 {
				continue
			}

			question := strings.TrimSpace(record[0])
			correctAnswer := strings.TrimSpace(record[1])

			if question == "" || correctAnswer == "" {
				continue
			}

			mcQuestion := createMultipleChoiceQuestionFromAnswers(question, correctAnswer, allAnswers, i)
			questions = append(questions, mcQuestion)
		}
	}

	return questions, nil
}

func createMultipleChoiceQuestionFromAnswers(question, correctAnswer string, allAnswers []string, seed int) models.Question {
	distractors := getRandomDistractors(correctAnswer, allAnswers, seed)

	allOptions := append([]string{correctAnswer}, distractors...)

	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed)))
	r.Shuffle(len(allOptions), func(i, j int) {
		allOptions[i], allOptions[j] = allOptions[j], allOptions[i]
	})

	correctIndex := 0
	for i, option := range allOptions {
		if option == correctAnswer {
			correctIndex = i
			break
		}
	}

	return models.Question{
		Question: question,
		Options:  allOptions,
		Correct:  correctIndex,
		Mode:     models.ModeMultipleChoice,
	}
}

func getRandomDistractors(correctAnswer string, allAnswers []string, seed int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed*13)))

	var distractors []string
	var otherAnswers []string

	for _, answer := range allAnswers {
		if answer != correctAnswer {
			otherAnswers = append(otherAnswers, answer)
		}
	}

	if len(otherAnswers) > 0 {
		r.Shuffle(len(otherAnswers), func(i, j int) {
			otherAnswers[i], otherAnswers[j] = otherAnswers[j], otherAnswers[i]
		})

		maxDistractors := len(otherAnswers)
		if maxDistractors > 3 {
			maxDistractors = 3
		}

		distractors = otherAnswers[:maxDistractors]
	}

	return distractors
}
