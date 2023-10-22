function gcdOfStrings(str1, str2) {
    if (str1 + str2 !== str2 + str1) {
        return "";
    }

    const gcd = getGcd(str1.length, str2.length);
    return str2.substring(0, gcd);
}

function getGcd(a, b) {
    while (b !== 0) {
        const temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

console.log(gcdOfStrings("ABCABC","ABCABCABC"))

