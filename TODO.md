add build task -> convert JSON rules into golang struct
	- make part of dockerfile?

add multiple rules files to be inputed
	bufio.MultiReader

add concurrency
	- each NewDiffItem = new go routine
	- buffer to max 100?
		- run benchmark tests (add to tests + makefile)