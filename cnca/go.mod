// Copyright 2019 Intel Corporation. All rights reserved.
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

module kubectl-cnca

replace cnca v0.0.0 => ./cmd

require (
	cnca v0.0.0
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	k8s.io/klog v1.0.0 // indirect
)
