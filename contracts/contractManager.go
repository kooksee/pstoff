package contracts

var eventNameHashMap map[string]*Contract

func GetContract(name string) *Contract {
	return eventNameHashMap[name]
}

func initContracts() error {
	eventNameHashMap = make(map[string]*Contract)
	for _, c := range cfg.Contracts {
		nContract, err := NewContract(c.Address, c.Abi)
		if err != nil {
			return err
		}
		eventNameHashMap[c.Name] = nContract
	}
	return nil
}
