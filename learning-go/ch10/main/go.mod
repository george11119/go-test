module example.com/main

go 1.23.0

toolchain go1.23.8

replace example.com/add => ../add

require example.com/add v0.0.0-00010101000000-000000000000

require golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
