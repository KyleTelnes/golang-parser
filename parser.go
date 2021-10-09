//Kyle Telnes
//CSC 3310 CPL taught by Dr. Arias
//10-04-2021

package main
import (
  "fmt"
  "io/ioutil"
  "log"
	"os"
	"strings"
	//"strconv"
	"unicode"
)

// this struct will make it easier to translate between 4point and
// Prolog / Scheme
// represents components of POINT_DEF statment
// has id and code of other language fields
type AssignStatement struct {
	_id string
	_codeOtherLang string
}

// GenerateProlog
// using the results of the parser, generates the corresponding code in Prolog
// returns:
// a string with the Prolog code
func GenerateProlog (fileName string) string {
	fmt.Println("Generating Prolog Code...")
	//read file and return all of the text into 'body'
	body, err0 := ioutil.ReadFile(fileName)
	if err0 != nil {
  	log.Fatalf("unable to read file: %v", err0)
  }
	//remove whitespace from body
	fileContents := strings.Fields(string(body))

	//process the Slice to make it easier to iterate through the statements
	fileContents = strings.SplitAfter(strings.Join(fileContents, ""), "SEMICOLON")

	//final contains the final Prolog code
	final := ""

	//processed is a string Slice containing the processed statements
	var processed = make([]AssignStatement, 0, 100)

	//using the grammar, a statement will either be a POINT_DEF or TEST
	//POINT_DEF statments are the only ones that contain '=' (ASSIGN)
	for i := 0; i < len(fileContents); i++ {
		if strings.Contains(fileContents[i], "ASSIGN") {
			//extract the information needed to create the statement in prolog
			//only the id and nums are needed, so trim everything else out
			fileContents[i] = Simplify(fileContents[i], "ASSIGN")
			fileContents[i] = Simplify(fileContents[i], "POINT")
			fileContents[i] = Simplify(fileContents[i], "LPAREN")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "RPAREN")
			fileContents[i] = Simplify(fileContents[i], "SEMICOLON")
			fileContents[i] = Simplify(fileContents[i], "PERIOD")

			//extract the id
			id := ""
			for j := 2; unicode.IsLower(rune(fileContents[i][j])); j++ {
				id += string(fileContents[i][j])
			}
			//extract nums
			num1 := ""
			for j := strings.Index(fileContents[i], "NUM") + 3; unicode.IsDigit(rune(fileContents[i][j])); j++ {
				num1 += string(fileContents[i][j])
			}
			num2 := ""
			for j := strings.LastIndex(fileContents[i], "NUM") + 2; j < len(fileContents[i]); j++ {
				if unicode.IsDigit(rune(fileContents[i][j])) {
					num2 += string(fileContents[i][j])
				}
			}
			//grow slice by 1
			processed = processed[:len(processed) + 1]
			//add statement info
			processed[i]._id = id
			processed[i]._codeOtherLang = "point2d" + "(" + num1 + ", " + num2 + ")"

		} else { //case for a test statement
			//extract the information needed to create the statement in prolog
			//only the id and nums are needed, so trim everything else out
			fileContents[i] = Simplify(fileContents[i], "LPAREN")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "RPAREN")
			fileContents[i] = Simplify(fileContents[i], "SEMICOLON")
			fileContents[i] = Simplify(fileContents[i], "PERIOD")


			if strings.Contains(fileContents[i], "TRIANGLE") {
				//match the three parameters to the point ids
				id1 := ""
				id2 := ""
				id3 := ""
				bookmark := 0
				statement := ""
				//first ID will always be at index 14
				for j := 14; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id1 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id2 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; j < len(fileContents[i]); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id3 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := 0; j < len(processed); j++ {
					if processed[j]._id == id1 {
						statement += "(" + processed[j]._codeOtherLang + ", "
					} else if processed[j]._id == id2 {
						statement += processed[j]._codeOtherLang + ", "
					} else if processed[j]._id == id3 {
						statement += processed[j]._codeOtherLang + ")"
					}
				}
				//print all of the statements onto the screen
				fmt.Println("\nProcessing triangle test:\n")
				fmt.Println("query(line" + statement + ").")
				fmt.Println("query(triangle" + statement + ").")
				fmt.Println("query(vertical" + statement + ").")
				fmt.Println("query(horizontal" + statement + ").")
				fmt.Println("query(equilateral" + statement + ").")
				fmt.Println("query(isosceles" + statement + ").")
				fmt.Println("query(right" + statement + ").")
				fmt.Println("query(scalene" + statement + ").")
				fmt.Println("query(acute" + statement + ").")
				fmt.Println("query(obtuse" + statement + ").")

			} else { //case for square, same code as triangle but with one more id added
				//match the four parameters to the point ids
				id1 := ""
				id2 := ""
				id3 := ""
				id4 := ""
				bookmark := 0
				statement := ""
				//first ID will always be at index 12
				for j := 12; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id1 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id2 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id3 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; j < len(fileContents[i]); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id4 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := 0; j < len(processed); j++ {
					if processed[j]._id == id1 {
						statement += "(" + processed[j]._codeOtherLang + ", "
					} else if processed[j]._id == id2 {
						statement += processed[j]._codeOtherLang + ", "
					} else if processed[j]._id == id3 {
						statement += processed[j]._codeOtherLang + ", "
					} else if processed[j]._id == id4 {
						statement += processed[j]._codeOtherLang + ")"
					}
				}
				//print all of the statements onto the screen
				fmt.Println("\nProcessing square test:\n")
				fmt.Println("query(square" + statement + ").")
			}
		}
	}
	final += "\nQuery processing:\n"
	//this statement will always be at the end of the prolog program as far as I know
	final += "\nwriteln(T) :- write(T), nl.\nmain:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),\n\thalt."
	return final
}

