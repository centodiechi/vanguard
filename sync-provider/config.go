package syncprovider

import "github.com/centodiechi/identity/client"

type SyncStore struct {
	IdtClient *client.IdentityClient
}

var ServiceProvider *SyncStore

func Initialize() error {
	identityClient, err := client.GetIdtClient()
	if err != nil {
		return err
	}
	ServiceProvider.IdtClient = identityClient
	return nil
}
