package jlexer

type FieldState int

const (
	FieldMissing FieldState = iota
	FieldNull
	FieldPresent
	FieldDuplicate
)

type FieldsTracker map[string]FieldState

func (f FieldsTracker) Add(field string) {
	if f[field] == FieldMissing {
		f[field] = FieldPresent
	} else {
		f[field] = FieldDuplicate
	}
}

func (f FieldsTracker) AddNull(field string) {
	if f[field] == FieldMissing {
		f[field] = FieldNull
	} else {
		f[field] = FieldDuplicate
	}
}

func (f FieldsTracker) GetState(field string) FieldState {
	result := f[field]
	if result == FieldNull {
		return FieldMissing
	}
	return result
}
