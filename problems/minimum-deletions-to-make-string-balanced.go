package problems

func minimumDeletions(s string) int {
    del := strings.Count(s, "a")
    ans := del
    for _, c := range s {
        // 'a' -> -1    'b' -> 1
        del += int((c-'a')*2 - 1)
        if del < ans {
            ans = del
        }
    }
    return ans
}
