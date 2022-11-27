// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute"
)

func TestWikilink(t *testing.T) {
	luteEngine := lute.New()

	test := "[reallink](linktext) [[wikilink]]"
	result := luteEngine.MarkdownStr("wikilink", test)
	expected := "<p><a href=\"linktext\">reallink</a> <a href=\"wikilink\" class=\"wikilink\">wikilink</a></p>\n"

	if result != expected {
		t.Fatalf("\ntried\n%s\nwanted\n%q\ngot\n%q\n", test, expected, result)
	}

}
