/*
   Gigasecond: add 1_000_000_000 seconds to a given Time
*/

package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1000000000) * time.Second
	return t.Add(gigasecond)
}
