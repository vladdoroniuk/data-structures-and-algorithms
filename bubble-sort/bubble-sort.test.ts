import { describe, it, beforeEach, expect } from "vitest";
import { bubble_sort, bubble_sort_optimized } from "./bubble-sort";

describe.concurrent("Bubble sort", () => {
  let sorted_list: number[];
  let reversed_list: number[];

  beforeEach(() => {
    sorted_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    reversed_list = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1];
  });

  it("sorts list with the 'bubble_sort' function", () => {
    bubble_sort(sorted_list);
    expect(sorted_list).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort(reversed_list);
    expect(reversed_list).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  });

  it("sorts list with the 'bubble_sort_optimized' function", () => {
    bubble_sort_optimized(sorted_list);
    expect(sorted_list).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort_optimized(reversed_list);
    expect(reversed_list).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  });
});
