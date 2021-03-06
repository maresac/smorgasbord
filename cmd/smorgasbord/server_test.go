/*
Copyright 2020 Smorgasbord Authors

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

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kubism/smorgasbord/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	It("can start with valid parameters", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		// Use different port to test whether server can start to avoid clashing
		// with client test
		port, err := util.GetFreePort()
		Expect(err).ToNot(HaveOccurred())
		addr := fmt.Sprintf("127.0.0.1:%d", port)
		args := validServerArgs()
		args[0] = fmt.Sprintf("--addr=%s", addr)
		output, err := executeCommandWithContext(ctx, newServerCmd, validServerArgs()...)
		Expect(err).ToNot(HaveOccurred())
		Expect(output).ToNot(Equal(""))
	})
	It("fails without proper flags", func() {
		_, err := executeCommandWithContext(context.Background(), newServerCmd)
		Expect(err).To(HaveOccurred())
	})
})
