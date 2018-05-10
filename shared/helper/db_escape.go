package helper

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//THIS files funtions has been borrowed from
func escapeString(txt string) string {
	var (
		esc string
		buf bytes.Buffer
	)
	last := 0
	for ii, bb := range txt {
		switch bb {
		case 0:
			esc = `\0`
		case '\n':
			esc = `\n`
		case '\r':
			esc = `\r`
		case '\\':
			esc = `\\`
		case '\'':
			esc = `\'`
		case '"':
			esc = `\"`
		case '\032':
			esc = `\Z`
		default:
			continue
		}
		io.WriteString(&buf, txt[last:ii])
		io.WriteString(&buf, esc)
		last = ii + 1
	}
	io.WriteString(&buf, txt[last:])
	return buf.String()
}

func escapeQuotes(txt string) string {
	var buf bytes.Buffer
	last := 0
	for ii, bb := range txt {
		if bb == '\'' {
			io.WriteString(&buf, txt[last:ii])
			io.WriteString(&buf, `''`)
			last = ii + 1
		}
	}
	io.WriteString(&buf, txt[last:])
	return buf.String()
}

// Escapes special characters in the txt, so it is safe to place returned string
// to Query method.
func MySqlEscape(txt string) string {
	//ME: by me-- not using commented featuers
	//if c.Status()&SERVER_STATUS_NO_BACKSLASH_ESCAPES != 0 {
	//	return escapeQuotes(txt)
	//}
	return escapeString(txt)
}

func SqlManyDollars(colSize, repeat int, isMysql bool) string {
	if isMysql {
		s := strings.Repeat("?,", colSize)
		s = "(" + s[0:len(s)-1] + "),"
		insVals_ := strings.Repeat(s, repeat)

		return insVals_[0 : len(insVals_)-1]
	}

	buff := bytes.NewBufferString("")
	cnt := 1
	for i := 0; i < repeat; i++ {
		buff.WriteString("(")
		for j := 0; j < colSize; j++ {
			buff.WriteString(fmt.Sprintf("$%d", cnt))
			if j+1 != colSize {
				buff.WriteString(",")
			}
			cnt++
		}
		buff.WriteString(")")
		if i+1 != repeat {
			buff.WriteString(",")
		}
	}
	return buff.String()
}
