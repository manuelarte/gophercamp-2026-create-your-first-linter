func NewAnalyzer() *analysis.Analyzer {
  return &analysis.Analyzer{
    Name:     "astexample",
        Doc:      "astexample to show hierarchy",
        Run:      run,
        Requires: []*analysis.Analyzer{inspect.Analyzer},
  }
}

func run(pass *analysis.Pass) (any, error) {
  for _, file := range pass.Files {
    _ = ast.Print(pass.Fset, file)
  }
  return nil, nil
}