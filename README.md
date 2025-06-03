# 📚 Memotty - Interactive CSV-Based Quiz Application

Memotty is a modern, interactive terminal quiz application built with Go and Bubbletea. It features dynamic CSV loading, intelligent distractor generation, and a beautiful terminal interface.

## 🎯 Key Features

- ✅ **Dynamic CSV Loading**: Load questions from any CSV file in `~/.memotty/`
- 📊 **Adaptive Options**: Question option count adapts to available answers (2-4 options)
- 🔄 **Dual Quiz Modes**: Multiple choice and written answer modes
- 📈 **Detailed Results**: Comprehensive analysis with correct/incorrect breakdown

## 🚀 Quick Start

### Installation
```bash
git clone https://github.com/yplog/memotty
cd memotty
go mod tidy
```

### Run the Application
```bash
go run cmd/memotty/main.go
```

## 📋 CSV File Format

Place your CSV files in `~/.memotty/` directory:

```csv
What is the synonym of happy?,joyful
What is the antonym of cold?,hot
What does the word run mean in the context of exercise?,to jog or sprint
What part of speech is the word quickly?,adverb
What is the plural form of mouse?,mice
Which language does the word fiancé originate from?,french
What is the past tense of go?,went
What does the prefix un- mean?,not or opposite
What is the comparative form of good?,better
What is the root word of beautiful?,beauty
```

### CSV Requirements
- **Format**: `question,answer` (comma-separated)
- **Location**: `~/.memotty/*.csv`
- **No headers**: Start directly with question data

## 🔧 Technical Features

### Dependencies
- **[Bubbletea](https://github.com/charmbracelet/bubbletea)**: Terminal UI framework
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)**: Styling and layout
- **Go 1.21+**: Modern Go features

## 🛠️ Customization

### Adding New Questions
1. Create/edit CSV file in `~/.memotty/`
2. Follow the `question,answer` format
3. Restart application to load new questions

### Creating Subject-Specific Quizzes
```csv
# math_basics.csv
What is 2 + 2?,4
What is the square root of 16?,4
What is 10 * 3?,30

# history_quiz.csv
When did World War II end?,1945
Who was the first US President?,George Washington
What year did the Berlin Wall fall?,1989
```

## 📈 Future Enhancements

- [ ] Timer-based quizzes
- [ ] Score tracking and statistics
- [ ] Multiple difficulty levels
- [ ] Multiple correct answers support (accept any of several valid answers)
- [ ] Export results to file
- [ ] Custom styling themes

## 🤝 Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.