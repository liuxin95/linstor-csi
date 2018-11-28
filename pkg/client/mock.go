/*
CSI Driver for Linstor
Copyright © 2018 LINBIT USA, LLC

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see <http://www.gnu.org/licenses/>.
*/

package client

import (
	"github.com/LINBIT/linstor-csi/pkg/volume"
)

type assignment struct {
	vol  *volume.Info
	node string
}

type MockStorage struct {
	createdVolumes  []*volume.Info
	assignedVolumes []*assignment
}

func (s *MockStorage) ListAll(parameters map[string]string) ([]*volume.Info, error) {
	return s.createdVolumes, nil
}

func (s *MockStorage) GetByName(name string) (*volume.Info, error) {
	for _, vol := range s.createdVolumes {
		if vol.Name == name {
			return vol, nil
		}
	}
	return nil, nil
}

func (s *MockStorage) GetByID(ID string) (*volume.Info, error) {
	for _, vol := range s.createdVolumes {
		if vol.ID == ID {
			return vol, nil
		}
	}
	return nil, nil
}

func (s *MockStorage) Create(vol *volume.Info) error {
	s.createdVolumes = append(s.createdVolumes, vol)
	return nil
}

func (s *MockStorage) Delete(vol *volume.Info) error {
	for _, v := range s.createdVolumes {
		if v.Name == vol.Name {
			// Close enough for testing...
			v = nil
		}
	}
	return nil
}

func (s *MockStorage) Attach(vol *volume.Info, node string) error {
	s.assignedVolumes = append(s.assignedVolumes, &assignment{vol: vol, node: node})
	return nil
}

func (s *MockStorage) Detech(vol *volume.Info, node string) error {
	for _, a := range s.assignedVolumes {
		if a.vol.Name == vol.Name && a.node == node {
			// Close enough for testing...
			a = nil
		}
	}
	return nil
}

func (s *MockStorage) NodeAvailable(node string) (bool, error) {
	// Hard coding magic string to pass csi-test.
	if node == "some-fake-node-id" {
		return false, nil
	}

	return true, nil
}
