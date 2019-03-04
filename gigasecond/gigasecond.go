package gigasecond

import "time"

const GIGASECOND = time.Second * 1e9

func AddGigasecond(startTime time.Time) time.Time {
    return startTime.Add(GIGASECOND)
}
