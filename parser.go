package main
import (
  "fmt"
  "io/ioutil"
  "log"
	"os"
)

// GenerateProlog
// using the results of the parser, generates the corresponding code in Prolog
// returns:
// a string with the Prolog code
func GenerateProlog () string {
	return "dummy text p"
}

// GenerateScheme
// using the results of the parser, generates the corresponding code in Scheme
// returns:
// a string with the Scheme code
func GenerateScheme () string {
	return "dummy text s"
}

// LexicallyAnalyze
// Lexically analyzes the text from the given file and writes the tokens
// to a file "tokenTable.tkn"
// returns:
// nothing
func LexicallyAnalyze (fileName string) {
	//read file and return all of the text into 'body'
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
  	log.Fatalf("unable to read file: %v", err)
  }
	//iterate through each character in 'body'
	for i := 0; i < len(string(body)); i++ {
		//Assign tokens to the lexemes
		switch {
		case string(body[i]) == "=": fmt.Println("ASSIGN")
		case string(body[i]) == ",": fmt.Println("COMMA")
		case string(body[i]) == ";": fmt.Println("SEMICOLON")
		case string(body[i]) == "(": fmt.Println("LPAREN")
		case string(body[i]) == ")": fmt.Println("RPAREN")
		case string(body[i]) == ".": fmt.Println("PERIOD")
		//these two empty cases are there to ignore whitespace chars
		case string(body[i]) == " ":
		case string(body[i]) == "\n":
		//assigns NUM Tokens
		case body[i] < 58 && body[i] > 47:
			num := string(body[i])
			len := 0
			for j := i + 1; body[j] < 58 && body[j] > 47; j++ {
				num += string(body[j])
				len++
			}
			i += len
			fmt.Println("NUM " + num)
		//assigns ID and other reserved word tokens
	  default:
			idName := string(body[i])
			idLength := 0
			for j := i + 1; body[j] < 123 && body[j] > 96; j++ {
				idName += string(body[j])
				idLength++
			}
			i += idLength
			//special cases for reserved words
			if idName == "point" {
				fmt.Println("POINT")
			} else if idName == "triangle" {
				fmt.Println("TRIANGLE")
			} else if idName == "test" {
				fmt.Println("TEST")
			} else if idName == "square" {
				fmt.Println("SQUARE")
			} else { //if it is not reserved, write as identifier
				fmt.Println("ID " + idName)
			}
		}
	}
}

func main() {
	//makes sure that there are the correct number of command line args
	if len(os.Args) < 3 {
		panic("Not enough command line arguments")
	} else if len(os.Args) > 3 {
		panic("Too many command line arguments")
	}

	//assign file name to read from
	fileName := os.Args[1]

	//do Lexical Analysis
	LexicallyAnalyze(fileName)

	//Decides whether to generate scheme or prolog code based on command line arg
	if os.Args[2] == "-s" {
		fmt.Println(GenerateScheme())
	} else if os.Args[2] == "-p" {
		fmt.Println(GenerateProlog())
	} else {
		panic("Include either \"-s\" or \"-p\" in the third command line argument")
	}
}
