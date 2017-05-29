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