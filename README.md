# TaskMan

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/taskman?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/taskman)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/taskman)
[![Releases](https://img.shields.io/github/release/ecnepsnai/taskman/all.svg?style=flat-square)](https://github.com/ecnepsnai/taskman/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/taskman.svg?style=flat-square)](https://github.com/ecnepsnai/taskman/blob/master/LICENSE)

TaskMan is an advanced task scheduling system for golang applications. It attempts to implement an interface similar
to the Task Scheduler on Microsoft Windows.

# Usage

TaskMan runs as an instance in your go application. You add tasks to that instance and when running it will run those
tasks when it's their time.

**Start an Instance:**

For taskman to run properly it must be able to save some data. TaskMan includes two default data providers, a JSON
provider which saves a JSON file to your system, as well as an in-memory provider that is not persistent.

You can implement your own mechanism for saving and loading data by implementing the `taskman.DataInterface` interface.

```go
instance, err := taskman.New("/path/to/somewhere", taskman.JSONDataStore)
if err != nil {
	// There was an error setting up taskman
}

// Instances need to run on their own goroutine as they block
go instance.Start()
```

**Define a Task:**

```go
action := taskman.Action{
	// The UserDataProvider is called every time (if specified) this action is ran to get the user data that will
	// be provided to the main function
	UserDataProvider: func(taskName string) (interface{}, error) {
		return string("world"), nil
	},
	// This method is the main run method for the action, the user data is whatever the UserDataProvider returned, or nil
	Func: func(userData interface{}) error {
		greeting := userData.(string)
		fmt.Printf("Hello %s!\n", greeting)
		return nil
	},
}

task := taskman.Task{
	// The name of the task is how it is identified within the instance
	Name:        "Example task",
	// Tasks can have multiple triggers, only one has to be positive for the action to run
	Triggers:    []taskman.Trigger{trigger},
	Actions:     []taskman.Action{action}, // See below for examples on how to define triggers
}

// The task will become active as soon as you add it
instance.AddTask(task)
```

**Define Triggers:**

```go
// Run once, right now
onetime := taskman.TriggerOnce(time.Now())

// Run every 2 days after October 23 2077 9:47AM (UTC)
daily := taskman.TriggerDaily(time.Date(2077, 10, 23, 9, 47, 0, 0, time.UTC), 2)

// Run on the last day of March
custom := taskman.Trigger{
	Type:  taskman.Monthly,
	Start: time.Now(),
	Months: map[taskman.Month]bool{
		taskman.March: true,
	},
	DaysOfMonth: []taskman.DayOfMonth{
		taskman.Last,
	},
}
```
