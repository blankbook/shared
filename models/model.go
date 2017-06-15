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
func ValidateRanges(p Model, colMinLen map[string]int,
                    colMaxLen map[string]int) error {
    var errorMessage string
    const errMsgMin = "field %s required with min length %d, "
    const errMsgMax = "field %s required with max length %d, "

    pRefl := reflect.ValueOf(p).Elem()
    pType := pRefl.Type()

    for i := 0; i < pRefl.NumField(); i++ {
        key := pType.Field(i).Name
        val := pRefl.Field(i).Interface()
        if targ, ok := colMinLen[key]; ok {
            if val != targ {
                errorMessage += fmt.Sprintf(errMsgMin, key, targ)
            }
        }
        if targ, ok := colMaxLen[key]; ok {
            if val != targ {
                errorMessage += fmt.Sprintf(errMsgMax, key, targ)
            }
        }
    }

    var err error
    if errorMessage != "" {
        err = errors.New(errorMessage[:len(errorMessage) - 2])
    }
    return err
}
