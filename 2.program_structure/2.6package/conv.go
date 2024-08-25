package tempconv

func C2F(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func F2C(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func K2C(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func C2K(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
