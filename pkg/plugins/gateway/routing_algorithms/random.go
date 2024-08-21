/*
Copyright 2024 The Aibrix Team.

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

package routingalgorithms

import (
	"context"
	"math/rand/v2"

	v1 "k8s.io/api/core/v1"
)

type randomRouter struct {
}

func NewRandomRouter() Router {
	return randomRouter{}
}

func (r randomRouter) Get(ctx context.Context, pods []v1.Pod) (string, error) {
	pod := pods[rand.IntN(3)]
	return pod.Status.PodIP + ":8000", nil // TODO (varun): remove static port
}
