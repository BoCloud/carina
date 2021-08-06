/*
   Copyright @ 2021 bocloud <fushaosong@beyondcent.com>.

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
package bcache

import (
	"github.com/bocloud/carina/pkg/devicemanager/types"
)

type Bcache interface {
	// create bcache
	CreateBcache(dev, cacheDev string) error
	RemoveBcache(dev, cacheDev string) error

	//
	GetDeviceBcache(dev string) (string, error)
	RegisterDevice(dev ...string) error
	ShowDevice(dev string) (*types.BcacheDeviceInfo, error)
}
