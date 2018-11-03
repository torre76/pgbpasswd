package files

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/torre76/pgbpasswd/types"
)

// authFileFileManager is a FileManager that writes PGBouncer Auth File Style files
type authFileFileManager struct {
	baseFileManager
}

func (fm *authFileFileManager) Read(fileName string) ([]types.LoginPassword, error) {
	if !fm.fileExists(fileName) {
		return nil, errors.New("Cannot read from file, does it exists?")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	// Ensure File will be closed at the end
	defer file.Close()

	result := []types.LoginPassword{}
	fileScanner := bufio.NewScanner(file)

	/* Complex RegExp to find string like
	 * "login" "password"
	 * even if
	 * * "login" has been escaped
	 * * There are any kind of characters after second quoting
	 * Logic has been adapted starting from:
	 * https://stackoverflow.com/questions/249791/regex-for-quoted-string-with-escaping-quotes
	 */
	re := regexp.MustCompile(`^-*"([^"\\]*(\\.[^"\\]*)*)"?.*?"(.*?)".*$`)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := re.FindStringSubmatch(line)

		if len(matches) < 4 {
			return nil, errors.New("File format is not valid")
		}

		result = append(result, *types.NewLoginHashedPassword(
			strings.Replace(matches[1], `\"`, `"`, -1),  // Replace escape if login has been escaped
			strings.Replace(matches[3], `\"`, `"`, -1))) // Replace escape if hashed password has been escaped
	}

	return result, nil
}

func (fm *authFileFileManager) Write(fileName string, elements []types.LoginPassword) error {
	return nil
}

// NewAuthFileFileManager create an instance of File Manager that deals with with PgBouncer Auth File format
func NewAuthFileFileManager() FileManager {
	return &authFileFileManager{}
}
