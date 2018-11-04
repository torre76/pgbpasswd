package files

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
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
	var buffer = elements
	var sourceFile *os.File

	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].Login < buffer[j].Login
	})

	if fm.fileExists(fileName) {
		// File Exists, lets try to load the content
		buffer, err := fm.Read(fileName)
		if err != nil {
			return err
		}

		// Check if the login already exists
		for _, lpOrigin := range buffer {
			for _, lpNew := range elements {
				if lpOrigin.Login == lpNew.Login {
					return errors.New("Login already present")
				}
			}
		}

		// buffer is valid, merge contains and sort it
		buffer = append(buffer, elements...)
		sort.Slice(buffer, func(i, j int) bool {
			return buffer[i].Login < buffer[j].Login
		})

		// Create a temporary file and copy to it
		tmpFile, err := ioutil.TempFile("", "pgbpasswd")
		if err != nil {
			return err
		}

		// Ensure temporary file will be removed
		defer os.Remove(tmpFile.Name())

		sourceFile, err := os.Open(fileName)
		if err != nil {
			return err
		}

		// Ensure source file will close even in case of error
		defer sourceFile.Close()

		// Create a temporary backup copy
		if _, err := io.Copy(tmpFile, sourceFile); err != nil {
			return err
		}

		// Removing source file to write the new one
		sourceFile.Close()
		fm.removeFile(fileName)

	}

	sourceFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// Ensure file will be closed
	defer sourceFile.Close()

	for _, lp := range buffer {
		line := fmt.Sprintf(`%q %q`, lp.Login, lp.HashedPassword)
		_, err := sourceFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// NewAuthFileFileManager create an instance of File Manager that deals with with PgBouncer Auth File format
func NewAuthFileFileManager() FileManager {
	return &authFileFileManager{}
}
