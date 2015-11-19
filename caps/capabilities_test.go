// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2015 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package caps

import (
	"testing"

	. "gopkg.in/check.v1"

	"github.com/ubuntu-core/snappy/testutil"
)

func Test(t *testing.T) {
	TestingT(t)
}

type CapabilitySuite struct{}

var _ = Suite(&CapabilitySuite{})

func (s *CapabilitySuite) TestValidateName(c *C) {
	c.Assert(ValidateName("name with space"), ErrorMatches,
		`"name with space" is not a valid snap name`)
	c.Assert(ValidateName("name-with-trailing-dash-"), ErrorMatches,
		`"name-with-trailing-dash-" is not a valid snap name`)
	c.Assert(ValidateName("name-with-3-dashes"), IsNil)
	c.Assert(ValidateName("name"), IsNil)
}

func (s *CapabilitySuite) TestLoadBuiltInTypes(c *C) {
	repo := NewRepository()
	err := LoadBuiltInTypes(repo)
	c.Assert(err, IsNil)
	c.Assert(repo.types, testutil.Contains, FileType)
	c.Assert(repo.types, HasLen, 1) // Update this whenever new built-in type is added
	err = LoadBuiltInTypes(repo)
	c.Assert(err, ErrorMatches, `cannot add type "file": name already exists`)
}

func (s *CapabilitySuite) TestString(c *C) {
	cap := &Capability{Name: "name", Label: "label", Type: FileType}
	c.Assert(cap.String(), Equals, "name")
}