// GenerateScheme
// using the results of the parser, generates the corresponding code in Scheme
// used same code as GenerateProlog function, with minor changes
// returns:
// a string with the Scheme code
func GenerateScheme (fileName string) string {
	fmt.Println("Generating Scheme Code...")
	//read file and return all of the text into 'body'
	body, err0 := ioutil.ReadFile(fileName)
	if err0 != nil {
  	log.Fatalf("unable to read file: %v", err0)
  }

	fileContents := strings.Fields(string(body))

	//process the Slice to make it easier to seperate the statments
	fileContents = strings.SplitAfter(strings.Join(fileContents, ""), "SEMICOLON")

	//processed is a string Slice containing the processed statements
	var processed = make([]AssignStatement, 0, 100)

	//using the grammar, a statement will either be a POINT_DEF or TEST
	//POINT_DEF statments are the only ones that contain '=' (ASSIGN)
	for i := 0; i < len(fileContents); i++ {
		if strings.Contains(fileContents[i], "ASSIGN") {
			//extract the information needed to create the statement in prolog
			//only the id and nums are needed, so trim everything else out
			fileContents[i] = Simplify(fileContents[i], "ASSIGN")
			fileContents[i] = Simplify(fileContents[i], "POINT")
			fileContents[i] = Simplify(fileContents[i], "LPAREN")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "RPAREN")
			fileContents[i] = Simplify(fileContents[i], "SEMICOLON")
			fileContents[i] = Simplify(fileContents[i], "PERIOD")

			//extract the id
			id := ""
			for j := 2; unicode.IsLower(rune(fileContents[i][j])); j++ {
				id += string(fileContents[i][j])
			}
			//extract nums
			num1 := ""
			for j := strings.Index(fileContents[i], "NUM") + 3; unicode.IsDigit(rune(fileContents[i][j])); j++ {
				num1 += string(fileContents[i][j])
			}
			num2 := ""
			for j := strings.LastIndex(fileContents[i], "NUM") + 2; j < len(fileContents[i]); j++ {
				if unicode.IsDigit(rune(fileContents[i][j])) {
					num2 += string(fileContents[i][j])
				}
			}
			//grow slice by 1
			processed = processed[:len(processed) + 1]
			//add statement info
			processed[i]._id = id
			processed[i]._codeOtherLang = "(" + "make-point" + " " + num1 + " " + num2 + ")"

		} else { //case for a test statement
			//extract the information needed to create the statement in prolog
			//only the id and nums are needed, so trim everything else out
			fileContents[i] = Simplify(fileContents[i], "LPAREN")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "COMMA")
			fileContents[i] = Simplify(fileContents[i], "RPAREN")
			fileContents[i] = Simplify(fileContents[i], "SEMICOLON")
			fileContents[i] = Simplify(fileContents[i], "PERIOD")


			if strings.Contains(fileContents[i], "TRIANGLE") {
				//match the three parameters to the point ids
				id1 := ""
				id2 := ""
				id3 := ""
				bookmark := 0
				statement := ""
				//first ID will always be at index 14
				for j := 14; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id1 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id2 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; j < len(fileContents[i]); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id3 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := 0; j < len(processed); j++ {
					if processed[j]._id == id1 {
						statement += processed[j]._codeOtherLang + " "
					} else if processed[j]._id == id2 {
						statement += processed[j]._codeOtherLang + " "
					} else if processed[j]._id == id3 {
						statement += processed[j]._codeOtherLang
					}
				}
				//print all of the statements onto the screen
				fmt.Println("(process-triangle " + statement + ")")
			} else { //case for square, same code as triangle but with one more id added
				//match the four parameters to the point ids
				id1 := ""
				id2 := ""
				id3 := ""
				id4 := ""
				bookmark := 0
				statement := ""
				//first ID will always be at index 12
				for j := 12; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id1 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id2 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; unicode.IsLower(rune(fileContents[i][j])); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id3 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := bookmark + 3; j < len(fileContents[i]); j++ {
					if unicode.IsLower(rune(fileContents[i][j])) {
						id4 += string(fileContents[i][j])
						bookmark = j
					}
				}
				for j := 0; j < len(processed); j++ {
					if processed[j]._id == id1 {
						statement += processed[j]._codeOtherLang + " "
					} else if processed[j]._id == id2 {
						statement += processed[j]._codeOtherLang + " "
					} else if processed[j]._id == id3 {
						statement += processed[j]._codeOtherLang + " "
					} else if processed[j]._id == id4 {
						statement += processed[j]._codeOtherLang
					}
				}
				//print all of the statements onto the screen
				fmt.Println("(process-square " + statement + ")")
			}
		}
	}

	//final contains the final Prolog code
	final := ""

	return final
}

