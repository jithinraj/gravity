/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"bytes"
	"path/filepath"
	"time"

	"github.com/gravitational/gravity/lib/app"
	apptest "github.com/gravitational/gravity/lib/app/service/test"
	"github.com/gravitational/gravity/lib/blob/fs"
	"github.com/gravitational/gravity/lib/defaults"
	"github.com/gravitational/gravity/lib/loc"
	"github.com/gravitational/gravity/lib/pack"
	"github.com/gravitational/gravity/lib/pack/localpack"
	"github.com/gravitational/gravity/lib/storage/keyval"

	"github.com/gravitational/trace"
	. "gopkg.in/check.v1"
)

type PullerSuite struct {
	srcPack pack.PackageService
	dstPack pack.PackageService
	srcApp  app.Applications
	dstApp  app.Applications
}

var _ = Suite(&PullerSuite{})

func (s *PullerSuite) SetUpSuite(c *C) {
	s.srcPack, s.srcApp = setupServices(c)
	s.dstPack, s.dstApp = setupServices(c)
}

func (s *PullerSuite) TestPullPackage(c *C) {
	err := s.srcPack.UpsertRepository("example.com", time.Time{})
	c.Assert(err, IsNil)

	loc := loc.MustParseLocator("example.com/package:0.0.1")

	_, err = s.srcPack.CreatePackage(loc, bytes.NewBuffer([]byte("data")))
	c.Assert(err, IsNil)

	env, err := PullPackage(PackagePullRequest{
		SrcPack: s.srcPack,
		DstPack: s.dstPack,
		Package: loc,
	})
	c.Assert(err, IsNil)
	c.Assert(env.Locator, Equals, loc)

	env, err = s.dstPack.ReadPackageEnvelope(loc)
	c.Assert(err, IsNil)
	c.Assert(env.Locator, Equals, loc)

	_, err = PullPackage(PackagePullRequest{
		SrcPack: s.srcPack,
		DstPack: s.dstPack,
		Package: loc,
	})
	c.Assert(trace.IsAlreadyExists(err), Equals, true)
}

func (s *PullerSuite) TestPullApp(c *C) {
	err := s.srcPack.UpsertRepository("example.com", time.Time{})
	c.Assert(err, IsNil)

	runtimePackage := loc.MustParseLocator("gravitational.io/planet:0.0.1")
	apptest.CreatePackage(s.srcPack, runtimePackage, nil, c)
	apptest.CreateRuntimeApplication(s.srcApp, c)

	locator := loc.MustParseLocator("example.com/package:0.0.2")
	apptest.CreateDummyApplication(s.srcApp, locator, c)

	pulled, err := PullApp(AppPullRequest{
		SrcPack: s.srcPack,
		DstPack: s.dstPack,
		SrcApp:  s.srcApp,
		DstApp:  s.dstApp,
		Package: locator,
	})
	c.Assert(err, IsNil)
	c.Assert(pulled.Package, Equals, locator)

	local, err := s.dstApp.GetApp(locator)
	c.Assert(err, IsNil)
	c.Assert(local.Package, Equals, locator)

	_, err = PullApp(AppPullRequest{
		SrcPack: s.srcPack,
		DstPack: s.dstPack,
		SrcApp:  s.srcApp,
		DstApp:  s.dstApp,
		Package: locator,
	})
	c.Assert(trace.IsAlreadyExists(err), Equals, true)
}

func setupServices(c *C) (pack.PackageService, app.Applications) {
	dir := c.MkDir()

	backend, err := keyval.NewBolt(keyval.BoltConfig{
		Path: filepath.Join(dir, "bolt.db"),
	})
	c.Assert(err, IsNil)

	objects, err := fs.New(dir)
	c.Assert(err, IsNil)

	packService, err := localpack.New(localpack.Config{
		Backend:     backend,
		UnpackedDir: filepath.Join(dir, defaults.UnpackedDir),
		Objects:     objects,
	})
	c.Assert(err, IsNil)

	appService, err := New(Config{
		Backend:  backend,
		StateDir: filepath.Join(dir, defaults.ImportDir),
		Packages: packService,
	})
	c.Assert(err, IsNil)

	return packService, appService
}
