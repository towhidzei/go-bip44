package bitcoin_address

type ChangeType uint32

const (
	ExternalChangeType ChangeType = 0
	ChangeChangeType   ChangeType = 1
)

const HardenedKeyZeroIndex = 0x80000000

type Purpose uint32

const (
	BIP44Purpose Purpose = 44
	BIP49Purpose Purpose = 49
	BIP84Purpose Purpose = 84
)

type CoinType uint32

const (
	BitcoinCoinType CoinType = 0
	TestnetCoinType CoinType = 1
)
