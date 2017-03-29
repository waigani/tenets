- The Tenet applies only to the client scope by being placed in
the root of the server pkg folder. 
- The scope can be further refined with CLQL, e.g. to apply only on methods of "server type" structs etc
- The scope can be made more flexible with the use of variables and "or"s. E.g.

go.selector_expr:
  or:
    and:
      expr: "json"
      name: "Marshal"
    and:
      expr: "other"
      name: "Marshal"