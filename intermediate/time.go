package intermediate

import (
	"fmt"
	"time"
)

func main() {

	// Current local time
	fmt.Println(time.Now())

	// Specific time
	specificTime := time.Date(2025, time.April, 30, 10, 0, 0, 0, time.UTC)
	fmt.Println("Specific time:", specificTime)

	// Parse time
	parseTime, _ := time.Parse("2006-01-02", "2005-05-01")      // Mon Jan 2 15:04:05 MST 2006
	parseTime1, _ := time.Parse("06-01-02", "20-05-01")         // Mon Jan 2 15:04:05 MST 2006
	parseTime2, _ := time.Parse("06-1-2", "20-5-1")             // Mon Jan 2 15:04:05 MST 2006
	parseTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") // Mon Jan 2 15:04:05 MST 2006

	fmt.Println(parseTime)
	fmt.Println(parseTime1)
	fmt.Println(parseTime2)
	fmt.Println(parseTime3)

	// Formatting time
	t := time.Now()
	fmt.Println("Formatted time:", t.Format("Monday 06-01-02 15-04-02"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time:", t.Round(time.Hour))

	// loc, _ := time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2024, time.July, 8, 14, 16, 40, 0, time.UTC)

	// // Convert this to the specific time zone
	// tLocal := t.In(loc)

	// // Perform rouding
	// roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)

	// fmt.Println("Origin Time (UTC)", t)
	// fmt.Println("Origin Time (Local)", tLocal)
	// fmt.Println("Rounded Time (UTC)", roundedTime)
	// fmt.Println("Rounded Time (Local)", roundedTimeLocal)

	fmt.Println("Truncated time:", t.Truncate(time.Hour))

	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")

	// Convert time to location
	tInNY := time.Now().In(loc)
	fmt.Println("Ho Chi Minh Time:", tInNY)

	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))
}
