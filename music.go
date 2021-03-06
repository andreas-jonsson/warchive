// +build midi

/*
Copyright (C) 2016-2021 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package warchive

import (
	"path"

	"github.com/andreas-jonsson/warchive/xmi"
)

type Music map[string][]byte

func LoadMusic(arch *Archive) (Music, error) {
	mus := make(Music)
	for file, data := range arch.Files {
		if path.Ext(file) == ".XMI" {
			mid, err := xmi.ToMidi(data, xmi.NoConversion)
			if err != nil {
				return nil, err
			}
			mus[file] = mid
		}
	}
	return mus, nil
}
