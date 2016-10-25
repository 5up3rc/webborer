// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

import (
	"github.com/Matir/gobuster/logging"
	"testing"
	"time"
)

func TestRobotsModeStrings(t *testing.T) {
	if len(robotsModeStrings) != robotsModeMax {
		t.Errorf("RobotsModeStrings != enum: %d vs %d", len(robotsModeStrings), robotsModeMax)
	}
}

// Test some defaults
func TestNewScanSettings(t *testing.T) {
	ss := NewScanSettings()
	if ss == nil {
		t.Fatalf("NewScanSettings returned nil!")
	}
	var foundLogLevel bool
	for _, l := range logging.LogLevelStrings {
		if l == ss.LogLevel {
			foundLogLevel = true
		}
	}
	if !foundLogLevel {
		t.Errorf("Invalid default loglevel: %s", ss.LogLevel)
	}
	if len(ss.Extensions) < 1 {
		t.Errorf("No default extensions!")
	}
	if len(ss.SpiderCodes) < 1 {
		t.Errorf("No HTTP codes to spider!")
	}
	if !ss.flagsSet {
		t.Errorf("Flags not initialized!")
	}
}

func TestStringSliceFlag(t *testing.T) {
	f := StringSliceFlag{}
	if f.String() != "" {
		t.Error("Expected empty string for empty StringSliceFlag.")
	}
	f.slice = &[]string{}
	s := "a,b,c"
	if err := f.Set(s); err != nil {
		t.Errorf("Error when setting StringSliceFlag: %v", err)
	}
	if len(*f.slice) != 3 {
		t.Errorf("len(f.slice) != 3, = %d", len(*f.slice))
	}
	if f.String() != s {
		t.Errorf("Differing strings: \"%s\" vs \"%s\".", f.String(), s)
	}
}

func TestIntSliceFlag(t *testing.T) {
	f := IntSliceFlag{}
	if f.String() != "" {
		t.Error("Expected empty string for empty IntSliceFlag.")
	}
	f.slice = &[]int{}
	s := "1,2,3"
	if err := f.Set(s); err != nil {
		t.Errorf("Error when setting IntSliceFlag: %v", err)
	}
	if len(*f.slice) != 3 {
		t.Errorf("len(f.slice) != 3, = %d", len(*f.slice))
	}
	if f.String() != s {
		t.Errorf("Differing strings: \"%s\" vs \"%s\".", f.String(), s)
	}
	if err := f.Set("xyz"); err == nil {
		t.Error("Expected error when setting invalid IntSliceFlag.")
	}
}

func TestDurationFlag_Empty(t *testing.T) {
	f := DurationFlag{}
	if f.String() != "" {
		t.Error("Expected empty string for empty DurationFlag.")
	}
}

func TestDurationFlag_String(t *testing.T) {
	d := time.Second
	f := DurationFlag{&d}
	if f.String() != "1s" {
		t.Errorf("Expected \"1s\" for duration: \"%s\"", f.String())
	}
}

func TestDurationFlag_Set_Valid(t *testing.T) {
	d := time.Duration(0)
	f := DurationFlag{&d}
	if err := f.Set("1s"); err != nil {
		t.Errorf("Error setting DurationFlag: %v", err)
	}
}

func TestDurationFlag_Set_Invalid(t *testing.T) {
	d := time.Duration(0)
	f := DurationFlag{&d}
	if err := f.Set("blah"); err == nil {
		t.Error("Expected error setting DurationFlag.")
	}
}

func TestRobotsFlag_Empty(t *testing.T) {
	f := robotsFlag{}
	if f.String() != "ignore" {
		t.Errorf("Expected robots flag ignore, got %s.", f.String())
	}
}

func TestRobotsFlag_String(t *testing.T) {
	i := 0
	f := robotsFlag{&i}
	if f.String() != "ignore" {
		t.Errorf("Expected robots flag ignore, got %s.", f.String())
	}
}

func TestRobotsFlag_Set_Valid(t *testing.T) {
	i := 0
	f := robotsFlag{&i}
	if err := f.Set("obey"); err != nil {
		t.Errorf("Expected no error setting robots flag, got %v", err)
	}
	if i != ObeyRobots {
		t.Errorf("Expected flag to be %d, got %d.", ObeyRobots, i)
	}
}

func TestRobotsFlag_Set_Invalid(t *testing.T) {
	i := 0
	f := robotsFlag{&i}
	if err := f.Set("wtfmate"); err == nil {
		t.Error("Expected error setting flag, got nil.")
	}
	if i != 0 {
		t.Errorf("Expected flag unchanged during error, got %d.", i)
	}
}
