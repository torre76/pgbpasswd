package files

import "github.com/torre76/pgbpasswd/types"

// HBAFileFileManager is a FileManager that writes Postgres HBA Auth File Style files
type HBAFileFileManager struct{}

func (fm *HBAFileFileManager) Write(fileName string, elements []types.LoginPassword) {

}
