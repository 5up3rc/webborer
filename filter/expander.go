// Copyright 2018 Google Inc. All Rights Reserved.
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

package filter

import (
	"github.com/matir/webborer/task"
	"github.com/matir/webborer/workqueue"
)

// An Expander is responsible for taking input URLs and expanding them to the
// various mutations to be processed.
type Expander interface {
	Expand(in <-chan *task.Task) <-chan *task.Task
	SetAddCount(workqueue.QueueAddCount)
}
