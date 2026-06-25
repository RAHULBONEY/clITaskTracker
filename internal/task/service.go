package task

import (
	"fmt"
	"hash/fnv"
	"time"
)

var Tasks []Task

func generateId(name string, t time.Time) int {
	uniqueString := fmt.Sprintf("%s-%d", name, t.UnixNano())
	h := fnv.New32a()
	h.Write([]byte(uniqueString))
	return int(h.Sum32())
}
func AddTask(name string) error {

	newTask := Task{
		Name:        name,
		IsCompleted: false,
		Time:        time.Now(),
		ID:          generateId(name, time.Now()),
	}
	Tasks = append(Tasks, newTask)
	return SaveTasks()

}

func GetTask() {
	if len(Tasks) == 0 {
		fmt.Println("Task list is empty")
		return
	}
	for _, value := range Tasks {
		fmt.Printf("The tasks name is %v,time of task %v, is complete %v\n", value.Name, value.Time, value.IsCompleted)
	}

}
func CompleteTask(id int) error {
	for i := range Tasks {
		if Tasks[i].ID == id {

			Tasks[i].IsCompleted = true
			Tasks[i].CompletedAt = time.Now()

			return SaveTasks()
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func DeleteTask(id int) error {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return SaveTasks()
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}
