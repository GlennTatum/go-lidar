/*
https://www.asprs.org/divisions-committees/lidar-division/laser-las-file-format-exchange-activities
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	LAS_FILE_SIG = "LASF"
)

type PBR_ITEM struct {
	i    int
	name string
	size int
}

// func PBR_Table(arg string) {
// 	var
// }

func main() {
	if err := run(os.Args); err != nil {
		panic(err)
	}
}

func run(args []string) error {
	fn := args[1]

	f, err := os.ReadFile(fn)
	if err != nil {
		return err
	}

	rd := bufio.NewReader(bytes.NewReader(f))
	file_signature(rd)

	return nil
}

/*
LAS VERSION 1.4 PUBLIC HEADER BLOCK ACCESSORS
(Ref. Table 3 LAS File Format Specification)

file_signature: complete
Point Data Record Format
- Note: There are various different formats of a point data record
*/
func file_signature(rd *bufio.Reader) (bool, error) {
	/*
		File Signature: The file signature must contain the four characters “LASF”, and it is required by
		the LAS specification. These four characters can be checked by user software as a quick look
		initial determination of file type.

		char[4]
		4 bytes
	*/
	var v string

	sig, err := rd.Peek(4) // char[4]
	if err != nil {
		return false, err
	}

	for _, b := range sig {
		ch := string(rune(b))
		fmt.Println(ch) // Debug

		ch += v

	}

	if v == LAS_FILE_SIG {
		return true, nil
	}

	return false, err
}
