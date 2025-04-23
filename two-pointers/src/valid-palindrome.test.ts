import { expect, test } from "vitest";
import { isPalindrome } from "./valid-palindrome";

test("Test 1", () => {
  expect(isPalindrome("A man, a plan, a canal: Panama")).toBe(true);
});

test("Test 2", () => {
  expect(isPalindrome("race a car")).toBe(false);
});
