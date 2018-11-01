package files

import "github.com/torre76/pgbpasswd/types"

// FileManager writes a collection of Login and Hashed Password to a file
// Since there are at least two different formats there will be multiple implementations
type FileManager interface {
	Write(fileName string, elements []types.LoginPassword)
}
