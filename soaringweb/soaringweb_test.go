// Copyright 2014 The ezgliding Authors.
//
// This file is part of ezgliding.
//
// ezgliding is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ezgliding is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ezgliding.  If not, see <http://www.gnu.org/licenses/>.
//
// Author: Ricardo Rocha <rocha.porto@gmail.com>

package soaringweb

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ParseTest struct {
	t  string
	rg string
	c  string
	r  []Release
}

var parseTests = []ParseTest{
	{"single entry",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<small>[ 13 July 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
		},
	},
	{"multiple entries",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<small>[ 13 July 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
</li>
<small>[ 25 January 2012 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/250112__AIRSPACE_France_2501e.txt">OpenAir format</a>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
			Release{Location: "http://soaringweb.org/Airspace/FR/250112__AIRSPACE_France_2501e.txt",
				Region: "FR", Date: time.Date(2012, time.January, 25, 0, 0, 0, 0, time.UTC)},
		},
	},
	{"single entry inverted date",
		"FR",
		`
<html>
<body>
<small>[ Version 2014-05e (140708): Effective 01 May 2014 ]</small>
<li>
<a href="http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt">OpenAir format</a>
<small>[ 13 July 2014 ]</small>
</li>
</body>
</html>
`,
		[]Release{
			Release{Location: "http://soaringweb.org/Airspace/FR/140708__AIRSPACE_France_1405e.txt",
				Region: "FR", Date: time.Date(2014, time.July, 13, 0, 0, 0, 0, time.UTC)},
		},
	},
}

func TestListLocal(t *testing.T) {
	releases, err := List(".", []string{"FR"})
	if err != nil {
		t.Errorf("Failed to list releases :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListHTTP(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, _ := ioutil.ReadFile("./FR")
		io.WriteString(w, string(resp))
	}))
	defer ts.Close()

	releases, err := List(ts.URL, []string{"FR"})
	if err != nil {
		t.Errorf("Failed to list releases :: %v", err)
	}
	if len(releases) < 1 {
		t.Errorf("Got wrong number of releases :: %v", len(releases))
	}
}

func TestListEmpty(t *testing.T) {
	_, err := List("", nil)
	if err == nil {
		t.Errorf("List empty string should give error")
	}
}

func TestListMissing(t *testing.T) {
	_, err := List("./nonexisting.file", nil)
	if err == nil {
		t.Errorf("List non existing should give error")
	}
}

func TestParse(t *testing.T) {
	for i := range parseTests {
		test := parseTests[i]

		releases, err := parse(test.rg, []byte(test.c))
		if err != nil {
			t.Errorf("Failed to parse '%v' :: %v", test.t, err)
		}
		if len(releases) != len(test.r) {
			t.Errorf("Wrong num of releases in '%v' :: %v but expected %v",
				test.t, len(releases), len(test.r))
		}
		for r := range releases {
			var release = releases[r]
			var expected = test.r[r]
			if release.Date != expected.Date {
				t.Errorf("Wrong date in release :: got %v expected %v", release, expected)
			}
			if release.Location != expected.Location {
				t.Errorf("Wrong location in release :: got %v expected %v", release, expected)
			}
			if release.Region != expected.Region {
				t.Errorf("Wrong region in release :: got %v expected %v", release, expected)
			}
		}
	}
}
