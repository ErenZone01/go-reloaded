package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ToUpper(s string) string {
	letter := []rune(s)
	var x string = ""

	for i := 0; i < len(s); i++ {
		if letter[i] >= 'a' && letter[i] <= 'z' {
			letter[i] = letter[i] - 32
			x += string(letter[i])
		} else {
			x += string(letter[i])
		}
	}
	return x
}

func Capitalize(s string) string {
	tab := []byte(s)
	if len(s) > 0 {
		if s[0] != ' ' && (s[0] >= 'a' && s[0] <= 'z') {
			tab[0] = (tab[0] - ' ')
		}
		for i := 0; i < len(s)-1; i++ {
			if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
				if s[i+1] >= 'a' && s[i+1] <= 'z' {
					tab[i+1] = (tab[i+1] - ' ')
				}
			} else if s[i+1] >= 'A' && s[i+1] <= 'Z' {
				tab[i+1] = (tab[i+1] + ' ')
			}
		}
	}
	return string(tab)
}

func ToLower(s string) string {
	letter := []rune(s)
	var x string = ""
	for i := 0; i < len(s); i++ {
		if letter[i] >= 'A' && letter[i] <= 'Z' {
			letter[i] = letter[i] + 32
			x += string(letter[i])
		} else {
			x += string(letter[i])
		}
	}
	return x
}

func SplitWhiteSpaces(s string) []string {
	phrase := make([]string, 0)
	var texte string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " || string(s[i]) == "\n" || i == len(s)-1 {
			if i == len(s)-1 {
				texte += string(s[i])
				phrase = append(phrase, texte)
				break
			}
			if string(s[i+1]) != " " {
				phrase = append(phrase, texte)
				texte = ""
			}
			continue
		}
		texte += string(s[i])
	}
	return phrase
}

func IterativePower(nb int, power int) int {
	var result int = 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}

func hexa(nmbr []int) string {
	var resultat []int
	var hex int
	var phrase string

	for i := len(nmbr) - 1; i >= 0; i-- {
		resultat = append(resultat, nmbr[i])
	}

	for i := 0; i < len(nmbr); i++ {
		hex += resultat[i] * IterativePower(16, i)
	}
	phrase = strconv.Itoa(hex)
	return phrase
}

func bin(nmbr int) string {
	var nombre int = nmbr
	var reste int
	var result int

	for i := 0; nombre != 0; i++ {
		reste = nombre % 10
		nombre /= 10
		result += reste * IterativePower(2, i)
	}
	convert := strconv.Itoa(result)
	return convert
}

func IsNumeric(s string) bool {
	nombre := []rune(s)
	for i := 0; i < len(s); i++ {
		if !(nombre[i] >= '0' && nombre[i] <= '9') {
			return false
		}
	}
	return true
}

func foundWord(s string) bool {
	var etat bool = false
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			etat = true
		}
	}
	if etat == false {
		return false
	}

	return true
}

