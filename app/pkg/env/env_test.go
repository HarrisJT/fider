package env_test

import (
	"os"
	"testing"

	. "github.com/getfider/fider/app/pkg/assert"
	"github.com/getfider/fider/app/pkg/env"
)

var envs = []struct {
	go_env string
	env    string
	isEnv  func() bool
}{
	{"test", "test", env.IsTest},
	{"development", "development", env.IsDevelopment},
	{"production", "production", env.IsProduction},
	{"anything", "production", env.IsProduction},
}

func TestPath(t *testing.T) {
	RegisterT(t)

	Expect(env.Path("etc/deep/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/deep/file.txt")
	Expect(env.Path("etc/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/file.txt")
	Expect(env.Path("///etc/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/file.txt")
	Expect(env.Path("/etc/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/file.txt")
	Expect(env.Path("./etc/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/file.txt")
	Expect(env.Path("file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/file.txt")
	Expect(env.Path("/file.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/file.txt")
	Expect(env.Path("")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider")

	Expect(env.Etc("a.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/a.txt")
	Expect(env.Etc("b.txt")).Equals(os.Getenv("GOPATH") + "/src/github.com/getfider/fider/etc/b.txt")
}

func TestIsEnvironment(t *testing.T) {
	RegisterT(t)

	current := env.Config.Environment
	defer func() {
		env.Config.Environment = current
	}()

	for _, testCase := range envs {
		env.Config.Environment = testCase.go_env
		actual := testCase.isEnv()
		Expect(actual).IsTrue()
	}
}

func TestHasLegal(t *testing.T) {
	RegisterT(t)

	Expect(env.HasLegal()).IsTrue()
}

func TestMultiTenantDomain(t *testing.T) {
	RegisterT(t)

	env.Config.HostDomain = "test.fider.io"
	Expect(env.MultiTenantDomain()).Equals(".test.fider.io")
	env.Config.HostDomain = "dev.fider.io"
	Expect(env.MultiTenantDomain()).Equals(".dev.fider.io")
	env.Config.HostDomain = "fider.io"
	Expect(env.MultiTenantDomain()).Equals(".fider.io")
}

func TestSubdomain(t *testing.T) {
	RegisterT(t)

	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("")

	env.Config.CDN.Host = "test.assets-fider.io:3000"

	Expect(env.Subdomain("demo.test.fider.io")).Equals("demo")
	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("demo")
	Expect(env.Subdomain("test.fider.io")).Equals("")
	Expect(env.Subdomain("test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("helloworld.com")).Equals("")

	env.Config.HostMode = "single"

	Expect(env.Subdomain("demo.test.fider.io")).Equals("")
	Expect(env.Subdomain("demo.test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("test.fider.io")).Equals("")
	Expect(env.Subdomain("test.assets-fider.io")).Equals("")
	Expect(env.Subdomain("helloworld.com")).Equals("")
}
