export function isPalindrome(s: string): boolean {
  const word = s.toLowerCase().replace(/[^0-9a-z]/gi, "");

  let lp = 0;
  let rp = word.length - 1;

  while (lp < rp) {
    if (word[lp] != word[rp]) {
      return false;
    }

    lp++;
    rp--;
  }

  return true;
}
