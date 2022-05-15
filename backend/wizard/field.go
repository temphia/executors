package simplewizard2

const ErrStringFieldEmpty = "Field cannot be empty"
const ErrStringWrongDataType = "Wrong data type"

func (sw *SimpleWizard) validateField(field *Field, data interface{}) string {

	if data == nil {
		if field.Optional {
			return ""
		} else {
			return ErrStringFieldEmpty
		}
	}

	switch field.Type {
	case BASIC_SHORTTEXT, BASIC_LONGTEXT, BASIC_SELECT:

		dstr, ok := data.(string)
		if !ok {
			return ErrStringWrongDataType
		}

		if len(dstr) == 0 {
			return ErrStringFieldEmpty
		}
	case BASIC_EMAIL:

	case BASIC_RANGE:
	case BASIC_MULTI_SELECT:
	case BASIC_PHONE:
	case BASIC_CHECKBOX:
	case BASIC_COLOR:
	case BASIC_DATE:
	case BASIC_DATETIME:

	case BASIC_NUMBER:

	}

	return ""

}
