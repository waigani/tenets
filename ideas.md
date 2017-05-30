# Unblocked Go Routine

### Problem

Often you want a go routine to spawn workers and only wait for them to complete using a waitgroup:

```
func main () {
    var wg sync.WaitGroup
    work := []int{1,2,3,4,5}

    for _, v := range work {
        wg.Add(1)
        go func (int value) {
            doStuff(value)
            wg.Done()
        }(v)
    }

    wg.Wait()
}
```

An anti-pattern to catch is calling `wg.Add(1)` inside the spawned go routine. This will allow the parent goroutine to continue, because there is no guarantee that `wg.Add(1)` is called before we reach `wg.Wait()`.

### CLQL Idea

You could capture the most obvious cases with:

```
go.func_decl:
    go.assign_stmt:
        <identify waitgroup creation>
            name: $wgName
    <go routine>:
        go.func_expr:
            rhs:
                name: "Add"
            lhs:
                name: $wgName
        go.func_expr:
            rhs:
                name: "Done"
            lhs:
                name: $wgName
```

And we should be able to generate this whole query by selecting the anti-pattern code above.

# Sprintf an Error

### Problem

If you want to add some detail to an error as it's being returned up the call stack, an (example of an) obivous way to do that is:

```
errors.New(fmt.Sprintf("%v.Query(_) = _, %v", client, err)
```

but it would be tidier to use `errors.Errorf()`.

### CLQL Idea

```
<go.call_expr:
    <lhs>: "errors"
    <rhs>: "New"
    <args>:
        <lhs>: "fmt"
        <rhs>: "Sprintf"
```

This is a good example of a tenet that could live inside the Errors package and apply to projects that used Errors. 

# Deferring a Function

### Problem

Sometimes you want to defer a function, so you call it inside the deferred function:

```
defer func() {
    wg.Wait()
}()
```

We would prefer the cleaner:

```
defer wg.Wait()
```

### CLQL Idea

You need to find a deferred function that *only* has a single function inside it. I'm not sure how this would be done, but you could find the inner function like:

```
go.<deferal>:
    go.func_decl:
        <go.call_expr
```

Then it would be a matter of figuring out how to ensure that the call expression is the only child of the function declaration. Perhaps the not operator can be used to find all node types excluding funcs.

# Utilizing if-semicolon syntax

### Problem

Code generally becomes more readable when concerns are collected together into their own blocks and lines. For example, lines like:

```
err := os.Chdir(dir)
if err != nil {
	return "", errors.Trace(err)
}
```

Can be shortened to:

```
if err := os.Chdir(dir); err != nil {
    return "", errors.Trace(err)
}
```

### CLQL Idea

We can capture the specific case with something like:

```
go.call_expr:
    <return_val>: 
        type: "error"
    go.assign_stmt:
        name: $err
go.if_stmt:
    <lhs>:
        name: $err
    <rhs>:
        <literal>: "nil"
```

More generally we may want to capture all the cases where; we assign variables that are going to be used in an immediately proceeding if statement, but aren't going to be used anywhere else in the scope. Any variable like that can be declared in the init statement of the if statement because it doesn't need to be accessed anywere else.

We would also want to capture the cases where variables are assigned but not declared in the call expression; the assignment can be put into the if's init statement because it refers to the variable in the parent scope, so it will update it as before.

# Spelling and Grammar

### Problem

We need to be able to check spelling and grammar errors in comments in code. 

### CLQL Idea

One solution would be to create a lexicon that does NLP, and turns a string of English into a parse tree based on English grammar. The parse tree could add edges describing the elements of the string incorrect (for a simple black and white parser).

```
go.func_decl:
    <header comment>:
        <nlp.grammar_error
```

The above query would induce the ingest engine to parse all go header comments, and ingest them into the store. When parts of the comments can't be parsed, the NLP lexicon would add a grammar_error edge with a description of what the problem is. 

# Single Cased Select Statement

### Problem

A select statement with a single statement is pointless, and probably indicates that the author intended to add another case. 

```
select {
case stuff, ok := <-stuffc:
    dostuff()
}
```

### CLQL Idea

We need to identify a select statement with only one case, so ideally we would be able to count the number of cases:

```
go.select_stmt:
    count:
        amount: 1
        go.case_stmt:
```

or, if we could access a guaranteed unique identifier on the case node:

```
go.select_stmt:
    go.case_stmt:
        uuid: $var
    go.case_stmt:
        !uuid: $var
```