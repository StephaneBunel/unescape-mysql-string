package main

// UnescapeMysqlString unescape string escaped by Mysql
func UnescapeMysqlString(es string) []byte {
	s, d := 0, 0
	us := make([]byte, len(es))
	for s < len(es) {
		c := es[s]
		s++
		if c == '\\' {
			c = es[s]
			s++
			switch c {
			case '0':
				us[d] = '\x00'
			case '\'', '\\', '"':
				us[d] = c
			case 'b':
				us[d] = '\b'
			case 'n':
				us[d] = '\n'
			case 'r':
				us[d] = '\r'
			case 't':
				us[d] = '\t'
			case 'Z':
				us[d] = '\x1a'
			default:
				us[d] = '\\'
				d++
				us[d] = c
			}
			d++
		} else {
			us[d] = c
			d++
		}
	}
	return us[:d]
}
