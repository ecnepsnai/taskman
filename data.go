package taskman

// Data describes the saved data for taskman
type Data struct {
	LastRun map[string]int64
}

var defaultData = Data{
	LastRun: map[string]int64{},
}
