package hash_func_tests

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/nechel11/test_ozon/internal/utils"
)


func Test_hash_func(t *testing.T) {
	testCases := []struct {
		File_input        string
		File_output       string
		WantBool          bool
		
	}{
		{
			// 20k unique strings with different lenght
			File_input:	"./unique_strings",
			File_output: "output_true",
			WantBool: true,
		},
		{
			File_input:	"./not_unique_strings",
			File_output: "output_false",
			WantBool: false,
		},
	}

	//Act 
	for _, tc := range testCases{
		file, err := os.Open(tc.File_input)
		if err != nil {
			log.Fatal(err)
		}
		fo, err := os.Create(tc.File_output)
		if err != nil {
			panic(err)
		}
		defer fo.Close()
		scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := utils.Hash_func(scanner.Text())
		fo.WriteString(s)
		fo.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, tc.WantBool, duplicates_in_file(tc.File_output))
	defer os.Remove(tc.File_output)
	}	
}

func duplicates_in_file(s string) bool{
	counts := make(map[string]int)
	lineCountsInFiles := make(map[string]map[string]int)


	f, err := os.Open(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem reading %v: %v\n", s, err)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if lineCountsInFiles[line] == nil {
				lineCountsInFiles[line] = make(map[string]int)
		}
		lineCountsInFiles[line][s]++
	}
	f.Close()
	for _, n := range counts {
		if n > 1 {
			return false
		}
	}
	return true
}