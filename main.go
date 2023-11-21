package main

import (
	"fmt"
	"time"
)

const (
	secondsPerMinute = 60
	secondsPerHour   = 60 * secondsPerMinute
	secondsPerDay    = 24 * secondsPerHour
	seconds20Day     = 20 * secondsPerDay
	seconds21Day     = 21 * secondsPerDay
	seconds2Hours    = 2 * secondsPerHour
)

const (
	daily  = 1
	weekly = 7
)

const (
	repeatTypeOnce   = 1 // 重复类型-只开当天
	repeatTypeDaily  = 2 // 重复类型-当天 && 之后每天
	repeatTypeWeekly = 3 // 重复类型-当天 && 之后每周
)

func main() {

}

func doTask() {
	fun := "do task"
	fmt.Println("1")
	timer := time.NewTimer(1 * time.Second)

	count := int64(0)
	// offset := int64(10)
	for {
		now := int64(1685383198) + count
		nextTickAtSec := GetNextHalfHourUnixSecond(now)
		// waitSeconds := nextTickAtSec - now + offset
		waitSeconds := nextTickAtSec - now
		timer.Reset(time.Duration(waitSeconds) * time.Second)

		fmt.Printf("%s start to loop task, nowTime:%v, nowTimeUnix:%v, nextTickAtSec:%v, waitSeconds:%v\n",
			fun, now, now, nextTickAtSec, waitSeconds)
		count += 3
		select {
		case <-timer.C:
			fmt.Println("I am doing...")
		}
	}
}

func GetNextHalfHourUnixSecond(timeStamp int64) int64 {

	return (timeStamp/3)*3 + 3
}

func TimeIn(t time.Time) time.Time {
	// 设置时区：中国上海市
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("111")
		return t
	}
	fmt.Println("222")
	return t.In(loc)
}

func zone(timestamp int64) {
	// Load the time zone location for America/New_York
	loc, err := time.LoadLocation("America/New_York")

	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	t := time.Unix(timestamp, 0).In(loc)
	fmt.Println("Current time in New York:", t)

	s := TimeIn(time.Unix(timestamp, 0))
	fmt.Println("Current time in Shanghai:", s)
}

func IsPoint(timestamp int64) bool {
	unixTime := time.Unix(timestamp, 0)
	unixTime = TimeIn(unixTime)
	_, min, sec := unixTime.Clock()
	// 整点(xx:00,xx:30)
	fmt.Println("min:", min)
	fmt.Println("sec:", sec)
	if (min == 0 || min == 30) && sec == 0 {
		return true
	}
	return false
}

func getUpdatePeriod(timestamp int64) (start, end int64) {
	start = timeZeroUnix(time.Now())                // 今天的凌晨时刻。
	end = timeZeroUnix(time.Unix(timestamp, 0)) - 1 // 要关闭的时间槽的前一天的结束时刻。
	if start >= end {
		start, end = 0, 0
	}
	return
}

func isPoint(timestamp int64) bool {
	unix := time.Unix(timestamp, 0)
	_, min, sec := unix.Clock()
	if (min == 0 || min == 30) && sec == 0 {
		return true
	}
	return false
}

func timezone() {
	timestamp := int64(1684677661)
	u := time.Unix(timestamp, 0)
	r := u.Format("2006-01-02 15:04:05")
	fmt.Println(r)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	// fmt.Printf("%T %v", loc, err)
	layout := "2006-01-02 15:04:05"
	value := "2023-05-21 22:01:01"
	t, _ := time.ParseInLocation(layout, value, loc)
	fmt.Println(t.Unix())
}

func getNextThreeWeekRangeTime() (start, end int64) {
	start = todayZeroUnix()
	end = start + seconds21Day
	return
}

func timerStudy() {
	fmt.Println("1")
	timeTicker := time.NewTimer(1 * time.Second)
	for {
		timeTicker.Reset(2 * time.Second)
		// <-timeTicker.C
		select {
		case <-timeTicker.C:
			fmt.Println("I am doing...")
		}
	}
}

func tickerStudy() {
	fmt.Println("1")
	timeTicker := time.NewTicker(1 * time.Second)
	for {
		timeTicker.Reset(2 * time.Second)
		<-timeTicker.C
		fmt.Println("2")
	}
}

func secondsAt2am(timestamp int64) int64 {
	start := timeZeroUnix(time.Unix(timestamp, 0))
	end := start + seconds2Hours
	if timestamp >= start && timestamp <= end {
		return end
	}
	return start + secondsPerDay + seconds2Hours
}

func GetNextFutureDayRangeTime() (start, end int64) {
	start = todayZeroUnix() + seconds21Day
	end = start + secondsPerDay
	return
}

func GetNextThreeWeekRangeTime() (start, end int64) {
	start = todayZeroUnix()
	end = start + seconds21Day
	return
}

func GetVisiblePeriodRangeTimeByStamp(timestamp int64) (start, end int64) {
	start = timeZeroUnix(time.Unix(timestamp, 0))
	end = todayZeroUnix() + seconds21Day
	return
}

func timeSlotsGenerator(begin, stop, repeat, timeslot int64) []int64 {
	if begin >= stop {
		return []int64{}
	}

	freq := getFreq(repeat)
	if freq == 0 {
		return []int64{}
	}

	offset := getOffset(timeslot)
	timestamp := begin + offset

	s := timeZeroUnix(time.Unix(begin, 0))
	e := timeZeroUnix(time.Unix(stop, 0))
	d := diffDaysBySeconds(s, e)

	return generateTimeSlots(timestamp, freq, d)
}

func getOffset(timeslot int64) (offset int64) {
	offset = timeslot - timeZeroUnix(time.Unix(timeslot, 0))
	return
}

func getFreq(r int64) int64 {
	switch r {
	case repeatTypeDaily:
		return daily
	case repeatTypeWeekly:
		return weekly
	default:
		return 0
	}
}

func generateTimeSlots(timestamp, freq, d int64) []int64 {
	timeslots := make([]int64, 0)
	i := int64(0)
	for i < d {
		timeslot := timestamp + i*secondsPerDay
		timeslots = append(timeslots, timeslot)
		i += freq
	}
	return timeslots
}

func selectStudy() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "one"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "two"
			time.Sleep(300 * time.Millisecond)
		}
	}()

	for i := 0; i < 2000; i++ {
		time.Sleep(1000 * time.Millisecond)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("no msg to received")
		}

	}
}

func todayZeroUnix() int64 {
	now := time.Now()
	s := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return s.Unix()
}

func timeZeroUnix(t time.Time) int64 {
	// t = TimeIn(t)
	s := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return s.Unix()
}

func diffDaysBySeconds(s, e int64) int64 {
	return diffDaysByTime(time.Unix(s, 0), time.Unix(e, 0))
}

func diffDaysByTime(s, e time.Time) int64 {
	s = time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, s.Location())
	e = time.Date(e.Year(), e.Month(), e.Day(), 0, 0, 0, 0, e.Location())
	days := e.Sub(s).Hours() / 24
	return int64(days)
}
