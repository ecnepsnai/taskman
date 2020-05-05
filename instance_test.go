package taskman_test

import (
	"testing"

	"github.com/ecnepsnai/taskman"
)

func TestMutateTasks(t *testing.T) {
	instance := taskman.Ephemeral()

	task1 := taskman.Task{Name: "task1"}
	task2 := taskman.Task{Name: "task2"}
	task3 := taskman.Task{Name: "task3"}

	instance.AddTask(task1)
	instance.AddTask(task2)
	instance.AddTask(task3)

	expected := 3
	got := len(instance.AllTasks())

	if got != expected {
		t.Errorf("Returned list of tasks was not correct length. Expected %d got %d", expected, got)
	}

	instance.RemoveTask(task1.Name)

	expected = 2
	got = len(instance.AllTasks())

	if got != expected {
		t.Errorf("Returned list of tasks was not correct length. Expected %d got %d", expected, got)
	}
}
