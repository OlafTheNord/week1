package service

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonReader(name string) error {
	employees := make([]byte, 0)

	r, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("cannot read file %s", err)
	}

	err = json.Unmarshal(r, &employees)
	if err != nil {
		return fmt.Errorf("cannot parse file %s", err)
	}
	return nil
}
