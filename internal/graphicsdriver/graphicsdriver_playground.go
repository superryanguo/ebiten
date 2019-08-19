// Copyright 2019 The Ebiten Authors
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

// +build !android
// +build !darwin
// +build !freebsd
// +build !ios
// +build !js
// +build !linux,cgo !cgo
// +build !windows

package graphicsdriver

import (
	"github.com/hajimehoshi/ebiten/internal/driver"
)

func Get() driver.Graphics {
	if !driver.IsPlayground {
		panic("ebiten: a graphics driver is not implemented on this environment: isn't cgo disabled?")
	}
	// TODO: Implement this
	return nil
}