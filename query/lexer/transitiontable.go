
package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 40 : // ['(','(']
				return 2
			case r == 41 : // [')',')']
				return 3
			case r == 44 : // [',',',']
				return 4
			case r == 45 : // ['-','-']
				return 5
			case r == 47 : // ['/','/']
				return 6
			case r == 48 : // ['0','0']
				return 7
			case 49 <= r && r <= 57 : // ['1','9']
				return 8
			case r == 58 : // [':',':']
				return 9
			case r == 61 : // ['=','=']
				return 10
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 91 : // ['[','[']
				return 12
			case r == 93 : // [']',']']
				return 13
			case r == 97 : // ['a','a']
				return 11
			case r == 98 : // ['b','b']
				return 14
			case r == 99 : // ['c','c']
				return 11
			case r == 100 : // ['d','d']
				return 15
			case 101 <= r && r <= 107 : // ['e','k']
				return 11
			case r == 108 : // ['l','l']
				return 16
			case r == 109 : // ['m','m']
				return 17
			case 110 <= r && r <= 111 : // ['n','o']
				return 11
			case r == 112 : // ['p','p']
				return 18
			case r == 113 : // ['q','q']
				return 11
			case r == 114 : // ['r','r']
				return 19
			case r == 115 : // ['s','s']
				return 20
			case 116 <= r && r <= 122 : // ['t','z']
				return 11
			case r == 123 : // ['{','{']
				return 21
			case r == 125 : // ['}','}']
				return 22
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S4
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 24
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 26
			case 48 <= r && r <= 57 : // ['0','9']
				return 27
			case 65 <= r && r <= 90 : // ['A','Z']
				return 28
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 8
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 120 : // ['a','x']
				return 11
			case r == 121 : // ['y','y']
				return 29
			case r == 122 : // ['z','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 30
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 31
			case 98 <= r && r <= 122 : // ['b','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 32
			case 98 <= r && r <= 100 : // ['b','d']
				return 11
			case r == 101 : // ['e','e']
				return 33
			case 102 <= r && r <= 104 : // ['f','h']
				return 11
			case r == 105 : // ['i','i']
				return 34
			case 106 <= r && r <= 122 : // ['j','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 35
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 36
			case 98 <= r && r <= 122 : // ['b','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 37
			case r == 117 : // ['u','u']
				return 38
			case 118 <= r && r <= 122 : // ['v','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			case r == 125 : // ['}','}']
				return 39
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case 48 <= r && r <= 57 : // ['0','9']
				return 23
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 40
			
			
			default:
				return 24
			}
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 41
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			default:
				return 42
			}
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 43
			case r == 47 : // ['/','/']
				return 43
			case 48 <= r && r <= 57 : // ['0','9']
				return 27
			case 65 <= r && r <= 90 : // ['A','Z']
				return 28
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 43
			case r == 47 : // ['/','/']
				return 43
			case 48 <= r && r <= 57 : // ['0','9']
				return 27
			case 65 <= r && r <= 90 : // ['A','Z']
				return 28
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 107 : // ['a','k']
				return 11
			case r == 108 : // ['l','l']
				return 44
			case 109 <= r && r <= 122 : // ['m','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 45
			case 117 <= r && r <= 122 : // ['u','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 119 : // ['a','w']
				return 11
			case r == 120 : // ['x','x']
				return 46
			case 121 <= r && r <= 122 : // ['y','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 47
			case 98 <= r && r <= 99 : // ['b','c']
				return 11
			case r == 100 : // ['d','d']
				return 48
			case 101 <= r && r <= 122 : // ['e','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 109 : // ['a','m']
				return 11
			case r == 110 : // ['n','n']
				return 49
			case 111 <= r && r <= 122 : // ['o','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 113 : // ['a','q']
				return 11
			case r == 114 : // ['r','r']
				return 50
			case 115 <= r && r <= 122 : // ['s','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 51
			case 117 <= r && r <= 122 : // ['u','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 99 : // ['a','c']
				return 11
			case r == 100 : // ['d','d']
				return 52
			case 101 <= r && r <= 122 : // ['e','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 108 : // ['a','l']
				return 11
			case r == 109 : // ['m','m']
				return 53
			case 110 <= r && r <= 122 : // ['n','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 42 : // ['*','*']
				return 40
			case r == 47 : // ['/','/']
				return 54
			
			
			default:
				return 24
			}
			
		},
	
		// S41
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case r == 10 : // ['\n','\n']
				return 41
			
			
			default:
				return 42
			}
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 43
			case r == 47 : // ['/','/']
				return 43
			case 48 <= r && r <= 57 : // ['0','9']
				return 27
			case 65 <= r && r <= 90 : // ['A','Z']
				return 28
			case r == 95 : // ['_','_']
				return 43
			case 97 <= r && r <= 122 : // ['a','z']
				return 28
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 55
			case 117 <= r && r <= 122 : // ['u','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 56
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 109 : // ['a','m']
				return 11
			case r == 110 : // ['n','n']
				return 57
			case 111 <= r && r <= 122 : // ['o','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 104 : // ['a','h']
				return 11
			case r == 105 : // ['i','i']
				return 58
			case 106 <= r && r <= 122 : // ['j','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 98 : // ['a','b']
				return 11
			case r == 99 : // ['c','c']
				return 59
			case 100 <= r && r <= 122 : // ['d','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S51
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 60
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S52
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 99 : // ['a','c']
				return 11
			case r == 100 : // ['d','d']
				return 61
			case 101 <= r && r <= 122 : // ['e','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S53
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S54
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S55
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 62
			case 98 <= r && r <= 122 : // ['b','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S56
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 114 : // ['a','r']
				return 11
			case r == 115 : // ['s','s']
				return 63
			case 116 <= r && r <= 122 : // ['t','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S57
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S58
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case r == 97 : // ['a','a']
				return 64
			case 98 <= r && r <= 122 : // ['b','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S59
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 65
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S60
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 66
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S61
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 67
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S62
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S63
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 68
			case 117 <= r && r <= 122 : // ['u','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S64
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 109 : // ['a','m']
				return 11
			case r == 110 : // ['n','n']
				return 69
			case 111 <= r && r <= 122 : // ['o','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S65
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 109 : // ['a','m']
				return 11
			case r == 110 : // ['n','n']
				return 70
			case 111 <= r && r <= 122 : // ['o','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S66
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 114 : // ['a','r']
				return 11
			case r == 115 : // ['s','s']
				return 71
			case 116 <= r && r <= 122 : // ['t','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S67
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 117 : // ['a','u']
				return 11
			case r == 118 : // ['v','v']
				return 72
			case 119 <= r && r <= 122 : // ['w','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S68
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S69
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S70
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 115 : // ['a','s']
				return 11
			case r == 116 : // ['t','t']
				return 73
			case 117 <= r && r <= 122 : // ['u','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S71
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 104 : // ['a','h']
				return 11
			case r == 105 : // ['i','i']
				return 74
			case 106 <= r && r <= 122 : // ['j','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S72
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S73
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 104 : // ['a','h']
				return 11
			case r == 105 : // ['i','i']
				return 75
			case 106 <= r && r <= 122 : // ['j','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S74
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 102 : // ['a','f']
				return 11
			case r == 103 : // ['g','g']
				return 76
			case 104 <= r && r <= 122 : // ['h','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S75
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 107 : // ['a','k']
				return 11
			case r == 108 : // ['l','l']
				return 77
			case 109 <= r && r <= 122 : // ['m','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S76
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 109 : // ['a','m']
				return 11
			case r == 110 : // ['n','n']
				return 78
			case 111 <= r && r <= 122 : // ['o','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S77
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 79
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S78
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 100 : // ['a','d']
				return 11
			case r == 101 : // ['e','e']
				return 80
			case 102 <= r && r <= 122 : // ['f','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S79
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S80
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 99 : // ['a','c']
				return 11
			case r == 100 : // ['d','d']
				return 81
			case 101 <= r && r <= 122 : // ['e','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
		// S81
		func(r rune) int {
			switch {
			case r == 45 : // ['-','-']
				return 25
			case r == 47 : // ['/','/']
				return 25
			case 48 <= r && r <= 57 : // ['0','9']
				return 7
			case 65 <= r && r <= 90 : // ['A','Z']
				return 11
			case r == 95 : // ['_','_']
				return 25
			case 97 <= r && r <= 122 : // ['a','z']
				return 11
			
			
			
			}
			return NoState
			
		},
	
}
