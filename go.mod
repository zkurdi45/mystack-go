module github.com/zkurdi45/mystack-go

go 1.24.4

replace github.com/zkurdi45/mystack-go => ./

require (
	github.com/a-h/templ v0.3.920
	github.com/go-chi/chi/v5 v5.2.2
	github.com/lib/pq v1.10.9
)

require github.com/google/go-cmp v0.7.0 // indirect
