package calculationservice

import (
	"fmt"
	"math"
	"time"
)

func CalculateZodiac(expression string) (string, error) {
    t, err := time.Parse("2006-01-02", expression)
    if err != nil {
        return "", fmt.Errorf("неверный формат даты: %w", err)
    }
    month := t.Month()
    day := t.Day()

    switch {
    case (month == 3 && day >= 21) || (month == 4 && day <= 19):
        return "Овен", nil
    case (month == 4 && day >= 20) || (month == 5 && day <= 20):
        return "Телец", nil
    case (month == 5 && day >= 21) || (month == 6 && day <= 20):
        return "Близнецы", nil
    case (month == 6 && day >= 21) || (month == 7 && day <= 22):
        return "Рак", nil
    case (month == 7 && day >= 23) || (month == 8 && day <= 22):
        return "Лев", nil
    case (month == 8 && day >= 23) || (month == 9 && day <= 22):
        return "Дева", nil
    case (month == 9 && day >= 23) || (month == 10 && day <= 22):
        return "Весы", nil
    case (month == 10 && day >= 23) || (month == 11 && day <= 21):
        return "Скорпион", nil
    case (month == 11 && day >= 22) || (month == 12 && day <= 21):
        return "Стрелец", nil
    case (month == 12 && day >= 22) || (month == 1 && day <= 19):
        return "Козерог", nil
    case (month == 1 && day >= 20) || (month == 2 && day <= 18):
        return "Водолей", nil
    case (month == 2 && day >= 19) || (month == 3 && day <= 20):
        return "Рыбы", nil
    default:
        return "", fmt.Errorf("дата вне диапазона зодиаков")
    }
}


func DaysUntilNextFullMoon(expression string) (int, error) {
    const moonCycle = 29.53 // Средняя длина лунного месяца
    // Дата последнего полнолуния (ориентировочно 19 августа 2024, можно обновлять)
    lastFullMoon := time.Date(2024, 8, 19, 18, 26, 0, 0, time.UTC)
    t, err := time.Parse("2006-01-02", expression)
    if err != nil {
        return 0, fmt.Errorf("неверный формат даты: %w", err)
    }
    daysSince := t.Sub(lastFullMoon).Hours() / 24
    // Сдвиг к следующему полнолунию
    mod := math.Mod(daysSince, moonCycle)
    if mod < 0 {
        mod += moonCycle
    }
    // Если сегодня полнолуние, выдаём 0
    if mod == 0 {
        return 0, nil
    }
    return int(moonCycle - mod + 0.5), nil // округление
}

func MoonPhaseString(expression string) (string, error) {
    const moonCycle = 29.53
    // Дата последнего полнолуния (см. предыдущая функция)
    lastFullMoon := time.Date(2024, 8, 19, 18, 26, 0, 0, time.UTC)
    t, err := time.Parse("2006-01-02", expression)
    if err != nil {
        return "", fmt.Errorf("неверный формат даты: %w", err)
    }
    daysSince := t.Sub(lastFullMoon).Hours() / 24
    // LibMod всегда положительный
    age := math.Mod(daysSince, moonCycle)
    if age < 0 {
        age += moonCycle
    }

    switch {
    case age < 1.5 || age > moonCycle-1.5:
        return "Полнолуние", nil
    case age < 7.4:
        return "Убывающая луна", nil
    case age < 8.4:
        return "Последняя четверть", nil
    case age < 14.8:
        return "Новолуние", nil
    case age < 15.8:
        return "Первая четверть", nil
    default:
        return "Растущая луна", nil
    }
}