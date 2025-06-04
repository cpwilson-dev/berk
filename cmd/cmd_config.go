package cmd

/*
currentLimits:
- only checks the local config
- `--remove-section` not included

checks:
- .berk/config file exists
- has flag. If not; display help

actions:
- `--add` => adds a key/value. Default local config
- `--get` => gets a key. Default global & local
- `--unset` => deletes a key. Default global & local
- `--list` => returns all keys. Default global & local

*/

func commandConfig() error {
	return nil
}
