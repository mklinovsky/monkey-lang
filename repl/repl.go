package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mklinovsky/monkey-lang/lexer"
	"github.com/mklinovsky/monkey-lang/token"
)

const PROMPT = ">> "

func Start(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Fprint(writer, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.CreateLexer(line)

		for currentToken := lexer.NextToken(); currentToken.Type != token.EOF; currentToken = lexer.NextToken() {
			fmt.Fprintf(writer, "%+v\n", currentToken)
		}
	}
}
