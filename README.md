# Create Your First Linter

Slides used in the talk "Create Your First Linter" in [gophercamp.cz](https://gophercamp.cz) 2026.

The goal of this presentation is to create a linter, and run:

- standalone
- inside [golangci-lint][golangci]

🔗[Link to the presentation][5]

## Tools

- I am using [slidev][slidev] to create the slides.

To start the slides, inside the [./slides/](./slides/) folder, run:

```bash
pnpm dev
```

## Go Projects

### Ast Example

Example on how the AST looks like for a small `.go` file

```bash
cd astexample && make t
```

### Unexported Constants Check

Exercise to implement Uber style guideline [Prefix Unexported Globals with _](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_).

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

To implement: [analyzer.go](./unexportedconstantscheck/analyzer.go)

### Custom Module

Small app to run [custom plugin linters][4].

[golangci]: https://golangci-lint.run/
[slidev]: https://sli.dev/
[3]: https://github.com/jj-vcs/jj
[4]: https://golangci-lint.run/plugins/module-plugins/
[5]: https://manuelarte.github.io/gophercamp-2026-create-your-first-linter
