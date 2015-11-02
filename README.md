# behaviour driven development logging (BDDL)
BDDL is a behaviour driven development colorized logging package, making it easy to see what you are testing in your terminal.

Tips on how to use it
Use the dot import so you can write you bddl in a given-when-then style.
You can also add as much parameters as you want and they wil show up in your terminal with a different color.

#### Example:

```
package main

import (
"testing"

	. "github.com/idebruijn/bddl"
)
var email := example@bddl.com
var password := password123

Given(t,"I create a customer")
When(t,"I login with", email, "and", password)
Then(t, "my login was successful")
```

```
Given I create a customer
When I login with ezample@bddl.com and password123
Then my login was successful
```
