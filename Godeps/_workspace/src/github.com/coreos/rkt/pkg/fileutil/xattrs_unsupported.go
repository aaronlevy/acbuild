// Copyright 2014 Red Hat, Inc
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

// +build !linux

// These functions are from github.com/docker/docker/pkg/system

package fileutil

func Lgetxattr(path string, attr string) ([]byte, error) {
	return nil, ErrNotSupportedPlatform
}

func Lsetxattr(path string, attr string, data []byte, flags int) error {
	return ErrNotSupportedPlatform
}
