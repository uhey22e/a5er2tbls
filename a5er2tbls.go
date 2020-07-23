package a5er2tbls

import (
	"bufio"
	"os"
	"regexp"

	"github.com/k1LoW/tbls/config"
)

func a5erSectionToTblsRelation(section []string) config.AdditionalRelation {
	conf := config.AdditionalRelation{}
	rt, _ := regexp.Compile(`^Entity2=(.*)$`)
	rpt, _ := regexp.Compile(`^Entity1=(.*)$`)
	rc, _ := regexp.Compile(`^Fields2=(.*)$`)
	rpc, _ := regexp.Compile(`^Fields1=(.*)$`)
	for _, l := range section {
		if v := rt.FindStringSubmatch(l); v != nil && len(v) == 2 {
			conf.Table = v[1]
			continue
		}
		if v := rpt.FindStringSubmatch(l); v != nil && len(v) == 2 {
			conf.ParentTable = v[1]
			continue
		}
		if v := rc.FindStringSubmatch(l); v != nil && len(v) == 2 {
			conf.Columns = append(conf.Columns, v[1])
			continue
		}
		if v := rpc.FindStringSubmatch(l); v != nil && len(v) == 2 {
			conf.ParentColumns = append(conf.ParentColumns, v[1])
			continue
		}
	}
	return conf
}

func splitA5erFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rh, _ := regexp.Compile(`^\[Relation\]$`)
	rt, _ := regexp.Compile(`^$`)

	// The .a5er file is composed of sections; "Relation", "Entity"...
	// Load every "Relation" sections into buffer
	sections := make([][]string, 0, 32)
	var buf []string
	st := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		scb := sc.Bytes()
		if st == 0 && rh.Match(scb) {
			st = 1
			buf = make([]string, 0, 32)
		}
		if st == 1 {
			buf = append(buf, sc.Text())
		}
		if st == 1 && rt.Match(scb) {
			sections = append(sections, buf)
			buf = nil
			st = 0
		}
	}
	if st == 1 {
		sections = append(sections, buf)
	}

	return sections, nil
}

// ParseRelations read the a5er file and get relationships as the tbls setting struct.
func ParseRelations(filename string) ([]config.AdditionalRelation, error) {
	sections, err := splitA5erFile(filename)
	if err != nil {
		return nil, err
	}

	confs := make([]config.AdditionalRelation, len(sections))
	for i, s := range sections {
		confs[i] = a5erSectionToTblsRelation(s)
	}

	return confs, nil
}
