package main

import (
	"fmt"
	"time"
)

/*
	    Map weekday string to time.Weekday. The weekday string could be the full letters or first-three letters.
	    Ex.
	    Sunday: time.Sunday
		Sun   : time.Sunday
		Monday: time.Monday
		Mon   : time.Monday
		:
*/
var daysOfWeek = map[string]time.Weekday{}

func init() {
	for day := time.Sunday; day <= time.Saturday; day++ {
		dayString := day.String()
		daysOfWeek[dayString] = day
		daysOfWeek[dayString[0:3]] = day
	}
}

func dayFromString(dayString string) (time.Weekday, error) {
	if weekday, exists := daysOfWeek[dayString]; exists {
		return weekday, nil
	}

	return time.Saturday, fmt.Errorf("invalid weekday '%s'", dayString)
}

func main() {
	// Get input from user's entry
	var input string

	fmt.Println("Enter a week day string to test:")
	_, inputErr := fmt.Scanln(&input)
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}

	weekday, convertErr := dayFromString(input)

	if convertErr != nil {
		fmt.Println(convertErr)
		return
	}

	var result string

	switch weekday {
	case time.Sunday:
		result = "A limbo.. where you really want to do something fun, but you know you have something like school work to do, " +
			"so you don't do what you want to do because you think you're gonna start working soon, " +
			"but then it's hours later and you have done neither what you wanted to do nor what you need to do."
	case time.Monday:
		result = "It's the day of the week you look forward to the most because you get to spend it with a woman you are completely in love with. " +
			"She is usually one of a kind and the most beautiful woman you've ever known. most people hate mondays " +
			"because it's the first day of the work week. for some very fortunate people, " +
			"it's a day you can forget about all your troubles and and hold your baby in your arms."
	case time.Tuesday:
		result = "It's the day of the week during which Soylent Green may be purchased for consumption. " +
			"Soylent Green is not readily available the other days of the week due to limited production -- " +
			"but don't worry, the more people die from starvation, the more Soylent Green there will be for the remaining populace!"
	case time.Wednesday:
		result = "It's the day that you wear a spiderman suit and goggles then scream at yourself in the mirror."
	case time.Thursday:
		result = "It's the day of the week where nothing counts! If something happens only once, and it was a Thursday, " +
			"then it is as if it never happened at all. Hands down, the best day of the week."
	case time.Friday:
		result = "It's the day in which a 15 year old generally eats cereal for breakfast directly before going down to the bus stop. " +
			"After that one conveniently forgets about the bus and then proceed to see one's friends driving by. " +
			"Their car is nearly full but one still feels the need to make a decision on which seat to take, front or back. " +
			"One then proceeds to drive around singing in autotune and dancing awkwardly. " +
			"On Friday it is also very common to be extremely excited and to have a ball."
	case time.Saturday:
		result = "It's the national holiday that occurs once every week and 52 times in a year. It is devoted for males to spend times with each other."
	}

	fmt.Println(result)
}
