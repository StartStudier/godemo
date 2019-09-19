package refl

import "reflect"

func add(args []reflect.Value) (result []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var ret reflect.Value

	switch args[0].Kind() {
	case reflect.Int:
		n := 0

		for _, a := range args {
			n += int(a.Int())
		}

		ret = reflect.ValueOf(n)

	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}

		ret = reflect.ValueOf(ss)

	}

	result = append(result, ret)
	return
}
