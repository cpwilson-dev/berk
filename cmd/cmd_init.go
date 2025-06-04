package cmd

/*
currentLimits:
- only creates an init file in cwd.
- will only run if there is no `.berk` directory

checks:
- verify that the directory exists

actions:
- creates a `.berk` directory
- inside directory creates the following
    - hooks/
    - info/
    - objects/
    - refs/
    - config
    - description
    - HEAD
*/

func commandInit() error {
	return nil
}
