// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

// Package logger provides support for logging to stdout and stderr.
package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/qingstor/log"
	"github.com/qingstor/log/level"
)

// ErrUnavailableLevel returns when level not available
var ErrUnavailableLevel = errors.New("level not available")

var mu = new(sync.Mutex)
var tf log.Transformer
var e log.Executor

// CheckLevel checks whether the log level is valid.
func CheckLevel(l string) error {
	if _, err := ParseLevel(l); err != nil {
		return fmt.Errorf(`%v: "%s"`, ErrUnavailableLevel, l)
	}
	return nil
}

// ParseLevel parse level from string into level.Level
func ParseLevel(l string) (level.Level, error) {
	l = strings.ToLower(l)
	for k, v := range level.Format[level.LowerCase] {
		if v == l {
			return level.Level(k), nil
		}
	}
	return level.Disable, ErrUnavailableLevel
}

// SetLevelAndWriter sets the log level and writer.
// Valid levels are "debug", "info", "warn", "error".
// Log with higher level will be written into writer.
// If writer nil, use Stderr as default.
func SetLevelAndWriter(l level.Level, w io.Writer) {
	mu.Lock()
	defer mu.Unlock()

	var writer io.Writer = os.Stderr
	if w != nil {
		writer = w
	}
	e = log.ExecuteMatchWrite(
		// Only print log that level is higher than Debug.
		log.MatchHigherLevel(l),
		// Write into stderr.
		writer,
	)
}

// SetMultipleExecutor sets log executor with multiple custom executor
// Such as writer higher level into one file and lower level into another file
func SetMultipleExecutor(ext ...log.Executor) {
	mu.Lock()
	defer mu.Unlock()

	e = log.ExecuteMultiple(ext...)
}

// GetLogger new a log entry with given executor and transformer
func GetLogger() *log.Logger {
	return log.New().WithExecutor(e).WithTransformer(tf)
}

func init() {
	var err error
	tf, err = log.NewText(&log.TextConfig{
		// Use unix timestamp nano for time
		TimeFormat: log.TimeFormatUnixNano,
		// Use upper case level
		LevelFormat: level.UpperCase,
		EntryFormat: "[{level}] - {time} {value}",
	})

	if err != nil {
		panic(fmt.Sprintf("failed to initialize QingStor SDK logger: %v", err))
	}

	// Only print warn and error logs in default
	SetLevelAndWriter(level.Info, nil)
}
