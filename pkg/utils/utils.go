/*
Copyright 2021 NDD.

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

// BoolPtr return pointer to boolean
func BoolPtr(b bool) *bool { return &b }

// StringPtr return pointer to boolean
func StringPtr(s string) *string { return &s }

// IntPtr return pointer to int
func IntPtr(i int) *int { return &i }

// Int8Ptr return pointer to int8
func Int8Ptr(i int8) *int8 { return &i }

// Int16Ptr return pointer to int16
func Int16Ptr(i int16) *int16 { return &i }

// Int32Ptr return pointer to int32
func Int32Ptr(i int32) *int32 { return &i }

// Int64Ptr return pointer to int64
func Int64Ptr(i int64) *int64 { return &i }

// Uint8Ptr return pointer to uint8
func Uint8Ptr(ui uint8) *uint8 { return &ui }

// Uint16Ptr return pointer to uint16
func Uint16Ptr(ui uint16) *uint16 { return &ui }

// Uint32Ptr return pointer to uint32
func Uint32Ptr(ui uint32) *uint32 { return &ui }

// Uint64Ptr return pointer to uint64
func Uint64Ptr(ui uint64) *uint64 { return &ui }

func RemoveString(slice []string, s string) (result []string) {
	for _, v := range slice {
		if v != s {
			result = append(result, v)
		}
	}
	return result
}
