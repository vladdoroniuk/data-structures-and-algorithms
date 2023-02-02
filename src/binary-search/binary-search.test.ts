import { describe, it, expect, beforeEach } from "vitest";
import { binary_search } from "./binary-search";

describe.concurrent("Binary search", () => {
  it("finds the element and returns its index", () => {
    let list_1 = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element_1 = 14;
    expect(binary_search(list_1, element_1)).toBe(9);

    let list_2 = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element_2 = 4;
    expect(binary_search(list_2, element_2)).toBe(2);
  });

  it("doesn't find the element and returns false", () => {
    let list = [1, 2, 4, 6, 7, 8, 10, 10, 10, 14, 15, 17, 17, 20];
    let element = 9;
    expect(binary_search(list, element)).toBe(false);
  });

  it("gets the empty array and returns false", () => {
    let list = [];
    let element = 1;
    expect(binary_search(list, element)).toBe(false);
  });
});
