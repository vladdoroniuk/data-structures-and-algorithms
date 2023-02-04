import { describe, it, expect } from "vitest";
import { selection_sort } from "./selection-sort";

describe.concurrent("Bubble sort", () => {
  it("sorts array with the 'selection_sort' function", () => {
    let sorted_array = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    let shuffled_array = [6, 2, 8, 9, 3, 4, 7, 1, 5, 10];
    let reversed_array = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1];

    selection_sort(sorted_array);
    expect(sorted_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    selection_sort(shuffled_array);
    expect(shuffled_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);

    selection_sort(reversed_array);
    expect(reversed_array).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  });
});