// Simplify
// a function that takes in a string and a substring to
// get rid of the substring within the string
// returns a string with the substring removed
func Simplify(string1 string, substr string) string {
	if !strings.Contains(string1, substr) {
		return string1
	} else {
		i := strings.Index(string1, substr)
		firstPart := string1[:i]
		string1 = strings.Replace(string1[i:], substr, "", len(substr))
		return firstPart + string1
	}

}

// LexicallyAnalyze
// Lexically analyzes the text from the given file and writes the tokens
// to a file "tknFileName.tkn"
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
		//these three empty cases are there to ignore whitespace chars
		case string(body[i]) == " ":
		case string(body[i]) == "\n":
		case string(body[i]) == "\t":
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
		case body[i] < 123 && body[i] > 96:
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
			}	else if idName == "square" {
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
		//if the character is not recognized in the language
	  default:
				panic("Lexical Error, " + string(body[i]) + " not recognized")
		}
	}
}

// Syntactically Analyze
// analyzes the syntax of the program using tokens from tknFileName.tkn
// the idea is to use recursive descent parsing
// returns:
// nothing
func SyntacticallyAnalyze (fileName string) {
	//read file and return all of the text into 'body'
	body, err0 := ioutil.ReadFile(fileName)
	if err0 != nil {
  	log.Fatalf("unable to read file: %v", err0)
  }

	//put all of the tokens into a slice called fileContents
	//so that each token can be accessed by an iterator
	fileContents := strings.Split(string(body), "\n")

	//iterator value to be passed by reference into Parsing functions
	i := 0

	ParseSTART(fileContents, &i)

	fmt.Println("Lexical and Syntax Analysis Passed!")
}

// ParseSTART
// Parameters: fileContents is the slice that all of the tokens are
//						 stored in
//						 i is the iterator, passed by reference so that it can be
//						 updated with more ease in the subsequent functions
// represents the first grammar rule, calling on the ParseSTMT rule
// returns whether the Semantic Analysis Passed
// true if it did false if not (but the panic statements
// will kill the program before a value of false can reach here)
func ParseSTART (fileContents []string, i *int) bool {
	return ParseSTMT_LIST(fileContents, i)
}

