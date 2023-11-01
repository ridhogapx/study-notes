package main

import (
	"context"
	"fmt"
)

func Foo(ctx context.Context) {
  fmt.Println("Barrrr")	
}

func withVal(ctx context.Context) {
  fmt.Println("The value:", ctx.Value("foo") )
}

func main() {
  ctx := context.TODO()
  Foo(ctx)

  // With value
  ctxWithVal := context.WithValue(context.Background(), "foo", "bar")
 
  withVal(ctxWithVal)
}