func main() {
	if len(os.Args[1:]) == 2 {
		arg := os.Args[1:]
		files, _ := ioutil.ReadFile(arg[0])
		texte := string(files)
		f, _ := os.Create(arg[1])
		var tab = SplitWhiteSpaces(texte)
		var nombre int
		var lettre rune
		var compare = []rune(texte)
		var hexadecimalMax = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
		var hexadecimalMin = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
		var decimal = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}

		if len(texte) != len(compare) {
			fmt.Println("Ce texte contient des caractères non ASCII")
			os.Exit(1)
		}
		var etat = true
		for i := 0; i < len(tab); i++ {
			if !etat {
				i = 0
				etat = true
			}

			if tab[i] == "(up)" {
				if i != 0 {
					for j := i - 1; j >= 0; j-- {
						if foundWord(tab[j]) {
							tab[j] = ToUpper(tab[j])
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}
					}
				} else {
					if len(tab) > 1 {
						tab = append(tab[:i], tab[i+1:]...)
						etat = false
					} else {
						tab[i] = ""
					}
				}

			}
			if tab[i] == "(cap)" {
				if i != 0 {
					for j := i - 1; j >= 0; j-- {
						if foundWord(tab[j]) {
							tab[j] = Capitalize(tab[j])
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}
					}
				} else {
					if len(tab) > 1 {
						tab = append(tab[:i], tab[i+1:]...)
						etat = false
					} else {
						tab[i] = ""
					}
				}

			}
			if tab[i] == "(low)" {
				if i != 0 {
					for j := i - 1; j >= 0; j-- {
						if foundWord(tab[j]) {
							tab[j] = ToLower(tab[j])
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}
					}
				} else {
					if len(tab) > 1 {
						tab = append(tab[:i], tab[i+1:]...)
						etat = false
					} else {
						tab[i] = ""
					}
				}

			}
			if tab[i] == "(hex)" {
				//var tabnmbr []int
				if i != 0 {
					var phrase []string
					var chiffre []int
					var color bool = true
					for j := 0; j < len(tab[i-1]); j++ {
						for k := 0; k < len(hexadecimalMax); k++ {
							if string(tab[i-1][j]) != hexadecimalMax[k] && string(tab[i-1][j]) != hexadecimalMin[k] {
								color = false
							} else {
								color = true
							}

							if string(tab[i-1][j]) == hexadecimalMax[k] {
								phrase = append(phrase, decimal[k])
								break
							}
							if string(tab[i-1][j]) == hexadecimalMin[k] {
								phrase = append(phrase, decimal[k])
								break
							}
						}
						if color != true {
							break
						}

					}
					if color {
						for k := 0; k < len(phrase); k++ {
							var total int
							total, _ = strconv.Atoi(phrase[k])
							chiffre = append(chiffre[:k], total)
						}

						tab[i-1] = hexa(chiffre)
						tab = append(tab[:i], tab[(i+1):]...)
						i = i - 1
					} else {
						tab[i] = ""
					}

				} else {
					if len(tab) > 1 {
						tab = append(tab[:i], tab[i+1:]...)
						etat = false
					} else {
						tab[i] = ""
					}
				}

			}
			if tab[i] == "(bin)" {
				if i != 0 {
					var chiffre = tab[i-1]
					if IsNumeric(chiffre) {
						total, _ := strconv.Atoi(chiffre)
						tab[i-1] = bin(total)
						tab = append(tab[:i], tab[(i+1):]...)
						i = i - 1
					} else {
						tab[i] = ""
					}

				} else {
					if len(tab) > 1 {
						tab = append(tab[:i], tab[i+1:]...)
						etat = false
					} else {
						tab[i] = ""
					}
				}

			}
			var nmbstring string
			if tab[i] == "(cap," {
				//if i < len(tab) {
				if i != 0 && i != len(tab)-1 {
					if tab[i+1][len(tab[i+1])-1] == ')' {
						for j := 0; j < len(tab[i+1]); j++ {
							if tab[i+1][j] == ')' {
								break
							} else {
								lettre = rune(tab[i+1][j])
								if tab[i+1][0] != '-' {
									nmbstring += string(lettre)
								}
								nombre = (nombre * 10) + int(lettre-'0')
							}

						}
					}

					if IsNumeric(nmbstring) {
						if tab[i+1][0] == '-' {
							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}
						var nbrWord int
						if tab[i+1][len(tab[i+1])-1] == ')' {
							for j := i - 1; j >= 0; j-- {
								if foundWord(tab[j]) {
									nbrWord++
								}

							}

							if nombre > nbrWord {
								for j := i - 1; nbrWord != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = Capitalize(tab[j])
										nbrWord--
									}

								}
							} else {
								for j := i - 1; nombre != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = Capitalize(tab[j])
										nombre--
									}
								}
							}

							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							nmbstring = ""
						}
					}
				} else {
					if len(tab) >= 2 {
						if tab[1][len(tab[1])-1] == ')' && len(tab[1]) > 1 {
							var strnmbr string
							for j := 0; j < len(tab[1])-1; j++ {
								if tab[i][0] == '-' {
									strnmbr += string(tab[1][j])
								}
							}
							if IsNumeric(strnmbr) {
								if tab[i+1][0] == '-' {
									tab = append(tab[:i], tab[(i+1):]...)
									tab = append(tab[:i], tab[(i+1):]...)
									i = i - 1
									break
								}
								if len(tab) == 2 {
									tab = append(tab[:i], tab[i+1:]...)
									tab[i] = ""
								} else {
									tab = append(tab[:i], tab[i+1:]...)
									tab = append(tab[:i], tab[i+1:]...)
								}

								etat = false
							}

						}
					}
				}
			}
			if tab[i] == "(up," {
				if i != 0 && i != len(tab)-1 {
					for j := 0; j < len(tab[i+1]); j++ {
						if tab[i+1][j] == ')' {
							break
						}
						lettre = rune(tab[i+1][j])
						if tab[i+1][0] != '-' {
							nmbstring += string(lettre)
						}

						nombre = (nombre * 10) + int(lettre-'0')
					}

					if IsNumeric(nmbstring) {
						var nbrWord int
						if tab[i+1][0] == '-' {
							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}

						if tab[i+1][len(tab[i+1])-1] == ')' {
							for j := i - 1; j >= 0; j-- {
								if foundWord(tab[j]) {
									nbrWord++
								}
							}

							if nombre > nbrWord {
								for j := i - 1; nbrWord != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = ToUpper(tab[j])
										nbrWord--
									}

								}
							} else {
								for j := i - 1; nombre != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = ToUpper(tab[j])
										nombre--
									}

								}
							}

							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
						}

					}
				} else {
					if len(tab) >= 2 {
						if tab[1][len(tab[1])-1] == ')' && len(tab[1]) > 1 {
							var strnmbr string
							for j := 0; j < len(tab[1])-1; j++ {
								if tab[i][0] == '-' {
									strnmbr += string(tab[1][j])
								}
							}
							if IsNumeric(strnmbr) {
								if tab[i+1][0] == '-' {
									tab = append(tab[:i], tab[(i+1):]...)
									tab = append(tab[:i], tab[(i+1):]...)
									i = i - 1
									break
								}
								if len(tab) == 2 {
									tab = append(tab[:i], tab[i+1:]...)
									tab[i] = ""
								} else {
									tab = append(tab[:i], tab[i+1:]...)
									tab = append(tab[:i], tab[i+1:]...)
								}

								etat = false
							}

						}
					}
				}
				nmbstring = ""
			}

			if tab[i] == "(low," {
				if i != 0 && i != len(tab)-1 {
					for j := 0; j < len(tab[i+1]); j++ {
						if tab[i+1][j] == ')' {
							break
						}
						lettre = rune(tab[i+1][j])
						if tab[i+1][0] != '-' {
							nmbstring += string(lettre)
						}
						nombre = (nombre * 10) + int(lettre-'0')
					}
					if IsNumeric(nmbstring) {
						if tab[i+1][0] == '-' {
							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
							break
						}
						var nbrWord int

						if tab[i+1][len(tab[i+1])-1] == ')' {
							for j := i - 1; j >= 0; j-- {
								if foundWord(tab[j]) {
									nbrWord++
								}

							}

							if nombre > nbrWord {

								for j := i - 1; nbrWord != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = ToLower(tab[j])
										nbrWord--
									}

								}
							} else {
								for j := i - 1; nombre != 0; j-- {
									if foundWord(tab[j]) {
										tab[j] = ToLower(tab[j])
										nombre--
									}

								}
							}
							tab = append(tab[:i], tab[(i+1):]...)
							tab = append(tab[:i], tab[(i+1):]...)
							i = i - 1
						}

					}
				} else {
					if len(tab) >= 2 {
						if tab[1][len(tab[1])-1] == ')' && len(tab[1]) > 1 {
							var strnmbr string
							for j := 0; j < len(tab[1])-1; j++ {
								if tab[i+1][0] != '-' {
									strnmbr += string(tab[1][j])
								}
							}
							if IsNumeric(strnmbr) {
								if tab[i+1][0] == '-' {
									tab = append(tab[:i], tab[(i+1):]...)
									tab = append(tab[:i], tab[(i+1):]...)
									i = i - 1
									break
								}
								if len(tab) == 2 {
									tab = append(tab[:i], tab[i+1:]...)
									tab[i] = ""
								} else {
									tab = append(tab[:i], tab[i+1:]...)
									tab = append(tab[:i], tab[i+1:]...)
								}

								etat = false
							}

						}
					}
				}

			}

		}
		var voyelle = []rune{'a', 'o', 'i', 'e', 'u', 'A', 'O', 'U', 'I', 'E', 'H', 'h'}
		var consonne = []rune{'b', 'c', 'd', 'f', 'g', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', 'B', 'C', 'D', 'F', 'G', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}

		for i := 0; i < len(tab); i++ {
			if len(tab) > 1 {
				if i+1 == len(tab) {
					break
				}
				if tab[i] == "a" || tab[i] == "A" || tab[i] == "an" || tab[i] == "An" {
					for j := 0; j < len(voyelle); j++ {
						if tab[i] == "a" && rune(tab[i+1][0]) == voyelle[j] {
							tab[i] = "an"
						} else if tab[i] == "A" && rune(tab[i+1][0]) == voyelle[j] {
							tab[i] = "An"
						} else if tab[i] == "An" && rune(tab[i+1][0]) == voyelle[j] {
							tab[i] = "An"
						} else if tab[i] == "an" && rune(tab[i+1][0]) == voyelle[j] {
							tab[i] = "an"
						}
					}
					for j := 0; j < len(consonne); j++ {
						if tab[i] == "a" && rune(tab[i+1][0]) == consonne[j] {
							tab[i] = "a"
						} else if tab[i] == "A" && rune(tab[i+1][0]) == consonne[j] {
							tab[i] = "A"
						} else if tab[i] == "An" && rune(tab[i+1][0]) == consonne[j] {
							tab[i] = "A"
						} else if tab[i] == "an" && rune(tab[i+1][0]) == consonne[j] {
							tab[i] = "a"
						}
					}
				}
			}

		}

		var phrase string
		for i := 0; i < len(tab); i++ {
			phrase += tab[i]
			if i+1 != len(tab) {
				phrase += " "
			}

		}

		var remplacant1 = []rune(phrase)
		var marqueur []int

		for i := 0; i < len(remplacant1); i++ {
			if i < len(remplacant1) {
				if i < len(remplacant1)-1 && remplacant1[i] == ' ' && (remplacant1[i+1] == ',' || remplacant1[i+1] == ';' || remplacant1[i+1] == ':' || remplacant1[i+1] == '!' || remplacant1[i+1] == '?' || remplacant1[i+1] == '.') {
					remplacant1 = append(remplacant1[:i], remplacant1[(i+1):]...)
					// tmp := remplacant1[i+1]
					// remplacant1[i+1] = remplacant1[i]
					// remplacant1[i] = tmp
					// i = i - 1
				}

				if (i == 0 && len(remplacant1) > 1) && (remplacant1[i] == ',' || remplacant1[i] == ';' || remplacant1[i] == ':' || remplacant1[i] == '!' || remplacant1[i] == '?' || remplacant1[i] == '.') && (remplacant1[i+1] != ',' && remplacant1[i+1] != ';' && remplacant1[i+1] != ':' && remplacant1[i+1] != '!' && remplacant1[i+1] != '?' && remplacant1[i+1] != '.') {
					marqueur = append(marqueur, i)
				}

				if i < len(remplacant1)-1 && i > 0 && (remplacant1[i] == ',' || remplacant1[i] == ';' || remplacant1[i] == ':' || remplacant1[i] == '!' || remplacant1[i] == '?' || remplacant1[i] == '.') && (remplacant1[i+1] != ' ' && remplacant1[i-1] != ' ') && (remplacant1[i+1] != ',' && remplacant1[i+1] != ';' && remplacant1[i+1] != ':' && remplacant1[i+1] != '!' && remplacant1[i+1] != '?' && remplacant1[i+1] != '.') {
					marqueur = append(marqueur, i)
				}

				// if i < len(remplacant1)-1 {
				// 	if remplacant1[i] == ' ' && remplacant1[i+1] == ' ' {
				// 		remplacant1 = append(remplacant1[:i], remplacant1[(i+1):]...)
				// 		i = i - 1
				// 	}
				// }

			}

		}
		var tabPhrase = remplacant1
		var compteur int
		var posx int
		var posy int

		phrase = ""
		for i := 0; i < len(remplacant1); i++ {
			phrase += string(remplacant1[i])
			for j := 0; j < len(marqueur); j++ {
				if marqueur[j] == i {
					phrase += " "
				}
			}
		}

		tabPhrase = []rune(phrase)

		for i := 0; i < len(tabPhrase); i++ {
			if tabPhrase[i] == '\'' {
				if i != 0 && i+1 != len(tabPhrase) && tabPhrase[i] == '\'' && tabPhrase[i+1] != ' ' && tabPhrase[i-1] != ' ' {
					continue
				} else {
					compteur++
					if compteur == 1 {
						posx = i
					}
					if compteur == 2 {
						posy = i
					}
					if compteur == 2 {
						if tabPhrase[posx+1] == ' ' {
							tmp := tabPhrase[posx]
							tabPhrase[posx] = tabPhrase[posx+1]
							tabPhrase[posx+1] = tmp

						}
						if tabPhrase[posy-1] == ' ' {
							tmp := tabPhrase[posy]
							tabPhrase[posy] = tabPhrase[posy-1]
							tabPhrase[posy-1] = tmp
						}
						compteur = 0
					}
				}

			}
		}

		phrase = ""
		for i := 0; i < len(tabPhrase); i++ {
			phrase += string(tabPhrase[i])
		}

		var remplacant = []rune(phrase)

		phrase = ""

		for i := 0; i < len(remplacant)-1; i++ {
			if remplacant[i] == ' ' && remplacant[i+1] == ' ' {
				remplacant = append(remplacant[:i], remplacant[(i+1):]...)
				i = i - 1
			}
		}

		for i := 0; i < len(remplacant); i++ {
			if remplacant[i] == ' ' && i+1 == len(remplacant) {
				break
			}

			if i == 0 && remplacant[0] == ' ' {
				remplacant = append(remplacant[:i], remplacant[i+1:]...)
			}
			phrase += string(remplacant[i])

		}

		f.WriteString(string(phrase))
	}

}
