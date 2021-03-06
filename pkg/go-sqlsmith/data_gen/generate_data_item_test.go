// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package data_gen

import (
	"github.com/zhouqiang-cl/wreck-it/pkg/go-sqlsmith/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	layout = "2006-01-02 15:04:05"
)

func TestSQLSmith_TestIfDaylightTime(t *testing.T) {
	assert.Equal(t, ifDaylightTime(util.TimeMustParse(layout, "1986-05-05 11:45:14")), true)
	assert.Equal(t, ifDaylightTime(util.TimeMustParse(layout, "1991-09-05 11:45:14")), true)
	assert.Equal(t, ifDaylightTime(util.TimeMustParse(layout, "1985-08-05 11:45:14")), false)
	assert.Equal(t, ifDaylightTime(util.TimeMustParse(layout, "1992-06-05 11:45:14")), false)
}
