package proto

import (
	"encoding/json"
)

type Resume struct {
	Name string
	Age string
	School string
	Sex string
}

func (this *Resume) ClonePtr() *Resume {
	return this
}

func (this *Resume) Clone() *Resume {
	copyResume := new(Resume)
	data, err := json.Marshal(this)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, copyResume)
	if err != nil {
		return nil
	}
	return copyResume
}

