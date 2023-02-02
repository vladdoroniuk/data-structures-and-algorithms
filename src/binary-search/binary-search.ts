export function binary_search(list: number[], element: number) {
  let n = list.length;
  let left = 0;
  let right = n - 1;

  while (left <= right) {
    let mid = Math.floor((left + right) / 2);

    if (list[mid] > element) {
      right = mid - 1;
    } else if (list[mid] < element) {
      left = mid + 1;
    } else {
      return mid;
    }
  }

  return false;
}
