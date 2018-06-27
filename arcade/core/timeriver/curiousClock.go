package timeriver

import "time"

const layout = "2006-01-02 15:04"

func curiousClock(someTime, leavingTime string) string {
	st, _ := time.Parse(layout, someTime)
	lt, _ := time.Parse(layout, leavingTime)
	difference := st.Sub(lt)

	return st.Add(difference).Format(layout)
}
