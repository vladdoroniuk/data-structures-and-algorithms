export function bubble_sort(list: number[]) {
  let n = list.length;

  for (let i = 0; i < n - 1; i++) {
    for (let j = 0; j < n - i - 1; j++) {
      if (list[j] > list[j + 1]) {
        let temp = list[j];
        list[j] = list[j + 1];
        list[j + 1] = temp;
      }
    }
  }
}

export function bubble_sort_optimized(list: number[]) {
  let n = list.length;

  for (let i = 0; i < n - 1; i++) {
    let is_swapped = false;
    for (let j = 0; j < n - i - 1; j++) {
      if (list[j] > list[j + 1]) {
        let temp = list[j];
        list[j] = list[j + 1];
        list[j + 1] = temp;
        is_swapped = true;
      }
    }

    if (!is_swapped) {
      break;
    }
  }
}
