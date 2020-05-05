package taskman

// Task describes a single task
type Task struct {
	Name        string
	Description string
	Triggers    []Trigger
	Actions     []Action
}

func (t Task) canRunNow() bool {
	return false
}

func (t Task) runNow() error {
	return nil
}
