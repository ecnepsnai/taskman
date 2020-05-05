package taskman

import (
	"fmt"
	"sync"
	"time"
)

// Instance describes a single instance of taskman
type Instance struct {
	data          *Data
	dataDir       string
	dataInterface DataInterface
	lock          *sync.RWMutex
	list          []Task
	Delay         time.Duration
}

// New create a new taskman instance
func New(dataDir string, dataInterface DataInterface) (*Instance, error) {
	if dataInterface == nil {
		log.Warn("No data interface provided, defaulting to InMemoryStore. Specify a data interface to silence this warning.")
		dataInterface = InMemoryDataStore
	}

	err := dataInterface.Setup(dataDir)
	if err != nil {
		return nil, err
	}
	data, err := dataInterface.Load()
	if err != nil {
		return nil, err
	}

	i := Instance{
		data:          data,
		dataDir:       dataDir,
		dataInterface: dataInterface,
		lock:          &sync.RWMutex{},
		list:          []Task{},
		Delay:         5 * time.Second,
	}
	return &i, nil
}

// Ephemeral create a new in-memory taskman instance
func Ephemeral() *Instance {
	instance, err := New("", InMemoryDataStore)
	if err != nil {
		panic(err)
	}
	return instance
}

// Start the taskman instance and watching for jobs to trigger.
// This will block whatever goroutine it is called on.
func (i *Instance) Start() error {
	for true {
		i.lock.RLock()
		for _, task := range i.list {
			if task.canRunNow() {
				if fatalErr := task.runNow(); fatalErr != nil {
					log.Fatal("Fatal error running task '%s': %s", task.Name, fatalErr.Error())
				}
			}
		}
		i.lock.RUnlock()

		time.Sleep(i.Delay)
	}

	// Should never happen BUT YOU NEVER KNOW?!?!
	return fmt.Errorf("instance loop exited")
}

// AddTask add a new task to the current instance
func (i *Instance) AddTask(task Task) {
	i.lock.Lock()
	defer i.lock.Unlock()

	log.Debug("Adding new task '%s'", task.Name)
	i.list = append(i.list, task)

	i.data.LastRun[task.Name] = 0
	i.saveData()
}

// RemoveTask remove any task with the given name from the current instance
func (i *Instance) RemoveTask(taskName string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	log.Debug("Removing task named '%s'", taskName)

	n := 0
	for _, x := range i.list {
		if x.Name != taskName {
			i.list[n] = x
			n++
		}
	}
	i.list = i.list[:n]
}

// AllTasks return a list of all tasks
func (i *Instance) AllTasks() []Task {
	return i.list
}

func (i *Instance) saveData() error {
	if err := i.dataInterface.Save(*i.data); err != nil {
		return err
	}
	return nil
}
