package account

import "testing"

func TestAccountCreation(t *testing.T) {
	_, _, b, err := GenerateAccount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
}

func TestAccountImport(t *testing.T) {
	b, err := GenerateAccountPEM()
	if err != nil {
		t.Fatal(err)
	}

	cert, pvtKey, err := ImportAccount(b)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cert.Subject.Organization)
	t.Log(pvtKey)

}
