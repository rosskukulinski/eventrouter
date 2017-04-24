/*
Copyright 2017 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sinks

import (
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/client-go/pkg/api/v1"
)

// StdoutSink is the most basic sink
// Useful when you already have ELK/EFK Stack
type StdoutSink struct {
	// TODO: create a channel and buffer for scaling
}

// NewStdoutSink will create a new
func NewStdoutSink() EventSinkInterface {
	return &StdoutSink{}
}

// UpdateEvents implements the EventSinkInterface
func (gs *StdoutSink) UpdateEvents(eNew *v1.Event, eOld *v1.Event) {
	var eJSON map[string]interface{}
	if eOld == nil {
		eJSON = map[string]interface{}{
			"verb":  "ADDED",
			"event": eNew,
		}
	} else {
		eJSON = map[string]interface{}{
			"verb":      "UPDATED",
			"event":     eNew,
			"old_event": eOld,
		}
	}
	if eJSONBytes, err := json.Marshal(eJSON); err == nil {
		fmt.Println(string(eJSONBytes))
	} else {
		fmt.Fprintf(os.Stderr, "Failed to json serialize event: %v", err)
	}
}
