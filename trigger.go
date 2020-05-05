package taskman

import "time"

// TriggerType the types of triggers
type TriggerType string

// OneTime a "one time" trigger type where after triggering once, it does not trigger again unless a repeat is configured
var OneTime = TriggerType("one_time")

// Daily a trigger that runs at the specified time each day
var Daily = TriggerType("daily")

// Weekly a trigger that runs on the specified days of the week at the specified time
var Weekly = TriggerType("weekly")

// Monthly a trigger that runs on a specified day of the month at the specified time
var Monthly = TriggerType("monthly")

// DayOfWeek the day of weeks
type DayOfWeek string

// Monday named after the Moon
var Monday = DayOfWeek("monday")

// Tuesday named after Mars
var Tuesday = DayOfWeek("tuesday")

// Wednesday named after Mercury
var Wednesday = DayOfWeek("wednesday")

// Thursday named after Jupiter
var Thursday = DayOfWeek("thursday")

// Firday named after Venus
var Firday = DayOfWeek("firday")

// Saturday named after Saturn
var Saturday = DayOfWeek("saturday")

// Sunday named after the Sun (the star at the center of our solar system, not the bad british tabloid)
var Sunday = DayOfWeek("sunday")

// Month the month of the year
type Month string

// January named for the Roman god Janus, protector of gates and doorways.
var January = Month("january")

// February from the Latin word februa, "to cleanse."
var February = Month("february")

// March named for the Roman god of war, Mars.
var March = Month("march")

// April from the Latin word aperio, "to open (bud)," because plants begin to grow in this month.
var April = Month("april")

// May named for the Roman goddess Maia, who oversaw the growth of plants.
var May = Month("may")

// June named for the Roman goddess Juno, patroness of marriage and the well-being of women.
var June = Month("june")

// July named to honor Roman dictator Julius Caesar.
var July = Month("july")

// August named to honor the first Roman emperor (and grandnephew of Julius Caesar), Augustus Caesar.
var August = Month("august")

// September named after the Latin word septem, meaning "seven," because it was the seventh month in the Roman calendar.
var September = Month("september")

// October named after the latin word octo, meaning "eight," because it was the eighth month in the Roman calendar.
var October = Month("october")

// November named after the Latin word novem, meaning "nine," because it was the ninth month in the Roman calendar.
var November = Month("november")

// December named after the latin word decem, meaning "ten," because it was the tenth month in the Roman calendar.
var December = Month("december")

// DayOfMonth the day of the month. 0 represents the last day, 1 represents the first.
// If provided a day that isn't present on a month (for example 31 for February) it does not trigger.
type DayOfMonth uint

// Last the last day of the month
var Last = DayOfMonth(0)

// Trigger describes a trigger object
type Trigger struct {
	Type   TriggerType
	Start  time.Time
	Repeat *Repeat
	Expire *time.Time

	// Properties for Daily trigger types
	DailyRecurrence uint

	// Properties for Weekly trigger types
	WeeklyRecurrence uint
	DaysOfWeek       map[DayOfWeek]bool

	// Properties for Monthly trigger types
	Months      map[Month]bool
	DaysOfMonth []DayOfMonth
}

// TriggerOnce conveience method to return a one-time trigger at the given date
func TriggerOnce(at time.Time) Trigger {
	return Trigger{
		Type:  OneTime,
		Start: at,
	}
}

// TriggerDaily conveience method to return a daily trigger at the given time every n days
func TriggerDaily(at time.Time, everyNDays uint) Trigger {
	return Trigger{
		Type:            Daily,
		Start:           at,
		DailyRecurrence: everyNDays,
	}
}

func (t Trigger) didFire(lastRun time.Time) bool {
	return false
}

// overrideNow if not nil, the now() function will return this time. Used for testing.
var overrideNow *time.Time

// now return the current time
func now() time.Time {
	if overrideNow != nil {
		return *overrideNow
	}
	return time.Now()
}
