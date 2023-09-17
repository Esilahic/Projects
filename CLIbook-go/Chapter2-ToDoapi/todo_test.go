package todo_test

import (
	"os"
	"testing"

	"github.com/Esilahic/CLIbook-go/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	TaskName := "New Task"
	l.Add(TaskName)

	if l[0].Task != TaskName {
		t.Errorf("Expected %s, got %s", TaskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	TaskName := "New Task"
	l.Add(TaskName)

	if l[0].Task != TaskName {
		t.Errorf("Expected %s, got %s", TaskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed")
	}
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{"New Task 1", "New Task 2", "New Task 3"}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %s, got %s", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %s, got %s", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	TaskName := "New Task"
	l1.Add(TaskName)

	if l1[0].Task != TaskName {
		t.Errorf("Expected %s, got %s", TaskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatal(err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatal(err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expected %s, got %s", l1[0].Task, l2[0].Task)
	}
}
