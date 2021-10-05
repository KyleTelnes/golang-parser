package main
import (
  "fmt"
  "io/ioutil"
  "log"
	"os"
	"strings"
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
	body, err0 := ioutil.ReadFile(fileName)
	if err0 != nil {
  	log.Fatalf("unable to read file: %v", err0)
  }
	//create a new file to write tokens into
	tknFileName := strings.TrimRight(fileName, ".cpl")
	file, err1 := os.Create(tknFileName + ".tkn")

    if err1 != nil {
        log.Fatalf("unable to create/open file: %v", err1)
    }

    defer file.Close()

	//iterate through each character in 'body'
	for i := 0; i < len(string(body)); i++ {
		//Assign tokens to the lexemes
		switch {
		case string(body[i]) == "=":
			//write to token file
			_, err2 := file.WriteString("ASSIGN\n")

    	if err2 != nil {
        	log.Fatalf("unable to write to file: %v", err2)
    	}
		case string(body[i]) == ",":
			//write to token file
			_, err2 := file.WriteString("COMMA\n")

			if err2 != nil {
					log.Fatalf("unable to write to file: %v", err2)
			}
		case string(body[i]) == ";":
			//write to token file
			_, err2 := file.WriteString("SEMICOLON\n")

			if err2 != nil {
					log.Fatalf("unable to write to file: %v", err2)
			}
		case string(body[i]) == "(":
			//write to token file
			_, err2 := file.WriteString("LPAREN\n")

    	if err2 != nil {
        	log.Fatalf("unable to write to file: %v", err2)
    	}
		case string(body[i]) == ")":
			//write to token file
			_, err2 := file.WriteString("RPAREN\n")

			if err2 != nil {
					log.Fatalf("unable to write to file: %v", err2)
			}
		case string(body[i]) == ".":
			//write to token file
			_, err2 := file.WriteString("PERIOD\n")

			if err2 != nil {
					log.Fatalf("unable to write to file: %v", err2)
			}
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
			//write to token file
			_, err2 := file.WriteString("NUM " + num + "\n")

    	if err2 != nil {
        	log.Fatalf("unable to write to file: %v", err2)
    	}
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
				//write to token file
				_, err2 := file.WriteString("POINT\n")

	    	if err2 != nil {
	        	log.Fatalf("unable to write to file: %v", err2)
	    	}
			} else if idName == "triangle" {
				//write to token file
				_, err2 := file.WriteString("TRIANGLE\n")

	    	if err2 != nil {
	        	log.Fatalf("unable to write to file: %v", err2)
	    	}
			} else if idName == "test" {
				//write to token file
				_, err2 := file.WriteString("TEST\n")

	    	if err2 != nil {
	        	log.Fatalf("unable to write to file: %v", err2)
	    	}
			} else if idName == "square" {
				//write to token file
				_, err2 := file.WriteString("SQUARE\n")

	    	if err2 != nil {
	        	log.Fatalf("unable to write to file: %v", err2)
	    	}
			} else { //if it is not reserved, write as identifier
				//write to token file
				_, err2 := file.WriteString("ID " + idName + "\n")

	    	if err2 != nil {
	        	log.Fatalf("unable to write to file: %v", err2)
	    	}
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
