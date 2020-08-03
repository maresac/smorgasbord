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
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	_ "github.com/kubism/smorgasbord/internal/flags"
	"github.com/kubism/smorgasbord/pkg/testutil"
	"github.com/kubism/smorgasbord/pkg/util"

	"github.com/spf13/cobra"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dex         *testutil.Dex
	serverAddr  string
	redirectURL string
)

func TestSmorgasbord(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cmd/smorgasbord")
}

var _ = BeforeSuite(func(done Done) {
	var err error
	serverPort, err := util.GetFreePort()
	Expect(err).ToNot(HaveOccurred())
	serverAddr = fmt.Sprintf("127.0.0.1:%d", serverPort)
	redirectURL = fmt.Sprintf("http://%s/auth/callback", serverAddr)
	dex, err = testutil.NewDex(redirectURL)
	Expect(err).ToNot(HaveOccurred())
	Expect(dex).ToNot(BeNil())
	close(done)
}, 240)

var _ = AfterSuite(func() {
	if dex != nil {
		_ = dex.Close()
	}
})

func executeCommandWithContext(ctx context.Context, newCommandFn func(io.Writer) *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root := newCommandFn(buf)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	err = root.ExecuteContext(ctx)
	return buf.String(), err
}

func validServerArgs() []string {
	return []string{
		fmt.Sprintf("--addr=%s", serverAddr),
		fmt.Sprintf("--client-id=%s", testutil.DexClientID),
		fmt.Sprintf("--client-secret=%s", testutil.DexClientSecret),
		fmt.Sprintf("--issuer-url=%s", dex.GetIssuerURL()),
		fmt.Sprintf("--redirect-url=%s", redirectURL),
		"--auth-code-url-appendix=\"&connector_id=mock\"",
		"--nonce=test",
	}

}