/*
 * MIT License
 *
 * Copyright (c) 2022-2024 Tochemey
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package actors

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPIDMap(t *testing.T) {
	// create the actor path
	actorPath := NewPath("Test", NewAddress("TestSys", "host", 444))
	// create the PID
	actorRef := &PID{actorPath: actorPath, fieldsLocker: &sync.RWMutex{}, stopLocker: &sync.Mutex{}}
	// create a new PID map
	pidMap := newPIDMap(5)
	// add to the map
	pidMap.set(actorRef)
	// assert the length of the map
	assert.EqualValues(t, 1, pidMap.len())
	// list the map
	lst := pidMap.pids()
	assert.Len(t, lst, 1)
	// fetch the inserted pid back
	actual, ok := pidMap.get(actorPath)
	assert.True(t, ok)
	assert.NotNil(t, actual)
	assert.IsType(t, new(PID), actual)
	// remove the pid from the map
	pidMap.delete(actorPath)
	// list the map
	lst = pidMap.pids()
	assert.Len(t, lst, 0)
	assert.EqualValues(t, 0, pidMap.len())
}
