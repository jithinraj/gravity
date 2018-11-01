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

package suite

import (
	"context"
	"io"
	"time"

	"github.com/gravitational/gravity/lib/ops"
	"github.com/gravitational/gravity/lib/storage"
	"github.com/gravitational/license/authority"

	teleauth "github.com/gravitational/teleport/lib/auth"
	teleclient "github.com/gravitational/teleport/lib/client"
	"github.com/gravitational/teleport/lib/services"

	"github.com/gravitational/trace"
)

type TestOpsProxy struct {
	O ops.Operator
}

func (p *TestOpsProxy) GetService(link storage.OpsCenterLink) (ops.Operator, error) {
	return p.O, nil
}

type TestProxy struct {
}

func (t *TestProxy) GetClient() teleauth.ClientI {
	return nil
}

func (t *TestProxy) ReverseTunnelAddr() string {
	return ""
}

// CertAuthorities returns a list of certificate
// authorities proxy wants remote teleport sites to trust
func (t *TestProxy) CertAuthorities(withPrivateKey bool) ([]services.CertAuthority, error) {
	return TestAuthorities, nil
}

// DeleteAuthority deletes teleport authorities for the provided domain name
func (t *TestProxy) DeleteAuthority(domainName string) error {
	return nil
}

// StartCertAuthority sets up trust for certificate authority
func (t *TestProxy) TrustCertAuthority(services.CertAuthority) error {
	return nil
}

func (t *TestProxy) GetProxyClient(ctx context.Context, siteName string, labels map[string]string) (*teleclient.ProxyClient, error) {
	return nil, trace.Errorf("not implemented")
}

func (t *TestProxy) GenerateUserCert(pub []byte, user string, ttl time.Duration) ([]byte, error) {
	return nil, nil
}

func (t *TestProxy) GetCertAuthorities(caType services.CertAuthType) ([]services.CertAuthority, error) {
	return nil, nil
}

func (t *TestProxy) GetCertAuthority(id services.CertAuthID, withPrivateKeys bool) (*authority.TLSKeyPair, error) {
	return nil, nil
}

func (t *TestProxy) GetLocalAuthorityDomain() string {
	return ""
}

// UpsertUser upserts user entry with teleport
func (t *TestProxy) UpsertUser(username, password string) (string, error) {
	return "", nil
}

// GetServers returns a list of servers matching particular label key value
// pair expression and returns a list of servers
// domainName is a site domain name
func (t *TestProxy) GetServers(ctx context.Context, domainName string, labels map[string]string) ([]services.Server, error) {
	return nil, nil
}

// GetServersCount returns a number of servers belonging to a particular site
func (t *TestProxy) GetServerCount(ctx context.Context, domainName string) (int, error) {
	return 0, nil
}

// ExecuteCommand executes a command on a remote node addrress
// for a given site domain
func (t *TestProxy) ExecuteCommand(ctx context.Context, domainName, nodeAddr, command string, out io.Writer) error {
	panic("not implemented")
}

func (t *TestProxy) GetPlanetLeaderIP() string {
	panic("")
}

var TestUserAuthority = services.NewCertAuthority(services.UserCA, "example.com", [][]byte{PubKey}, nil, nil)

var TestAuthorities = []services.CertAuthority{
	TestUserAuthority,
}

var PrivKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAvJGHcmQNWUjY2eKasmw171qZR0B5FOnzy/nAGB1JAE+QokFe
Bjo8Gkk3L2TSuVNn0NI5uo5Jwp7GYtbfSbowo11E922Bwp0sFoVzeeUMyLud9EPz
Hl8+VvE8WEa1lC4D4aqravAfTeeePrONIYoBttX5oYXQ7aZkM8N7yS7KWNOZpy9f
n1vkSCpDOK29edLHWVyiDcXzULxEbXhPFl9Ly9shuEbqic2LRggxBnh3fhy53u8X
5qj8bp+21GGsQJaZYZtc9ieNYamo/KQcA0hFfUgTmV74ehY0vZ7yQk+2dW22cFqw
Dv+xNmnNHlfuYhHNCfk8rnztxfbqHfifgCArQQIDAQABAoIBADhq8jNva+8CtJ68
BbzMU3bBjIqc550yQhcNKkQMvwKwy31AQXlrgv/6V+B+Me3w3mbD/zGp0LfB+Wkp
ELVmV5cJGNFOmjw3+jDizKHzvddxCtlCW0MDDAvHMV7YCQvEmLSz84WTQkp0ugvY
fKlEOS8S5hVFjDUOS3yRSD/xF+lrIlYUaR4gXnDAJZx9ttgfZlHOp8ehxk+1bn59
3Fv1fCXcCKmKUlTk1kFasD8P+2M3MKP42Ih5ap9cfLSVPiBS/6JRBxIlZrHM9/2a
w6vEp+qMwwgCmxLPMwZfem6LNHO/huTrWKf4ltVubb5bUXIe22udKp2WK4NWc3Ka
uG8EleECgYEA4A9Mwd0QJs0j1kpuJDNIjfFx6IROv3QAb0QPq0+192ZF8P9AEj8B
TNDQVzb/skM+2NDdvhZ5v4+OJQcUNpEskhX+5ikk8QHGAUY6vT8rO6oiIRMaxLuJ
OEDc2Qms1OmctTmgSVyaxfXIK2/GDdvOizt0Z7Y7abza4bigEm49hyMCgYEA13MI
H429Ua0tnVVmGJ/4OjnKbgtF7i02r50vDVktPruKWNy1bhRkRyaOoCH7Zt9WXF2j
GapZZN1N/clO4vf9gikH0VCo4Tc2JR635dXdfISlt8NLXmR800Ms1UCAKlwIOQjz
dgHcvEbvFwSe1MFgOJVGL82G2rUA/zDVOKdjXEsCgYAZxyjZlQlqrWdWHDIX0B6k
1gZ47d/xfvMd2gLDfuQ8lnOtinBgqQcJQ2z028sHQ11TrJQWbpeLRoTgFbRposIx
/H3bFRi+8alKND5Fz6K1tpk+nOgTglADPNMr1UUhKc9xujOKvTDBXcmt1ao/pe5Z
bnmyBPFI9QVpusgP1scVaQKBgE5mJYaV5VZbVkXyVXyQeZt2fBsfLwtEmKm+4OhS
kwxI4kcDyWGNOhBKD4xl0T3V928VA8zLGEyD22WGY5Zj93PtylJ4r3uEw8cuLm0M
LdSp0EPWZQ6sMmAOCbpwBjNj2fonL7C5bMF2bnpJzCJPW9w7NZcfivr68qnp8yzy
fE2RAoGBALWvlHVH/29KOVmM52sOk49tcyc3czjs/YANvbokiItxOB8VPY6QQQnS
/CBsCZxUuWegYmkUnstHDmY1LYqjxW4goOqizIksaReivPmsTuQ1qd+aqXTfg2pt
uy6c6X17xkP5q2Lq4i90ikyWm3Oc25aUEw48pRyK/6rABRUzpDLB
-----END RSA PRIVATE KEY-----`)

var PubKey = []byte(`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC8kYdyZA1ZSNjZ4pqybDXvWplHQHkU6fPL+cAYHUkAT5CiQV4GOjwaSTcvZNK5U2fQ0jm6jknCnsZi1t9JujCjXUT3bYHCnSwWhXN55QzIu530Q/MeXz5W8TxYRrWULgPhqqtq8B9N554+s40higG21fmhhdDtpmQzw3vJLspY05mnL1+fW+RIKkM4rb150sdZXKINxfNQvERteE8WX0vL2yG4RuqJzYtGCDEGeHd+HLne7xfmqPxun7bUYaxAlplhm1z2J41hqaj8pBwDSEV9SBOZXvh6FjS9nvJCT7Z1bbZwWrAO/7E2ac0eV+5iEc0J+TyufO3F9uod+J+AICtB`)
