package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Training содержит необходимые данные о тренировке
type Training struct {
	personaldata.Personal               //персональные данные
	Steps                 int           //количество шагов
	TrainingType          string        //тип тренировки (бег или хотьба)
	Duration              time.Duration //продолжительность тренировки
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m"
// и записывает данные в соответствующие поля структуры Training
//
// Параметры:
//
// datastring string - строка с данными о количестве шагов, виде и продолжительности тренировки.
func (t *Training) Parse(datastring string) (err error) {
	input := strings.Split(datastring, ",")

	if len(input) != 3 {
		return errors.New("invalid input data format")
	}

	t.Steps, err = strconv.Atoi(input[0])
	if err != nil {
		return fmt.Errorf("failed to get integer number of steps: %w", err)
	}

	t.TrainingType = input[1]
	if input[1] != "Бег" && input[1] != "Ходьба" {
		return errors.New("unknown training type")
	}

	t.Duration, err = time.ParseDuration(input[2])
	if err != nil {
		return errors.New("failed to get walk duration")
	}

	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке,
// исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {

	if t.Duration <= 0 {
		return "", errors.New("the duration of the walk should be more than 0")
	}

	var (
		err      error
		calories float64
	)

	switch t.TrainingType {

	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}

	s := fmt.Sprintf(`Тип тренировки: %s
		Длительность: %.2f ч.
		Дистанция: %.2f км.
		Скорость: %.2f км/ч
		Сожгли калорий: %.2f`, t.TrainingType,
		t.Duration.Hours(), spentenergy.Distance(t.Steps),
		spentenergy.MeanSpeed(t.Steps, t.Duration), calories)

	return s, nil
}