// ParseSTMT_LIST
// represents the second grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParseSTMT_LIST(fileContents []string, i *int) bool {
	if ParseSTMT(fileContents, i) && fileContents[*i] == "PERIOD" {
		return true
	} else if fileContents[*i] == "SEMICOLON" { //alternate case for if there is more to the program
		//token has been used up, so add 1 to the iterator
		*i++
		//recursively call the rule again if there are more statements to be expected
		return ParseSTMT_LIST(fileContents, i)
	} else { //if the statements don't end with ';' or '.' , bad syntax
		panic("syntax error: semicolon or period expected")
		return false
	}
}

// ParseSTMT
// represents the third grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParseSTMT(fileContents []string, i *int) bool {
	if ParsePOINT_DEF(fileContents, i) {
		return true
	} else if ParseTEST(fileContents, i) {
		return true
	} else {
		panic("syntax error: incorrect statement")
		return false
	}
}

// ParsePOINT_DEF
// represents the fourth grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParsePOINT_DEF(fileContents []string, i *int) bool {
	//[:2] trims out the identifier name
	if fileContents[*i][:2] == "ID" && fileContents[*i + 1] == "ASSIGN" && fileContents[*i + 2] == "POINT" && fileContents[*i + 3] == "LPAREN" && fileContents[*i + 4][:3] == "NUM" && fileContents[*i + 5] == "COMMA" && fileContents[*i + 6][:3] == "NUM" && fileContents[*i + 7] == "RPAREN" {
		//advance the iterator 8 times for every token that was
		//analyzed with the long if statement
		*i += 8
		return true
	} else if fileContents[*i] == "TEST" { //this case is here so that an error isn't thrown if test is supposed to be analyzed instead
		return false
	} else {
		panic("syntax error: incorrect point declaration")
		return false
	}
}

// ParseTEST
// represents the fifth grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParseTEST(fileContents []string, i *int) bool {
	if fileContents[*i] == "TEST" && fileContents[*i + 1] == "LPAREN" {
		//two tokens are analyzed, so advance the iterator twice
		*i += 2
		if ParseOPTION(fileContents, i) {
			if fileContents[*i] == "COMMA" {
				//only one token is analyzed, so advance the iterator once
				*i++
				if ParsePOINT_LIST(fileContents, i) {
					if fileContents[*i] == "RPAREN" {
						*i++
						return true
					}
				}
			}
		}
	}
	panic("syntax error: incorrect test declaration")
	return false
}

// ParseOPTION
// represents the sixth grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParseOPTION(fileContents []string, i *int) bool {
	if fileContents[*i] == "TRIANGLE" || fileContents[*i] == "SQUARE" {
		//only one option here, so advance the iterator once
		*i++
		return true
	} else {
		panic("syntax error: incorrect option")
		return false
	}
}

// ParsePOINT_LIST
// represents the seventh grammar rule, uses the
// iterator to access the tokens that it needs for analysis
// returns whether the Syntax is correct so far
// true if it is, false if not
func ParsePOINT_LIST(fileContents []string, i *int) bool {
	if fileContents[*i][:2] == "ID" && fileContents[*i + 1] == "COMMA" {
		//advance the iterator twice because two tokens are analyzed
		*i += 2
		//call point list recursively to analyze the rest of the ID COMMA combos
		return ParsePOINT_LIST(fileContents, i)
	}
	if fileContents[*i][:2] == "ID" { //ID should be last in the chain of ID COMMAs
		*i++
		return true
	} else {
		panic("syntax error: wrong parameters")
		return false
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
	fmt.Println("Processing input file " + fileName)
	LexicallyAnalyze(fileName)

	//trim out .cpl from fileName
	tknFileName := strings.TrimRight(fileName, ".cpl")

	//do Syntax Analysis
	SyntacticallyAnalyze(tknFileName + ".tkn")

	//Decides whether to generate scheme or prolog code based on command line arg
	if os.Args[2] == "-s" {
		fmt.Println(GenerateScheme(tknFileName + ".tkn"))
	} else if os.Args[2] == "-p" {
		fmt.Println(GenerateProlog(tknFileName + ".tkn"))
	} else {
		panic("Include either \"-s\" or \"-p\" in the third command line argument")
	}
}
