# Z Function

Given a string s, Z function returns a list of integers that has the same size with s.
Let us call the returned list z and the ith index of z, z[i] i from 0 to n -1.

z[i] denotes, longest prefix of z that matches the string starting from z[i].
E.g.
	let s = "aabcaabxaaaz"
	then z becomes, [0, 1, 0, 0, 3, 1, 0, 0, 2, 2, 1, 0]
	here z[4] = 3 because, z[3-5] = aab which is a prefix of z, and this is the longest we can get for position 3.

Then, how is this z array is gonna help us. With slight modification we can use it to find all the occurences of a  string t inside string s. And we calculate this very efficiently.
We can append an extra character which s doesn't contain and then we can append t like:
s1 = t + "#" + s
assuming "#" is not in our alphabet and it is not included in s and t.
Here for every element in our z array for the positions after "#" sign, if z of such a position is equal to the length(t) then we know that t occur in that position of s.

To read extensive description of the algorithm:
https://cp-algorithms.com/string/z-function.html
