%{
package main

import (
	"fmt"
)
%}

%union
{
	node string
	vvar int
}

%token <vvar> VAR

%%
sen: VAR | VAR sen
%%
