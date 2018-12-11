// Copyright 2018 tinystack Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package govalidate

import (
    "reflect"
    "strconv"
)

func init() {
    validateHandlerMap[VALID_STR_MIN] = strMinHandler
    validateHandlerMap[VALID_STR_MAX] = strMaxHandler
    validateHandlerMap[VALID_STR_LEN] = strEqLenHandler
}

func strMinHandler(value reflect.Value, params []string) bool {
    return strMinMaxHandler(value, params, 'l')
}

func strMaxHandler(value reflect.Value, params []string) bool {
    return strMinMaxHandler(value, params, 'g')
}

func strEqLenHandler(value reflect.Value, params []string) bool {
    return strMinMaxHandler(value, params, 'e')
}

func strMinMaxHandler(value reflect.Value, params []string, op byte) bool {
    if len(params) != 1 {
        return true
    }
    slen, _ := strconv.Atoi(params[0])
    if value.Kind() == reflect.String {
        l := len(value.String())
        switch op {
        case 'l':
            if l < slen {
                return false
            }
        case 'e':
            if l != slen {
                return false
            }
        case 'g':
            if l > slen {
                return false
            }
        }
    }
    return true
}
