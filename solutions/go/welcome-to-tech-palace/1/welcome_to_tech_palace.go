package techpalace
import "strings"
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := ""
	for i := 1; i <= numStarsPerLine; i++{
		starLine += "*"
	}
	return starLine + "\n" + welcomeMsg + "\n" + starLine
}

func CleanupMessage(oldMsg string) string {
	star := '*'
	var result strings.Builder 
	for _, c := range oldMsg {
		if c != star {
			result.WriteRune(c)
		}
	}
	return strings.TrimSpace(result.String())
}
