package vars

import "fmt"

type stringTest struct {
	key  string
	in   string
	want string
}

var stringTests = []stringTest{
	{"GOARCH", "amd64", "amd64"},
	{"GOHOSTARCH", "amd", "amd"},
	{"GOHOSTOS", "linux", "linux"},
	{"GOOS", "linux", "linux"},
	{"GOPATH", "/go-workspace", "/go-workspace"},
	{"GOROOT", "/usr/lib/golang", "/usr/lib/golang"},
	{"GOTOOLDIR", "/usr/lib/golang/pkg/tool/linux_amd64", "/usr/lib/golang/pkg/tool/linux_amd64"},
	{"GCCGO", "gccgo", "gccgo"},
	{"CC", "gcc", "gcc"},
	{"GOGCCFLAGS", "-fPIC -m64 -pthread -fmessage-length=0", "-fPIC -m64 -pthread -fmessage-length=0"},
	{"CXX", "g++", "g++"},
	{"PKG_CONFIG", "pkg-config", "pkg-config"},
	{"CGO_ENABLED", "1", "1"},
	{"CGO_CFLAGS", "-g -O2", "-g -O2"},
	{"CGO_CPPFLAGS", "", ""},
	{"CGO_CXXFLAGS", "-g -O2", "-g -O2"},
	{"CGO_FFLAGS", "-g -O2", "-g -O2"},
	{"CGO_LDFLAGS", "-g -O2", "-g -O2"},
}

func genStringTestBytes() []byte {
	var out []byte
	for _, data := range stringTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}