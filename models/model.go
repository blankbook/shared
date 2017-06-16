// Package model provides an interface for all models, ensuring all models 
// have validation functions of the same format
package models

import (
    "reflect"
    "fmt"
    "errors"
)

// Model is an interface for the models of the data from incoming requests
type Model interface {
    Validate() error
}

// ValidateRanges validates whether the values of the fields of the model 
// lie within the specified ranges
func ValidateRanges(p Model, requiredCols map[string] bool, colMin,
                    colMax map[string]int64) error {
    var errorMessage string
    const errMsgStrMin = "field %s min length is %d, "
    const errMsgStrMax = "field %s max length is %d, "
    const errMsgIntMin = "field %s min val is %d, "
    const errMsgIntMax = "field %s max val is %d, "
    const errMsgMissing = "field %s required, "

    pRefl := reflect.ValueOf(p).Elem()
    pType := pRefl.Type()

    for i := 0; i < pRefl.NumField(); i++ {
        key := pType.Field(i).Name

        t := pRefl.Field(i).Type()
        var strval string
        var intval int64

        if t == reflect.TypeOf(strval) {
            strval := pRefl.Field(i).String()
            if strval == "" {
                continue;
            }
            if targ, ok := colMin[key]; ok && int64(len(strval)) < targ {
                errorMessage += fmt.Sprintf(errMsgStrMin, key, targ)
            }
            if targ, ok := colMax[key]; ok && int64(len(strval)) > targ {
                errorMessage += fmt.Sprintf(errMsgStrMax, key, targ)
            }
        } else if t == reflect.TypeOf(intval) {
            intval := pRefl.Field(i).Int()
            if targ, ok := colMin[key]; ok && intval < targ {
                errorMessage += fmt.Sprintf(errMsgIntMin, key, targ)
            }
            if targ, ok := colMax[key]; ok && intval > targ {
                errorMessage += fmt.Sprintf(errMsgIntMax, key, targ)
            }
        }

        if _, ok := requiredCols[key]; ok {
            requiredCols[key] = false
        }
    }

    for k, v := range requiredCols {
        if v {
            errorMessage += fmt.Sprint(errMsgMissing, k)
        }
    }

    var err error
    if errorMessage != "" {
        err = errors.New(errorMessage[:len(errorMessage) - 2])
    }
    return err
}
