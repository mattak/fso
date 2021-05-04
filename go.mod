module github.com/mattak/fso

go 1.14

require (
	cloud.google.com/go/firestore v1.5.0 // indirect
	github.com/mattak/fso/cmd v0.0.0-00010101000000-000000000000
)

replace github.com/mattak/fso/cmd => ./cmd
