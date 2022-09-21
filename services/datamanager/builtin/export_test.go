// export_test.go adds functionality to the builtin package that we only want to use and expose during testing.
package builtin

import (
	"go.viam.com/rdk/services/datamanager/datasync"
)

// SetSyncerConstructor sets the syncer constructor for the data manager to use when creating its syncer.
func (svc *builtIn) SetSyncerConstructor(fn datasync.ManagerConstructor) {
	svc.syncerConstructor = fn
}

// SetWaitAfterLastModifiedSecs sets the wait time for the syncer to use when initialized/changed in Service.Update.
func (svc *builtIn) SetWaitAfterLastModifiedSecs(s int) {
	svc.waitAfterLastModifiedSecs = s
}

// Make getServiceConfig global for tests.
var GetServiceConfig = getServiceConfig

// Make getDurationFromHz global for tests.
var GetDurationFromHz = getDurationFromHz

func (svc *builtIn) GetSessionID() string {
	return *svc.getSessionID()
}
