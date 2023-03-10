// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/tellmeac/entgo-example/ent/schema"
	"github.com/tellmeac/entgo-example/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescStars is the schema descriptor for Stars field.
	userDescStars := userFields[0].Descriptor()
	// user.StarsValidator is a validator for the "Stars" field. It is called by the builders before save.
	user.StarsValidator = func() func(int) error {
		validators := userDescStars.Validators
		fns := [...]func(int) error{
			validators[0].(func(int) error),
			validators[1].(func(int) error),
		}
		return func(_Stars int) error {
			for _, fn := range fns {
				if err := fn(_Stars); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
