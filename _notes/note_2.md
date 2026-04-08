# Sliding Window // Скользящее окно

**Sliding Window (скользящее окно)** — это техника в алгоритмах и программировании, которая позволяет эффективно обрабатывать подмассивы или подстроки фиксированного или переменного размера, “двигая окно” по данным вместо пересчёта заново каждый раз.

Сложность O(n)

Каркас приложения:

```golang
func main() {
	begin := 0
	windowState := 0// состояние рамки
    result := 0

	for end := 0; end < len(nums); end++ {
        windowState += end
        end - begin + 1 // window size
		if // window_condition {
			begin += 1 // shrink window
		}
	}
}
```