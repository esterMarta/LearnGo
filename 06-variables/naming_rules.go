package variables

// VARIABLE NAMING RULES
func NamingRules() {
	// CORRECT DECLARATIONS
	var speed int
	var SpeedD int

	// underscore is allowed but not recomended
	var _speed int

	// Unicode letters are allowed
	var 速度 int

	// Keep the compiler happy
	_, _, _, _ = speed, SpeedD, _speed, 速度
}
