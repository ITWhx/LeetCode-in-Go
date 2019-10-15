/**
*
*@author 吴昊轩
*@create 2019-10-1517:07
 */
package _151_Reverse_Words_in_a_String

import "strings"

func reverseWords(s string) string {
	parts := strings.Split(s, " ")
	size := len(parts)
	rparts := []string{}
	for i := size - 1; i >= 0; i-- {
		if parts[i] != "" {
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, " ")
}
