/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gs

import (
	"io"
	"os"
	"path/filepath"
)

type Resource interface {
	io.ReadCloser
	Name() string
}

type ResourceLocator interface {
	Locate(filename string) ([]Resource, error)
}

// FileResourceLocator locate Resource from file system.
type FileResourceLocator struct {
	ConfigLocations []string `value:"${spring.config.locations:=config/}"`
}

func (locator *FileResourceLocator) Locate(filename string) ([]Resource, error) {
	var resources []Resource
	for _, location := range locator.ConfigLocations {
		fileLocation := filepath.Join(location, filename)
		file, err := os.Open(fileLocation)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		resources = append(resources, file)
	}
	return resources, nil
}
