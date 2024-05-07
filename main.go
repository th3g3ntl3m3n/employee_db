package main

import (
	"github.com/th3g3ntl3m3n/emplyee_db/cmd"
)

func main() {
	if err := cmd.NewServer().ListenAndServe(); err != nil {
		panic("can't start server")
	}
}
