package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type RemoteCallbackInterface interface {
	HTTPSAuthCallBack() git2go.CredentialsCallback
	SSHAUthCallBack() git2go.CredentialsCallback
	NoAuthCallBack() git2go.CredentialsCallback
	CertCallback() git2go.CertificateCheckCallback
}

type RemoteCallbackStruct struct {
	RepoName   string
	UserName   string
	Password   string
	SSHKeyPath string
}

func (grc *RemoteCallbackStruct) CertCallback() git2go.CertificateCheckCallback {
	logger.Log("Verification to allow remote host certificates", global.StatusInfo)
	return func(cert *git2go.Certificate, valid bool, hostname string) git2go.ErrorCode {
		return 0
	}
}

func (grc *RemoteCallbackStruct) SSHAUthCallBack() git2go.CredentialsCallback {
	logger.Log("Initiating SSH credential creation", global.StatusInfo)
	return func(url string, username_from_url string, allowed_types git2go.CredentialType) (*git2go.Credential, error) {
		if strings.Contains(allowed_types.String(), "SSH") {
			if runtime.GOOS != "windows" {
				_, agentErr := exec.LookPath("ssh-agent")
				if agentErr != nil {
					logger.Log(agentErr.Error(), global.StatusError)
					return nil, agentErr
				}

				_, sshAddErr := exec.LookPath("ssh-agent")
				if sshAddErr != nil {
					logger.Log(sshAddErr.Error(), global.StatusError)
					return nil, sshAddErr
				}

				sshAdd, _ := exec.LookPath("ssh-add")
				cmd := exec.Cmd{
					Path: sshAdd,
					Args: []string{"", grc.SSHKeyPath},
				}
				_, err := cmd.Output()
				if err != nil {
					logger.Log("Unable to add SSH key : "+err.Error(), global.StatusError)
					return nil, err
				}
				return git2go.NewCredentialSSHKeyFromAgent(username_from_url)
			}

			logger.Log("Starting up pageant agent for windows", global.StatusInfo)
			execPath, execPathErr := os.Executable()

			if execPathErr != nil {
				logger.Log(execPathErr.Error(), global.StatusError)
				return nil, execPathErr
			}

			execPath = filepath.Dir(execPath)
			etcPath := execPath + "\\etc\\"
			keyGenPath := etcPath + "\\" + global.PuttyGenExeName
			ppkFileName := etcPath + "\\" + grc.RepoName + ".ppk"

			cmdArgs := []string{"", grc.SSHKeyPath, "-o", ppkFileName}
			cmd := exec.Cmd{
				Path: keyGenPath,
				Args: cmdArgs,
			}
			_, cmdErr := cmd.Output()

			if cmdErr == nil {
				pageantPath := etcPath + global.PageantExeName
				cmdArgs := []string{"", ppkFileName}

				cmd := exec.Cmd{Path: pageantPath, Args: cmdArgs}
				_, cmdErr := cmd.Output()

				if cmdErr == nil {
					return git2go.NewCredentialSSHKeyFromAgent(username_from_url)
				} else {
					logger.Log(cmdErr.Error(), global.StatusError)
					return nil, types.Error{Msg: "AUTH Failed on a non-windows platform"}
				}
			} else {
				logger.Log("Unable to convert private key", global.StatusWarning)
				logger.Log(cmdErr.Error(), global.StatusError)
			}
		} else {
			logger.Log("Expected auth type is not received", global.StatusWarning)
			logger.Log(fmt.Sprintf("Expected : %s || Recived : %s", git2go.CredentialTypeSSHKey.String(), allowed_types.String()), global.StatusWarning)
			return nil, types.Error{Msg: "AUTH Failed due to incompatible auth mode " + allowed_types.String()}
		}
		return nil, types.Error{Msg: "AUTH Failed for SSH auth mode"}
	}
}

func (grc *RemoteCallbackStruct) HTTPSAuthCallBack() git2go.CredentialsCallback {
	return func(url string, username_from_url string, allowed_types git2go.CredentialType) (*git2go.Credential, error) {
		if allowed_types == git2go.CredentialTypeUserpassPlaintext {
			return git2go.NewCredentialUserpassPlaintext(grc.UserName, grc.Password)
		} else {
			return nil, types.Error{Msg: "HTTPS Basic auth failed"}
		}
	}
}

func (grc *RemoteCallbackStruct) NoAuthCallBack() git2go.CredentialsCallback {
	return func(url string, username_from_url string, allowed_types git2go.CredentialType) (*git2go.Credential, error) {
		return git2go.NewCredentialDefault()
	}
}
