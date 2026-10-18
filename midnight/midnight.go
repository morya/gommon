package midnight

import "time"

func GetMidNight() time.Time {
	n := time.Now()
	return GetMidNightOf(n)
}

func GetMidNightOf(t time.Time) time.Time {
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return midnight
}
