export function selection_sort(array: number[]) {
  let n = array.length;

  for (let i = 0; i < n; i++) {
    let jMin = i;
    for (let j = i + 1; j < n; j++) {
      if (array[j] < array[jMin]) {
        jMin = j;
      }
    }

    if (jMin !== i) {
      let temp = array[i];
      array[i] = array[jMin];
      array[jMin] = temp;
    }
  }
}
