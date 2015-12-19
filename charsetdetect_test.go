package charsetdetect

import (
	"os"
	"fmt"
	"path"
	"strings"
	"testing"
	"io/ioutil"
	"path/filepath"
)

type TestCase struct {
	Lang	string
	Charset	string
	Path	string
}

func Test_DetectCharset(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Errorf("Unable to get current working directory: %v", err)
	}

	var (
		parts   []string
		lang    string
		tests = make(map[string][]TestCase)
		content []byte
	)

	visit := func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			parts = strings.Split(f.Name(), ".")
			if len(parts) < 2 {
				return fmt.Errorf("Unable to process file '%s'", path)
			}

			lang = parts[0]
			if lang == "none" {
				lang = "__"
			}

			test_case := TestCase{
				Lang	: lang,
				Charset	: parts[1],
				Path	: path,
			}

			if _, ok := tests[lang]; ok {
				tests[lang] = append(tests[lang], test_case)
			} else {
				tests[lang] = []TestCase{test_case}
			}
		}

		return err
	}

	err = filepath.Walk(path.Join(pwd, "test-data"), visit)
	if err != nil {
		t.Error(err)
	}

	if len(tests) == 0 {
		t.Errorf("Tests are empty")
	}

	for lang, test_cases := range tests {
		for _, test_case := range test_cases {
			content, err = ioutil.ReadFile(test_case.Path)
			if err != nil {
				t.Errorf("Unable to read file '%s': %v", test_case.Path, err)
			} else {
				charset, err := DetectCharset(content)
				if err != nil {
					t.Errorf("Unable to detect encoding of '%s': %v", test_case.Path, err)
				}
				if charset == test_case.Charset {
					t.Logf("%s: %20s == %s", lang, charset, test_case.Charset)
				} else {
					t.Errorf("%s: %20s != %s", lang, charset, test_case.Charset)
				}
			}

		}
	}
}
