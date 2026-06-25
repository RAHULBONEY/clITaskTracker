package task

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func SaveTasks() error {
	bytes, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
func LoadTasks() error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(bytes) == 0 {
		return nil
	}
	err = json.Unmarshal(bytes, &Tasks)
	if err != nil {
		return err
	}
	return nil
}
