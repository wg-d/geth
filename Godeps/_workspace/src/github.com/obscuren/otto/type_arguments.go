package otto

import (
	"strconv"
)

func (runtime *_runtime) newArgumentsObject(indexOfParameterName []string, environment _environment, length int) *_object {
	self := runtime.newClassObject("Arguments")

	for index, _ := range indexOfParameterName {
		name := strconv.FormatInt(int64(index), 10)
		objectDefineOwnProperty(self, name, _property{Value{}, 0111}, false)
	}

	self.objectClass = _classArguments
	self.value = _argumentsObject{
		indexOfParameterName: indexOfParameterName,
		environment:          environment,
	}

	self.prototype = runtime.Global.ObjectPrototype

	self.defineProperty("length", toValue_int(length), 0101, false)

	return self
}

type _argumentsObject struct {
	indexOfParameterName []string
	// function(abc, def, ghi)
	// indexOfParameterName[0] = "abc"
	// indexOfParameterName[1] = "def"
	// indexOfParameterName[2] = "ghi"
	// ...
	environment _environment
}

func (self0 _argumentsObject) clone(clone *_clone) _argumentsObject {
	indexOfParameterName := make([]string, len(self0.indexOfParameterName))
	copy(indexOfParameterName, self0.indexOfParameterName)
	return _argumentsObject{
		indexOfParameterName,
		clone.environment(self0.environment),
	}
}

func (self _argumentsObject) get(name string) (Value, bool) {
	index := stringToArrayIndex(name)
	if index >= 0 && index < int64(len(self.indexOfParameterName)) {
		name := self.indexOfParameterName[index]
		if name == "" {
			return Value{}, false
		}
		return self.environment.GetBindingValue(name, false), true
	}
	return Value{}, false
}

func (self _argumentsObject) put(name string, value Value) {
	index := stringToArrayIndex(name)
	name = self.indexOfParameterName[index]
	self.environment.SetMutableBinding(name, value, false)
}

func (self _argumentsObject) delete(name string) {
	index := stringToArrayIndex(name)
	self.indexOfParameterName[index] = ""
}

func argumentsGet(self *_object, name string) Value {
	if value, exists := self.value.(_argumentsObject).get(name); exists {
		return value
	}
	return objectGet(self, name)
}

func argumentsGetOwnProperty(self *_object, name string) *_property {
	property := objectGetOwnProperty(self, name)
	if value, exists := self.value.(_argumentsObject).get(name); exists {
		property.value = value
	}
	return property
}

func argumentsDefineOwnProperty(self *_object, name string, descriptor _property, throw bool) bool {
	if _, exists := self.value.(_argumentsObject).get(name); exists {
		if !objectDefineOwnProperty(self, name, descriptor, false) {
			return typeErrorResult(throw)
		}
		if value, valid := descriptor.value.(Value); valid {
			self.value.(_argumentsObject).put(name, value)
		}
		return true
	}
	return objectDefineOwnProperty(self, name, descriptor, throw)
}

func argumentsDelete(self *_object, name string, throw bool) bool {
	if !objectDelete(self, name, throw) {
		return false
	}
	if _, exists := self.value.(_argumentsObject).get(name); exists {
		self.value.(_argumentsObject).delete(name)
	}
	return true
}
