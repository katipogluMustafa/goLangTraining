package validationKit

import "testing"

func TestCheckUsernameSyntaxMinimumCharacterLength(t *testing.T) {
	result := CheckUsernameSyntax("")

	if result != false {
		t.Errorf("Failed the minimum character check.")
	}
}

func TestCheckUsernameSyntaxMaximumCharacterLength(t *testing.T) {
	result := CheckUsernameSyntax("aaaaaaaaaaaaaaaa")

	if result != false {
		t.Errorf("Failed the maximum character check.")
	}
}

func TestCheckUsernameSyntaxSymbols(t *testing.T) {
	result := CheckUsernameSyntax("g-@p!h#r")

	if result != false {
		t.Errorf("Failed the special character check.")
	}
}

func TestCheckUsernameSyntaxUnderscore(t *testing.T) {
	result := CheckUsernameSyntax("the_gopher")

	if result != true {
		t.Errorf("Failed the check to allow underscore characters.")
	}
}

func TestCheckUsernameSyntaxAtSignInsideUsername(t *testing.T) {
	result := CheckUsernameSyntax("the@gopher")

	if result != false {
		t.Errorf("Failed the @ sign check.The @ Sign was found in another place besides the start of the string.")
	}
}

func TestCheckUsernameSyntaxRandomUsernames(t *testing.T) {

	for i := 0; i < 10008; i++ {
		username := GenerateRandomUsername()
		result := CheckUsernameSyntax(username)
		if result != true {
			t.Error("The random username, ", username, ", failed to pass the username check.")
			t.Fatal("Quitting!")
		}
	}

}
