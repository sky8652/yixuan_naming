/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2019 HereweTech Co.LTD
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

/**
 * @file common_names.go
 * @package list
 * @author Dr.NP <np@corp.herewetech.com>
 * @since 08/02/2019
 */

package list

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commonNamesM map[string]int

// QueryCommonNames : Check and query common name
func QueryCommonNames(input string) int {
	if commonNamesM != nil {
		return commonNamesM[input]
	}

	return 0
}

// LoadCommonNames : Common names
func LoadCommonNames(dir string) (int, error) {
	var (
		fullPath string
		f        *os.File
		scanner  *bufio.Scanner
		line     string
		err      error
		total    int
	)

	commonNamesM = make(map[string]int)
	fullPath = fmt.Sprintf("%s/list/CommonChineseNames.txt", dir)
	f, err = os.Open(fullPath)
	if err != nil {
		commonNamesM = nil
		return 0, fmt.Errorf("Load common names list failed")
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() == true {
		line = strings.TrimSpace(scanner.Text())
		commonNamesM[line] = 1
		total++
	}

	f.Close()

	return total, nil
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
