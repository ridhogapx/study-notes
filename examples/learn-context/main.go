package main

import (
	"context"
	"fmt"
)

func Foo(ctx context.Context) {
  fmt.Println("Barrrr")	
}

func main() {
  ctx := context.TODO()
  Foo(ctx)
}
