package util

func insertSort(arr []float64, l int, r int) {
    for i := l + 1; i <= r; i++ {
        if arr[i-1] > arr[i] {
            t := arr[i]
            j := i
            for j > l && arr[j-1] > t {
                arr[j] = arr[j-1]
                j--
            }
            arr[j] = t
        }
    }
}

func getPivot(arr []float64, l int, r int) float64 {
    if l == r {
        return arr[l]
    }
    i := 0
    n := 0
    for i = l; i < r-5; i += 5 {
        insertSort(arr, i, i+4)
        n = i - l
        arr[l+n/5], arr[i+2] = arr[i+2], arr[l+n/5]
    }
    num := r - i + 1
    if num > 0 {
        insertSort(arr, i, i+num-1)
        n = i - l
        arr[l+n/5], arr[i+num/2] = arr[i+num/2], arr[l+n/5]
    }
    n /= 5
    if n == l {
        return arr[l]
    }
    return getPivot(arr, l, l+n)
}

func getPivotIndex(arr []float64, l int, r int, pivot float64) int {
    for i := l; i <= r; i++ {
        if arr[i] == pivot {
            return i
        }
    }
    return -1
}

func partition(arr []float64, l int, r int, pivotIndex int) int {
    arr[pivotIndex], arr[l] = arr[l], arr[pivotIndex]
    i := l
    j := r
    pivot := arr[l]
    for i < j {
        for arr[j] >= pivot && i < j {
            j--
        }
        arr[i] = arr[j]
        for arr[i] <= pivot && i < j {
            i++
        }
        arr[j] = arr[i]
    }
    arr[i] = pivot
    return i
}

func bfptrFloat64(arr []float64, l int, r int, k int) float64 {
    pivot := getPivot(arr, l, r)
    pivotIndex := getPivotIndex(arr, l, r, pivot)
    i := partition(arr, l, r, pivotIndex)
    m := i - l + 1
    if m == k {
        return arr[i]
    }
    if m > k {
        return bfptrFloat64(arr, l, i-1, k)
    }
    return bfptrFloat64(arr, i+1, r, k-m)
}
