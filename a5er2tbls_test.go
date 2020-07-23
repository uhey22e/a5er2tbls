package a5er2tbls

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestA5erSectionToTblsRelation(t *testing.T) {
	data := strings.NewReader(`
[Relation]
Entity1=user
Entity2=post
RelationType1=2
RelationType2=3
Fields1=user_id
Fields2=author_user_id
Cardinarity1=
Cardinarity2=

`)
	res := a5erSectionToTblsRelation(data)
	t.Log(res)

	if res.ParentTable != "user" {
		t.Error("Failed to parse parent table.")
	}
	if res.Table != "post" {
		t.Error("Failed to parse child table.")
	}
	if res.ParentColumns[0] != "user_id" {
		t.Error("Failed to parse child table.")
	}
	if res.Columns[0] != "author_user_id" {
		t.Error("Failed to parse child table.")
	}
}

func TestSplitA5erFile(t *testing.T) {
	filename := filepath.Join("testdata", "erd.a5er")
	sections, err := splitA5erFile(filename)
	if err != nil {
		t.Error("Failed to split file.", err)
	}
	if len(sections) != 3 {
		t.Error("Result is invalid.")
	}

	for _, s := range sections {
		t.Log(s.String())
	}
}
