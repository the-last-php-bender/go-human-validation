# Go Human Validation

**Go Human Validation** is a Go library built on top of [`go-playground/validator`](https://github.com/go-playground/validator) that provides **human-readable validation error messages** for Go structs.  

It is lightweight, production-ready, and designed to make API responses and backend validations user-friendly.

---

## Features

- Converts validator errors into **human-readable messages**.
- Fully compatible with `go-playground/validator` tags.
- Lightweight and framework-agnostic.
- Easily extendable for custom messages and rules.
- Tested and production-ready.

---

## Installation

```bash
go get github.com/the-last-php-bender/go-human-validation
```

---

## Usage

### Define your struct with validation tags

```go
type RegisterRequest struct {
    FirstName string `json:"first_name" validate:"required,min=2"`
    LastName  string `json:"last_name" validate:"required,min=2"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=6"`
}
```

### Validate and get human-readable errors

```go
package main

import (
    "fmt"
    "github.com/the-last-php-bender/go-human-validation/pkg/validation"
)

func main() {
    req := RegisterRequest{
        FirstName: "",
        Email: "invalid-email",
        Password: "123",
    }

    errs := validation.Validate(req)
    if len(errs) > 0 {
        for _, e := range errs {
            fmt.Println(e) // e.g., "First name is required", "Email must be a valid email address"
        }
    }
}
```

---

## Contributing

If you want to contribute:

1. Fork the repository
2. Make your changes
3. Create a pull request

Don't forget to ⭐ the repo if you find it useful!

---

## License

This project is licensed under the MIT License.
