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

	"github.com/qingstor/log"
	"github.com/qingstor/log/level"
)

// ErrUnavailableLevel returns when level not available
var ErrUnavailableLevel = errors.New("level not available")

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

// GetLogger new a default logger
func GetLogger() *log.Logger {
	return log.New()
}

// GetLoggerWithLevelAndWriter new a logger with given level and writer
func GetLoggerWithLevelAndWriter(l string, w io.Writer) *log.Logger {
	lvl, _ := ParseLevel(l)
	var writer io.Writer = os.Stderr
	if w != nil {
		writer = w
	}
	e := log.ExecuteMatchWrite(
		// Only print log that level is higher than Debug.
		log.MatchHigherLevel(lvl),
		// Write into stderr.
		writer,
	)
	tf, _ := log.NewText(&log.TextConfig{
		// Use unix timestamp nano for time
		TimeFormat: log.TimeFormatUnixNano,
		// Use upper case level
		LevelFormat: level.UpperCase,
		EntryFormat: "[{level}] - {time} {value}",
	})
	return log.New().WithExecutor(e).WithTransformer(tf)
}
