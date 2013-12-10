// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package god

import (
	"path"

	"github.com/rakyll/god/config"
	"github.com/rakyll/god/remote"
)

type Options struct {
	Path        string
	IsRecursive bool
	IsForce     bool
}

type God struct {
	context *config.Context
	rem     *remote.Remote

	opts *Options
}

func New(context *config.Context, opts *Options) *God {
	var r *remote.Remote
	if context != nil {
		r = remote.New(context)
	}
	// TODO: should always start with /
	opts.Path = path.Clean(opts.Path)
	return &God{
		context: context,
		rem:     r,
		opts:    opts,
	}
}
