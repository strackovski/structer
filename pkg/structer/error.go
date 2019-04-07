// Copyright (c) 2019 Vladimir Strackovski <vladimir.strackovski@gmail.com>
//
// This file is part of the structer project which is released under APACHE 2.0
// license. See LICENSE file or go to https://www.apache.org/licenses/LICENSE-2.0
// for full license details.

package structer

import "fmt"

// Check handles common error checking.
func Check(e error) {
    if e != nil {
        fmt.Printf("%s", e)
        return
    }
}
