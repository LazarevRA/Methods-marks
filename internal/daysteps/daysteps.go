package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// DaySteps содержит все необходимые данные о дневных прогулках:
// количество шагов, длительность,
// а также данные из структуры personaldata.Personal
type DaySteps struct {
	personaldata.Personal               //персональные данные
	Steps                 int           //количество шагов
	Duration              time.Duration //продолжительность тренировки
}

// Parse принимает строку формата "678,0h50m" с количеством
// шагов и временем прогулки и записывает их в структуру
func (ds *DaySteps) Parse(datastring string) (err error) {
	input := strings.Split(datastring, ",")

	if len(input) != 2 {
		return errors.New("invalid input data format")
	}

	ds.Steps, err = strconv.Atoi(input[0])
	if err != nil {
		return fmt.Errorf("failed to get integer number of steps: %w", err)
	}

	ds.Duration, err = time.ParseDuration(input[1])
	if err != nil {
		return errors.New("failed to get walk duration")
	}

	return nil
}

// ActionInfo формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("the duration of the walk should be more than 0")
	}

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf(`Количество шагов: %d
		Дистанция составила %.2f км.
		Вы сожгли %.2f ккал.`, ds.Steps, spentenergy.Distance(ds.Steps), calories)

	return s, nil
}
