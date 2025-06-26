# Mobikit Validator for Go

**Mobikit** is a lightweight, extensible, and developer-friendly input validation library for Go.  
It enables you to define flexible validation rules for user inputs or any dynamic data source, with support for custom rules, clean API, and localization.

Whether you're building RESTful APIs, CLI tools, or backend systems, Mobikit gives you an easy way to keep your inputs clean and predictable.

---

## ✨ Features

- ✅ Simple and intuitive API
- 🧩 Custom rule support
- 🌍 Localization-ready (JSON-based error messages)
- 🧪 Easy to test and extend
- 📦 Lightweight, zero external dependencies

---

## 📦 Installation

To install the package:

```bash
go get -u github.com/Armnajafi/mobikit-go
```

## Sample usage
``` GO
package main

import (
	mobikit_go "github.com/Armnajafi/mobikit-go"
	"github.com/Armnajafi/mobikit-go/rules"
)

func main() {
	validator := mobikit_go.NewValidator()

	emailRule := rules.EmailRule{
		Message: "Invalid email address",
	}

	rulesMap := map[string]mobikit_go.Rule{
		"email": emailRule,
	}

	inputs := map[string]interface{}{
		"email": "test@example.com",
	}
	result := validator.Validate(inputs, rulesMap)
	if !result.IsValid {
		print("email is invalid")
	}
}

```
