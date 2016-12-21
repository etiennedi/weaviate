package models




import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// StateDef state def
// swagger:model StateDef
type StateDef map[string]StateDefAnon

// Validate validates this state def
func (m StateDef) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(m) { // not required
		return nil
	}

	for k := range m {

		if swag.IsZero(m[k]) { // not required
			continue
		}

		if val, ok := m[k]; ok {

			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// StateDefAnon state def anon
// swagger:model StateDefAnon
type StateDefAnon struct {

	// Identifies what kind of resource this is. Value: the fixed string "weave#stateDef".
	Kind *string `json:"kind,omitempty"`

	// Minimal role required to view state.
	MinimalRole string `json:"minimalRole,omitempty"`

	// Name of the state field.
	Name string `json:"name,omitempty"`
}

// Validate validates this state def anon
func (m *StateDefAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMinimalRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stateDefAnonTypeMinimalRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manager","owner","user","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stateDefAnonTypeMinimalRolePropEnum = append(stateDefAnonTypeMinimalRolePropEnum, v)
	}
}

const (
	stateDefAnonMinimalRoleManager string = "manager"
	stateDefAnonMinimalRoleOwner   string = "owner"
	stateDefAnonMinimalRoleUser    string = "user"
	stateDefAnonMinimalRoleViewer  string = "viewer"
)

// prop value enum
func (m *StateDefAnon) validateMinimalRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stateDefAnonTypeMinimalRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StateDefAnon) validateMinimalRole(formats strfmt.Registry) error {

	if swag.IsZero(m.MinimalRole) { // not required
		return nil
	}

	// value enum
	if err := m.validateMinimalRoleEnum("minimalRole", "body", m.MinimalRole); err != nil {
		return err
	}

	return nil
}