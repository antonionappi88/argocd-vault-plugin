// Copyright (C) 2022, CERN
// This software is distributed under the terms of the GNU General Public
// Licence version 3 (GPL Version 3), copied verbatim in the file "COPYING".
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as Intergovernmental Organization
// or submit itself to any jurisdiction.
//
// Authors: Antonio Nappi
package backends

import (
	"encoding/json"
	"io"
	"os/exec"
)

// Empty structure
type TeigiSecretmanager struct {
	Port     string
	Hostname string
	Username string
	Password string
}

type TeigiSecret struct {
	Secret          string `json:"secret"`
	Service         string `json:"service"`
}

// NewTeigiSecretmanagerBackend initializes a new 1Password Connect backend
func NewTeigiSecretmanagerBackend(hostname string, port string, username string, password string) *TeigiSecretmanager {
	return &TeigiSecretmanager{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
	}
}

// Login does kinit
func (a *TeigiSecretmanager) Login() error {

	c1 := exec.Command("echo", a.Password)
	c2 := exec.Command("kinit", a.Username)

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	c1.Start()
	c2.Start()

	c1.Wait()
	w.Close()
	c2.Wait()

	return nil

}

// NOT USED
// GetSecrets gets secrets from 1Password Connect server and returns the formatted data
func (a *TeigiSecretmanager) GetSecrets(service string, version string, annotations map[string]string) (map[string]interface{}, error) {
	// Not used. We don't want to get all secret of teigi

	data := make(map[string]interface{})
	return data, nil
}

// GetIndividualSecret will get the specific secret (placeholder) from the 1Password connect backend
// For 1Password, we only support placeholders replaced from the k/v pairs of a secret which cannot be individually addressed
// So, we use GetSecrets and extract the specific placeholder we want
func (a *TeigiSecretmanager) GetIndividualSecret(service, secret, version string, annotations map[string]string) (interface{}, error) {

	out, err := exec.Command("/usr/bin/tbag", "--tbag-hostname", a.Hostname, "--tbag-port", a.Port, "show", "--service", service, secret).CombinedOutput()
	if err != nil {
		return nil, err
	}

	var result TeigiSecret
	json.Unmarshal(out, &result)
	data := make(map[string]interface{})
	data[secret] = result.Secret
	return data[secret], nil
}
