package apigenerator

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Entries struct {
	Desc  *string  `xml:"desc"`
	Entry []*Entry `xml:"entry"`
}

//Parse a directory for all *.xml files and the one containing <entry>
func Parse(directory string) (e []*Entry, err error) {
	p := &parser{result: make([]*Entry, 0)}

	err = filepath.Walk(directory, p.walk)
	return p.result, err
}

type parser struct {
	result []*Entry
	inside bool
}

func (p *parser) walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if p.inside && info.IsDir() {
		return filepath.SkipDir
	}
	p.inside = true

	//path is the actual file
	if strings.HasSuffix(path, ".xml") {

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		defer file.Close()

		//log.Printf("Parsing %v", path)
		err = p.parseEntry(content)
		if err != nil {
			log.Printf("Not an entry %v", path)
			return err
		}
		err = p.parseEntries(content, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *parser) parseEntry(content []byte) (err error) {
	//try to read an "entry"
	e := new(Entry)
	err = xml.Unmarshal(content, e)
	if err != nil {
		return
	}

	if e.RawName == "" {
		return nil
	}
	p.result = append(p.result, e)
	return
}

func (p *parser) parseEntries(content []byte, path string) (err error) {
	//try to read an "entry"
	entries := make([]*Entries, 0, 2)
	err = xml.Unmarshal(content, &entries)
	if err != nil {
		return
	}

	if len(entries) == 0 {
		log.Printf("empty entries")
		return nil
	}

	for _, x := range entries {
		if x.Entry != nil {
			p.result = append(p.result, x.Entry...)
		}
		// if x.Desc != nil {
		// 	log.Printf("found entry %v", *x.Desc)
		// }
	}
	return nil
}
