const reverseString = (s) => {
	let temp 
	const length = s.length;
	for (i = 0; i < length / 2; i++) {
		temp = s[i]
		s[i] = s[length-1-i]
		s[length-1-i] = temp
	}
}

function main() {
	const input = ["H", "e", "l", "l", "o"]
	reverseString(input)
	console.log(input)
}

main()