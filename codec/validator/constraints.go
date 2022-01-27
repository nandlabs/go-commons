package validator

import (
	"fmt"
	"reflect"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(val reflect.Value, typ reflect.Type, param string) error {
	switch typ.Kind() {
	case reflect.String:
		c, err := convertBool(param)
		if err != nil {
			return err
		}
		if c == true {
			in, _ := val.Interface().(string)
			if in == "" {
				return ErrRequired
			}
		}
	case reflect.Bool:
	case reflect.Int:
	case reflect.Float32:
	case reflect.Uint:
	}
	return nil
}

func nillable(val reflect.Value, typ reflect.Type, param string) error {
	return nil
}

func def(val reflect.Value, typ reflect.Type, param string) error {
	return nil
}

/**
Numerical Type Constraints
*/

func min(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in > cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		/*c, err := convertFloat(param)
		if err != nil {
			return err
		}
		valid = input.Float() < c*/
		valid = true
	case reflect.Float64:
		valid = true
	}
	if !valid {
		return ErrMin
	}
	return nil
}

func max(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in < cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint:
		c, err := convertUint(param, 0)
		if err != nil {
			return err
		}
		cUint := uint(c)
		in, _ := val.Interface().(uint)
		valid = in < cUint
		valid = true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		c, err := convertFloat(param, 32)
		if err != nil {
			return err
		}
		cFloat := float32(c)
		in, _ := val.Interface().(float32)
		valid = in < cFloat
	case reflect.Float64:
		c, err := convertFloat(param, 64)
		if err != nil {
			return err
		}
		cFloat := c
		in, _ := val.Interface().(float64)
		valid = in < cFloat
	}
	if !valid {
		return ErrMax
	}
	return nil
}

/**
move the below functions to a generic function to consider the both min and exclusive-min
*/
func exclusiveMin(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in >= cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		/*c, err := convertFloat(param)
		if err != nil {
			return err
		}
		valid = input.Float() < c*/
		valid = true
	case reflect.Float64:
		valid = true
	}
	if !valid {
		return ErrMin
	}
	return nil
}

func exclusiveMax(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in <= cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint:
		c, err := convertUint(param, 0)
		if err != nil {
			return err
		}
		cUint := uint(c)
		in, _ := val.Interface().(uint)
		valid = in <= cUint
		valid = true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		c, err := convertFloat(param, 32)
		if err != nil {
			return err
		}
		cFloat := float32(c)
		in, _ := val.Interface().(float32)
		valid = in <= cFloat
	case reflect.Float64:
		c, err := convertFloat(param, 64)
		if err != nil {
			return err
		}
		cFloat := c
		in, _ := val.Interface().(float64)
		valid = in <= cFloat
	}
	if !valid {
		return ErrMax
	}
	return nil
}

func multipleOf(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	in, _ := val.Interface().(int)
	c, err := convertInt(param, 0)
	cInt := int(c)
	if err != nil {
		return err
	}
	valid = in%cInt == 0
	if !valid {
		return ErrMultipleOf
	}
	return nil
}

/**
String Type Constraints
*/

func minLength(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	valid = lv > lc
	if !valid {
		return ErrMinLength
	}
	return nil
}

func maxLength(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	valid = lv < lc
	if !valid {
		return ErrMaxLength
	}
	return nil
}

func pattern(val reflect.Value, typ reflect.Type, param string) error {
	return nil
}
