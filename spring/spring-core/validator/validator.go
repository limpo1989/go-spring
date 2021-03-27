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

package validator

// Validator 参数校验器接口。
type Validator interface {
	Validate(i interface{}) error
}

var v Validator

// Init 初始化参数校验器。
func Init(r Validator) { v = r }

// Validate 参数校验。
func Validate(i interface{}) error {
	if v != nil {
		return v.Validate(i)
	}
	return nil
}
