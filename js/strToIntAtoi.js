var myAtoi = (s) => {
    var str = s.trim(' ');
    let res = '';
    let mult = 1;
    let usedMult = false;
    for (let ch of str) {
        if (ch === '-' && !usedMult) { mult = -1; usedMult = true; continue; }
        if (ch === '+'&& !usedMult)  { usedMult = true; continue; }
        if (Number(ch) || ch == '0') {
            res += ch; 
        }
        else if (res !== '') break;
        else return 0;
    }
    
    let val = Number(res) * mult;

    if (val < -2147483648) return -2147483648;
    if (val > 2147483647) return 2147483647;
    return val;
};

console.log(myAtoi("21474836460"));

