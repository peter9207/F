

to build parser libraries

get the grammar from https://github.com/antlr/grammars-v4/tree/master/javascript

get the jar file to generate by following instructions at https://github.com/antlr/antlr4/blob/master/doc/getting-started.md

generate code by running `antlr4 -Dlanguage=Go -o paser_test JavaScriptParser.g4`
