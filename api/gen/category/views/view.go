// Code generated by goa v3.18.0, DO NOT EDIT.
//
// category views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// CategoryView is a type that runs validations on a projected type.
type CategoryView struct {
	// ID is the unique id of the category
	ID *uint
	// Name of category
	Name *string
}

var ()

// ValidateCategoryView runs the validations defined on CategoryView.
func ValidateCategoryView(result *CategoryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}