package vars

import (
	"fmt"
	"strconv"
	"strings"
)

type stringTest struct {
	key  string
	in   string
	want string
}

type atobTest struct {
	key     string
	in      string
	want    bool
	wantErr error
}

type atofTest struct {
	key     string
	in      string
	want    string
	wantErr error
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

var atobTests = []atobTest{
	{"ATOB_1", "", false, strconv.ErrSyntax},
	{"ATOB_2", "asdf", false, strconv.ErrSyntax},
	{"ATOB_3", "0", false, nil},
	{"ATOB_4", "f", false, nil},
	{"ATOB_5", "F", false, nil},
	{"ATOB_6", "FALSE", false, nil},
	{"ATOB_7", "false", false, nil},
	{"ATOB_8", "False", false, nil},
	{"ATOB_9", "1", true, nil},
	{"ATOB_10", "t", true, nil},
	{"ATOB_11", "T", true, nil},
	{"ATOB_12", "TRUE", true, nil},
	{"ATOB_13", "true", true, nil},
	{"ATOB_14", "True", true, nil},
}

var atofTests = []atofTest{
	{"FLOAT_1", "", "0", strconv.ErrSyntax},
	{"FLOAT_2", "1", "1", nil},
	{"FLOAT_3", "+1", "1", nil},
	{"FLOAT_4", "1x", "0", strconv.ErrSyntax},
	{"FLOAT_5", "1.1.", "0", strconv.ErrSyntax},
	{"FLOAT_6", "1e23", "1e+23", nil},
	{"FLOAT_7", "1E23", "1e+23", nil},
	{"FLOAT_8", "100000000000000000000000", "1e+23", nil},
	{"FLOAT_9", "1e-100", "1e-100", nil},
	{"FLOAT_10", "123456700", "1.234567e+08", nil},
	{"FLOAT_11", "99999999999999974834176", "9.999999999999997e+22", nil},
	{"FLOAT_12", "100000000000000000000001", "1.0000000000000001e+23", nil},
	{"FLOAT_13", "100000000000000008388608", "1.0000000000000001e+23", nil},
	{"FLOAT_14", "100000000000000016777215", "1.0000000000000001e+23", nil},
	{"FLOAT_15", "100000000000000016777216", "1.0000000000000003e+23", nil},
	{"FLOAT_16", "-1", "-1", nil},
	{"FLOAT_17", "-0.1", "-0.1", nil},
	{"FLOAT_18", "-0", "-0", nil},
	{"FLOAT_19", "1e-20", "1e-20", nil},
	{"FLOAT_20", "625e-3", "0.625", nil},

	{"FLOAT_21", "0", "0", nil},
	{"FLOAT_22", "0e0", "0", nil},
	{"FLOAT_24", "-0e0", "-0", nil},
	{"FLOAT_25", "+0e0", "0", nil},
	{"FLOAT_26", "0e-0", "0", nil},
	{"FLOAT_27", "-0e-0", "-0", nil},
	{"FLOAT_28", "+0e-0", "0", nil},
	{"FLOAT_29", "0e+0", "0", nil},
	{"FLOAT_30", "-0e+0", "-0", nil},
	{"FLOAT_31", "+0e+0", "0", nil},
	{"FLOAT_32", "0e+01234567890123456789", "0", nil},
	{"FLOAT_33", "0.00e-01234567890123456789", "0", nil},
	{"FLOAT_34", "-0e+01234567890123456789", "-0", nil},
	{"FLOAT_35", "-0.00e-01234567890123456789", "-0", nil},
	{"FLOAT_36", "0e291", "0", nil},
	{"FLOAT_37", "0e292", "0", nil},
	{"FLOAT_38", "0e347", "0", nil},
	{"FLOAT_39", "0e348", "0", nil},
	{"FLOAT_40", "-0e291", "-0", nil},
	{"FLOAT_41", "-0e292", "-0", nil},
	{"FLOAT_42", "-0e347", "-0", nil},
	{"FLOAT_43", "-0e348", "-0", nil},

	{"FLOAT_44", "nan", "NaN", nil},
	{"FLOAT_45", "NaN", "NaN", nil},
	{"FLOAT_46", "NAN", "NaN", nil},

	{"FLOAT_47", "inf", "+Inf", nil},
	{"FLOAT_48", "-Inf", "-Inf", nil},
	{"FLOAT_49", "+INF", "+Inf", nil},
	{"FLOAT_50", "-Infinity", "-Inf", nil},
	{"FLOAT_51", "+INFINITY", "+Inf", nil},
	{"FLOAT_52", "Infinity", "+Inf", nil},

	{"FLOAT_53", "1.7976931348623157e308", "1.7976931348623157e+308", nil},
	{"FLOAT_54", "-1.7976931348623157e308", "-1.7976931348623157e+308", nil},

	{"FLOAT_55", "1.7976931348623159e308", "+Inf", strconv.ErrRange},
	{"FLOAT_56", "-1.7976931348623159e308", "-Inf", strconv.ErrRange},

	{"FLOAT_57", "1.7976931348623158e308", "1.7976931348623157e+308", nil},
	{"FLOAT_58", "-1.7976931348623158e308", "-1.7976931348623157e+308", nil},

	{"FLOAT_59", "1.797693134862315808e308", "+Inf", strconv.ErrRange},
	{"FLOAT_60", "-1.797693134862315808e308", "-Inf", strconv.ErrRange},

	{"FLOAT_61", "1e308", "1e+308", nil},
	{"FLOAT_62", "2e308", "+Inf", strconv.ErrRange},
	{"FLOAT_63", "1e309", "+Inf", strconv.ErrRange},

	{"FLOAT_64", "1e310", "+Inf", strconv.ErrRange},
	{"FLOAT_65", "-1e310", "-Inf", strconv.ErrRange},
	{"FLOAT_66", "1e400", "+Inf", strconv.ErrRange},
	{"FLOAT_67", "-1e400", "-Inf", strconv.ErrRange},
	{"FLOAT_68", "1e400000", "+Inf", strconv.ErrRange},
	{"FLOAT_69", "-1e400000", "-Inf", strconv.ErrRange},

	{"FLOAT_70", "1e-305", "1e-305", nil},
	{"FLOAT_71", "1e-306", "1e-306", nil},
	{"FLOAT_72", "1e-307", "1e-307", nil},
	{"FLOAT_73", "1e-308", "1e-308", nil},
	{"FLOAT_74", "1e-309", "1e-309", nil},
	{"FLOAT_75", "1e-310", "1e-310", nil},
	{"FLOAT_76", "1e-322", "1e-322", nil},

	{"FLOAT_77", "5e-324", "5e-324", nil},
	{"FLOAT_78", "4e-324", "5e-324", nil},
	{"FLOAT_79", "3e-324", "5e-324", nil},

	{"FLOAT_80", "2e-324", "0", nil},

	{"FLOAT_81", "1e-350", "0", nil},
	{"FLOAT_82", "1e-400000", "0", nil},

	{"FLOAT_83", "1e-4294967296", "0", nil},
	{"FLOAT_84", "1e+4294967296", "+Inf", strconv.ErrRange},
	{"FLOAT_85", "1e-18446744073709551616", "0", nil},
	{"FLOAT_86", "1e+18446744073709551616", "+Inf", strconv.ErrRange},

	{"FLOAT_87", "1e", "0", strconv.ErrSyntax},
	{"FLOAT_88", "1e-", "0", strconv.ErrSyntax},
	{"FLOAT_89", ".e-1", "0", strconv.ErrSyntax},
	{"FLOAT_90", "1\x00.2", "0", strconv.ErrSyntax},

	{"FLOAT_91", "2.2250738585072012e-308", "2.2250738585072014e-308", nil},

	{"FLOAT_92", "2.2250738585072011e-308", "2.225073858507201e-308", nil},

	{"FLOAT_93", "4.630813248087435e+307", "4.630813248087435e+307", nil},

	{"FLOAT_94", "22.222222222222222", "22.22222222222222", nil},
	{"FLOAT_95", "2." + strings.Repeat("2", 4000) + "e+1", "22.22222222222222", nil},

	{"FLOAT_96", "1.00000000000000011102230246251565404236316680908203125", "1", nil},

	{"FLOAT_97", "1.00000000000000011102230246251565404236316680908203124", "1", nil},

	{"FLOAT_98", "1.00000000000000011102230246251565404236316680908203126", "1.0000000000000002", nil},

	{"FLOAT_99", "1.00000000000000011102230246251565404236316680908203125" + strings.Repeat("0", 10000) + "1", "1.0000000000000002", nil},
}

func genAtobTestBytes() []byte {
	var out []byte
	for _, data := range atobTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

func genStringTestBytes() []byte {
	var out []byte
	for _, data := range stringTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

func genAtofTestBytes() []byte {
	var out []byte
	for _, data := range atofTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}
