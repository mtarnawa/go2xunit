package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type TestOut struct {
	Time    time.Time
	Action  string
	Package string
	Test    string
	Output  string
	Elapsed float64
}

type TestKey struct {
	Package string
	Test    string
}

type Test struct {
	// TBD
}

func parse(in io.Reader) ([]*Test, error) {
	file, err := os.Open("test.json")
	if err != nil {
		return nil, err
	}

	tests := map[TestKey][]*TestOut{}

	dec := json.NewDecoder(file)
	for {
		var out TestOut
		err := dec.Decode(&out)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		key := TestKey{out.Package, out.Test}
		parts := tests[key]
		if parts == nil {
			parts = []*TestOut{&out}
		} else {
			parts = append(parts, &out)
		}
		tests[key] = parts
	}

	// DEBUG
	for key, parts := range tests {
		fmt.Println(key)
		fmt.Println(len(parts))
		for _, part := range parts {
			fmt.Printf("\t%v", part)
		}
	}

	return nil, nil
}
