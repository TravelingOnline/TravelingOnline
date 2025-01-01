package clients

type IBankClient interface {
	CreateWallet(string) error
}
