package unit

import (
	"fmt"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/google/uuid"
)

type BasicInfo struct {
	// Name is the plugin/unit name
	Name string `json:"name"`
	//Description describes the units that are created by this plugin
	Description string `json:"description"`
	// Version is the current plugin version, should adhere to the semantic
	// versioning 2.0.0 spec
	Version string `json:"version"`
	// ID is the plugin unique ID, which should be kept the same, even as
	// the plugin version is advanced
	ID uuid.UUID `json:"id"`
}

func (info *BasicInfo) Verify() error {
	failureMsgs := []string{}

	if info.Name == "" {
		failureMsgs = append(failureMsgs, "name is empty")
	}

	if info.ID == uuid.Nil {
		failureMsgs = append(failureMsgs, "ID is empty")
	}

	if info.Version == "" {
		failureMsgs = append(failureMsgs, "version is empty")
	} else {
		if _, err := semver.ParseTolerant(info.Version); err != nil {
			failureMsgs = append(failureMsgs, err.Error())
		}
	}

	if info.Description == "" {
		failureMsgs = append(failureMsgs, "version is empty")
	}

	if len(failureMsgs) > 0 {
		failures := strings.Join(failureMsgs, ",")
		return fmt.Errorf("basic info is invalid: %s", failures)
	}

	return nil
}
