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

package log

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Logger = slog.Logger

var loggers sync.Map

func init() {
	slogOptions := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if slog.SourceKey == attr.Key {
				source := attr.Value.Any().(*slog.Source)
				idx := strings.LastIndexByte(source.File, '/')
				if idx == -1 {
					return attr
				}
				// Find the penultimate separator.
				idx = strings.LastIndexByte(source.File[:idx], '/')
				if idx == -1 {
					return attr
				}

				source.File = source.File[idx+1:]
			}

			return attr
		},
	}
	SetLogger("go-spring", slog.New(slog.NewTextHandler(os.Stdout, slogOptions)), true)
}

type namedLogger struct {
	name   string
	logger *Logger
}

func SetLogger(loggerName string, logger *Logger, primary ...bool) {
	named := &namedLogger{name: loggerName, logger: logger}
	loggers.Store(loggerName, named)

	if len(primary) > 0 && primary[0] {
		loggers.Store("", named)
	}
}

func GetLogger(loggerName string, typeName string) *Logger {
	if l, ok := loggers.Load(loggerName); ok {
		named := l.(*namedLogger)
		return named.logger.With("logger", named.name, "type", filepath.Base(typeName))
	}
	return nil
}
