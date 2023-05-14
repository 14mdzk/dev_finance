package helper

import "time"

type SqlDateFormat struct {
	time.Time
}

func (sdf *SqlDateFormat) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "" {
		return nil
	}

	tt, err := time.Parse(`"`+time.DateOnly+`"`, string(data))
	*sdf = SqlDateFormat{tt}

	return err
}
