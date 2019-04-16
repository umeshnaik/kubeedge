/*
Copyright 2019 The Kubeedge Authors.

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
package options

const (
	// CfgPath flag sets the path to kubeadm config file.
	CfgPath = "config"
	//kubeedgeVersion sets the path of kubeedgeVersion

	KubeedgeVersion = "kubeedge-version"
	//DockerVersion sets the path of dockerVersion
	DockerVersion = "docker-version"
	//DockerVersion sets the path of kubeedgeversion
	Kubernetesversion = "kubernetes-version"
	// CertificateKey flag sets the key used to encrypt and decrypt certificate secrets
	CertificateKey = "certificate-key"
	// IgnorePreflightErrors sets the path a list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.
	IgnorePreflightErrors = "ignore-preflight-errors"
	//sets the path of the server address
	Server = "server"
	//sets the path of the Certpath
	CertPath = "certPath"
)
