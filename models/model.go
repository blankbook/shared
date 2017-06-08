// Package model provides an interface for all models, ensuring all models 
// have validation functions of the same format
package models

// Model is an interface for the models of the data from incoming requests
type Model interface {
    Validate() bool
}
