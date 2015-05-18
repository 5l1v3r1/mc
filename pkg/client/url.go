/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"bytes"
	"net/url"
	"os"
)

// URL client url structure
type URL struct {
	Type   URLType
	Scheme string
	Host   string
	Path   string
}

// URLType - enum of different url types
type URLType int

// enum types
const (
	Unknown    URLType = iota // Unknown type
	Object                    // Minio and S3 compatible object storage
	Filesystem                // POSIX compatible file systems
)

// String converts type to string.
func (t URLType) String() string {
	switch t {
	case Object:
		return "Object"
	case Filesystem:
		return "Filesystem"
	default:
		return "Unknown"
	}
}

// Parse url parse
func Parse(urlStr string) *URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	if u.Scheme == "http" || u.Scheme == "https" {
		return &URL{
			Scheme: u.Scheme,
			Type:   Object,
			Host:   u.Host,
			Path:   u.Path,
		}
	}
	path, err := url.QueryUnescape(u.Path)
	if err != nil {
		return nil
	}
	return &URL{
		Type: Filesystem,
		Path: path,
	}
}

func (u *URL) String() string {
	var buf bytes.Buffer
	if u.Scheme != "" {
		buf.WriteString(u.Scheme)
		buf.WriteByte(':')
	}
	if u.Scheme != "" || u.Host != "" {
		buf.WriteString("//")
		if h := u.Host; h != "" {
			buf.WriteString(h)
		}
	}
	if u.Path != "" && u.Path[0] != '/' && u.Host != "" {
		buf.WriteByte(os.PathSeparator)
	}
	buf.WriteString(u.Path)
	return buf.String()
}
