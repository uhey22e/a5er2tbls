package a5er2tbls

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/k1LoW/tbls/config"
)

func a5erSectionToTblsRelation(section io.Reader) config.AdditionalRelation {
	conf := config.AdditionalRelation{}
	rt, _ := regexp.Compile(`^Entity2=`)
	rpt, _ := regexp.Compile(`^Entity1=`)
	rc, _ := regexp.Compile(`^Fields2=`)
	rpc, _ := regexp.Compile(`^Fields1=`)
	sc := bufio.NewScanner(section)
	for sc.Scan() {
		b := sc.Bytes()
		if lrt := len(rt.Find(b)); lrt > 0 {
			conf.Table = string(b[lrt:])
			continue
		}
		if lrpt := len(rpt.Find(b)); lrpt > 0 {
			conf.ParentTable = string(b[lrpt:])
			continue
		}
		if lrc := len(rc.Find(b)); lrc > 0 {
			conf.Columns = append(conf.Columns, string(b[lrc:]))
			continue
		}
		if lrpc := len(rpc.Find(b)); lrpc > 0 {
			conf.ParentColumns = append(conf.Columns, string(b[lrpc:]))
			continue
		}
	}
	return conf
}

func splitA5erFile(filename string) ([]*bytes.Buffer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rh, _ := regexp.Compile(`^\[Relation\]$`)
	rt, _ := regexp.Compile(`^$`)

	// The .a5er file is composed of sections; "Relation", "Entity"...
	// Load every "Relation" sections into buffer
	sections := make([]*bytes.Buffer, 0, 32)
	st := 0
	buf := bytes.NewBuffer(make([]byte, 512))
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		b := sc.Bytes()
		if st == 0 && rh.Match(b) {
			st = 1
		}
		if st == 1 {
			buf.Write(b)
			buf.WriteString("\n")
		}
		if st == 1 && rt.Match(b) {
			st = 0
			sections = append(sections, bytes.NewBuffer(make([]byte, buf.Len())))
			buf.WriteTo(sections[len(sections)-1])
			buf.Reset()
		}
	}
	if st == 1 {
		sections = append(sections, bytes.NewBuffer(make([]byte, buf.Len())))
		buf.WriteTo(sections[len(sections)-1])
	}

	return sections, nil
}

// ParseRelations read the a5er file and get relationships as the tbls setting struct.
func ParseRelations(filename string) []config.AdditionalRelation {
	sections, err := splitA5erFile(filename)
	if err != nil {
		panic(err)
	}

	confs := make([]config.AdditionalRelation, len(sections))
	for i, s := range sections {
		confs[i] = a5erSectionToTblsRelation(s)
	}

	return confs
}
