// context should be request specific and should be extra info that doesn't impact normal workflow
// one useful case is to add request ID
package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

// to avoid number collision
// type key is unexported, only this package has access to
type key int
const requestIDKey = key(42)

func Println(ctx context.Context, msg string) {
	if id, ok := ctx.Value(requestIDKey).(int64); !ok {
		log.Println("could not find request ID in context")
		return
	} else {
		log.Printf("[%d] %s", id, msg)
	}
}

// wrap handler with context
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
