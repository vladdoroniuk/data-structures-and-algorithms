import { describe, it, expect, beforeEach } from "vitest";
import { binary_search } from "./binary-search";

describe.concurrent("Binary search", () => {
  it("finds the element and returns its index", () => {
    let array_1 = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element_1 = 14;
    expect(binary_search(array_1, element_1)).toBe(9);

    let array_2 = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element_2 = 4;
    expect(binary_search(array_2, element_2)).toBe(2);
  });

  it("doesn't find the element and returns false", () => {
    let array = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element = 9;
    expect(binary_search(array, element)).toBe(false);
  });

  it("gets an empty array and returns false", () => {
    let array = [];
    let element = 1;
    expect(binary_search(array, element)).toBe(false);
  });
});
