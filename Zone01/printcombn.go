package pescine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i < '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	if n == 2 {
		for a := '0'; a <= '8'; a++ {
			for b := a + 1; b <= '9'; b++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				if a < '8' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	if n == 3 {
		for i := '0'; i <= '7'; i++ {
			for m := i + 1; m <= '8'; m++ {
				for a := m + 1; a <= '9'; a++ {
					z01.PrintRune(i)
					z01.PrintRune(m)
					z01.PrintRune(a)
					if i < '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	if n == 4 {
		for a := '0'; a <= '6'; a++ {
			for b := a + 1; b <= '7'; b++ {
				for c := b + 1; c <= '8'; c++ {
					for d := c + 1; d <= '9'; d++ {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if a < '6' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	if n == 5 {
		for a := '0'; a <= '5'; a++ {
			for b := a + 1; b <= '6'; b++ {
				for c := b + 1; c <= '7'; c++ {
					for d := c + 1; d <= '8'; d++ {
						for e := d + 1; e <= '9'; e++ {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(e)
							if a < '5' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	if n == 6 {
		for a := '0'; a <= '4'; a++ {
			for b := a + 1; b <= '5'; b++ {
				for c := b + 1; c <= '6'; c++ {
					for d := c + 1; d <= '7'; d++ {
						for e := d + 1; e <= '8'; e++ {
							for f := e + 1; f <= '9'; f++ {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune(f)
								if a < '4' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 7 {
		for a := '0'; a <= '3'; a++ {
			for b := a + 1; b <= '4'; b++ {
				for c := b + 1; c <= '5'; c++ {
					for d := c + 1; d <= '6'; d++ {
						for e := d + 1; e <= '7'; e++ {
							for f := e + 1; f <= '8'; f++ {
								for j := f + 1; j <= '9'; j++ {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(j)
									if a < '3' {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 8 {
		for a := '0'; a <= '2'; a++ {
			for b := a + 1; b <= '3'; b++ {
				for c := b + 1; c <= '4'; c++ {
					for d := c + 1; d <= '5'; d++ {
						for e := d + 1; e <= '6'; e++ {
							for f := e + 1; f <= '7'; f++ {
								for j := f + 1; j <= '8'; j++ {
									for h := j + 1; h <= '9'; h++ {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(j)
										z01.PrintRune(h)
										if a < '2' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 9 {
		for a := '0'; a <= '1'; a++ {
			for b := a + 1; b <= '2'; b++ {
				for c := b + 1; c <= '3'; c++ {
					for d := c + 1; d <= '4'; d++ {
						for e := d + 1; e <= '5'; e++ {
							for f := e + 1; f <= '6'; f++ {
								for j := f + 1; j <= '7'; j++ {
									for h := j + 1; h <= '8'; h++ {
										for i := h + 1; i <= '9'; i++ {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(j)
											z01.PrintRune(h)
											z01.PrintRune(i)
											if a < '1' {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
