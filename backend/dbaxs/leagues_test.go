package dbaxs

import "testing"

// Not developed yet
func TestQuickSort(t *testing.T) {
	setTestServerInfo()
	ConnDB(server.username, server.password, server.ip, server.port, server.db_name)

}
