export function bubble_sort(array: number[]) {
  let n = array.length;

  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - i - 1; j++) {
      if (array[j] > array[j + 1]) {
        let temp = array[j];
        array[j] = array[j + 1];
        array[j + 1] = temp;
      }
    }
  }
}

export function bubble_sort_optimized(array: number[]) {
  let n = array.length;

  for (let i = 0; i < n - 1; i++) {
    let is_swapped = false;
    for (let j = 0; j < n - i - 1; j++) {
      if (array[j] > array[j + 1]) {
        let temp = array[j];
        array[j] = array[j + 1];
        array[j + 1] = temp;
        is_swapped = true;
      }
    }

    if (!is_swapped) {
      break;
    }
  }
}
