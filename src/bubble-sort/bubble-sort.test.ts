import { describe, it, beforeEach, expect } from "vitest";
import { bubble_sort, bubble_sort_optimized } from "./bubble-sort";

describe.concurrent("Bubble sort", () => {
  let sorted_array: number[];
  let shuffled_array: number[];
  let reversed_array: number[];

  beforeEach(() => {
    sorted_array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    shuffled_array = [6, 2, 8, 9, 3, 4, 7, 1, 5, 10];
    reversed_array = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1];
  });

  it("sorts array with the 'bubble_sort' function", () => {
    bubble_sort(sorted_array);
    expect(sorted_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort(shuffled_array);
    expect(shuffled_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort(reversed_array);
    expect(reversed_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  });

  it("sorts array with the 'bubble_sort_optimized' function", () => {
    bubble_sort_optimized(sorted_array);
    expect(sorted_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort(shuffled_array);
    expect(shuffled_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    bubble_sort_optimized(reversed_array);
    expect(reversed_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  });
});
