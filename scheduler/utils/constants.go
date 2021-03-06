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
package utils

const (
	// PluginName is the name of the CSI plugin.
	CSIPluginName = "carina.storage.io"
	// storage class disk group
	DeviceDiskKey = "carina.storage.io/disk-type"
	// pv csi VolumeAttributes
	VolumeDeviceNode = "carina.storage.io/node"
	// device plugin
	DeviceCapacityKeyPrefix = "carina.storage.io/"

	// bcahce scheduler
	VolumeBackendDiskType = "carina.storage.io/backend-disk-type"
	VolumeCacheDiskType   = "carina.storage.io/cache-disk-type"
	// value: 1-100 Cache Capacity Ratio
	VolumeCacheDiskRatio = "carina.storage.io/cache-disk-ratio"
)
