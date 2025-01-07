package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "Attimo/logging"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	datetimeFormat = "2006-01-02 15:04:05"
)

type TimeHandler struct {
	timeInput *inputModel
	value     string
	keys      selectionKeyMap
}

func NewTimeHandler(prompt string, logger *log.Logger) (*TimeHandler, error) {
	if logger == nil {
		return nil, fmt.Errorf(log.LoggerNilString)
	}

	timeInput, err := newInputModel(prompt, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create time input: %v", err)
	}
	timeInput.input.Focus()

	return &TimeHandler{
		timeInput: timeInput,
		keys:      newSelectionKeyMap(),
	}, nil
}

func (th *TimeHandler) parseTimeInput(input string) (string, error) {
	input = strings.TrimSpace(input)
	now := time.Now()

	switch {
	case input == "":
		return now.Format(datetimeFormat), nil

	case strings.HasPrefix(input, "+"):
		minutesStr := strings.TrimPrefix(input, "+")
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			return "", fmt.Errorf("invalid minutes format: %v", err)
		}

		return now.Add(time.Duration(minutes) * time.Minute).Format(datetimeFormat), nil

	case strings.HasPrefix(input, "-"):
		minutesStr := strings.TrimPrefix(input, "-")
		minutes, err := strconv.Atoi(minutesStr)
		if err != nil {
			return "", fmt.Errorf("invalid minutes format: %v", err)
		}

		return now.Add(-time.Duration(minutes) * time.Minute).Format(datetimeFormat), nil

	default:
		_, err := time.Parse(datetimeFormat, input)
		if err != nil {
			return "", fmt.Errorf("invalid time format: %v", err)
		}
		return input, nil
	}
}

func (th *TimeHandler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, th.keys.Quit) {
			return th, nil
		}

		if key.Matches(msg, th.keys.Enter) {
			// Parse and validate the time input
			parsedTime, err := th.parseTimeInput(th.timeInput.input.Value())
			if err != nil {
				th.timeInput.SetStatus(StatusError, fmt.Sprintf("Invalid time format: %v", err))
				return th, nil
			}
			// Store the final value
			th.value = parsedTime
			return th, tea.Quit
		}

		var cmd tea.Cmd
		th.timeInput.input, cmd = th.timeInput.input.Update(msg)
		return th, cmd
	}

	return th, nil
}

func (th *TimeHandler) View() string {
	return th.timeInput.View()
}

func (th *TimeHandler) GetValue() string {
	return th.value
}

func (th *TimeHandler) Init() tea.Cmd {
	return nil
}
