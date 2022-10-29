package vfs

import "go.nandlabs.io/commons/vfs/local"

type Manager struct {
	fileSystems map[string]VFileSystem
}

func (d *Manager) IsSupported(scheme string) (supported bool) {
	_, supported = d.fileSystems[scheme]
	return
}

func (d *Manager) Register(fs VFileSystem) {
	for _, s := range fs.Schemes() {
		d.fileSystems[s] = fs

	}
}

func Get() *Manager {
	m := &Manager{fileSystems: make(map[string]VFileSystem)}
	localFs := &local.OsFs{}
	m.Register(localFs)
	return m

}
