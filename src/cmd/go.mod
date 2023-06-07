module cmd

go 1.21

replace golang.org/x/mod => github.com/millergarym/golang_mod v0.0.1

require (
	github.com/google/pprof v0.0.0-20230602150820-91b7bce49751
	golang.org/x/arch v0.3.0
	golang.org/x/mod v0.10.1-0.20230606122920-62c7e578f1a7
	golang.org/x/sync v0.2.1-0.20230523185436-4966af63bb0c
	golang.org/x/sys v0.8.1-0.20230523194307-b5c7a0975ddc
	golang.org/x/term v0.8.0
	golang.org/x/tools v0.9.3
)

require (
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20230524184225-eabc099b10ab // indirect
)
