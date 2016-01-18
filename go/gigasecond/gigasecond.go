package gigasecond

import "time"

const TestVersion = 2

func AddGigasecond(birthday time.Time) time.Time {
	return birthday.Add(time.Duration(1e9) * time.Second)
}

var Birthday = time.Date(2016, time.January, 13, 0, 0, 0, 0, time.UTC)
