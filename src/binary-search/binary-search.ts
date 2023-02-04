export function binary_search(array: number[], element: number) {
  let n = array.length;
  let left = 0;
  let right = n - 1;

  while (left <= right) {
    let midFloat = (left + right) / 2;
    let mid = midFloat - (midFloat % 1);

    if (array[mid] > element) {
      right = mid - 1;
    } else if (array[mid] < element) {
      left = mid + 1;
    } else {
      return mid;
    }
  }

  return false;
}
