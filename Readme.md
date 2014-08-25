# GO Validator

This is a package, to _ease the Validation_ process of Data with many kind of validation methods. It returns An error list which is a representation of _key value pair of errors_ similar to Rails.

The Error list can be used independent of the `Validate`. `ErrorList` itself has some methods that _eases up Searching for errors on key and to Add/Append Errors._

###Validate

####Usage

```
	var validator = el.Validate{}
	validator.MaxLengthOfString("phone", 50, "00919999999999")
	validator.MatchRegExp("phone", "^([0-9])+$", "00919999999999")
```

Parameters:

* The First Parameter is always the :key to be used to log errors against


Output:

* Errors are a list of error objects `[]error`, and iss appended to the Validate object.

####Sample

This is how `Validator` looks like if the above Validations fail

```
Validate{
  ErrorList{
  	Errors: []KeyErrors{
  		{
  			Field: "phone",
  			Errors: []error{
  				error.New("phone: Max allowed length is 20 Characters, found 50")
  				error.New("phone: Format not correct")
  			},
  	},
  },
}
```

####Passing Accross Objects

```
// type Params struct

func (p *Params) Validate() *el.ErrorList {
	var validator = el.Validate{}
	p.User.Validate(&validator)
	p.Education.Validate(&validator)
	return &validator.ErrorList
}

// type User struct

func (p User) Validate(validator *el.Validate) {
	validator.MaxLengthOfString("phone", 50, p.Phone)
	validator.MatchRegExp("phone", "^([0-9])+$", p.Phone)
}

// type Education struct

func (p Education) Validate(validator *el.Validate) {
	validator.MaxLengthOfString("course_name", 60, p.Firstname)
	validator.MaxLengthOfString("course_email", 50, p.Email)
	validator.MatchEmail("course_email", p.Email)
}

```

All errors will be added to the same Validate object and you can use the validate.ErrorList object to use the errors.

###ErrorList

You can use ErrorList without using Validate.

####Usage

To add a new key and error, and to add more errors on to the same key, use method `Append`

```
errList := ErrorList{}
errList.Append("key", error.New("Error Message"))
```

To find a key and its Errors, use method `Find`. It returns an error as second parameter if the Key is not found.

```
errList.Find("key") // returns: KeyError, error(if Key not found, else nil)
```

To find index of a key and its Errors, use method `FindIndex`. It returns an error as second parameter if the Key is not found.

```
errList.FindIndex("key") // return: Index of KeyError in ErrorList, error(if Key not found, else nil)
```