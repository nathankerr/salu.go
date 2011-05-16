%{
package main

import (
	"fmt"
)
%}

%union
{
	text string
}

%token <text> WORD
%token <text> STOP
%token EOF

%%
sen: WORD STOP
{
	fmt.Println("sentence", $1, $2)
	return 1
}
%%
