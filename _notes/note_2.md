# Sliding Window // Скользящее окно

**Sliding Window (скользящее окно)** — это техника в алгоритмах и программировании, которая позволяет эффективно обрабатывать подмассивы или подстроки фиксированного или переменного размера, “двигая окно” по данным вместо пересчёта заново каждый раз.

Сложность O(n)

Каркас приложения:

```golang
func main() {
	begin := 0
	window_state // состояние рамки
    result

	for end := 0; end < len(nums); end++ {
        window_state
        window_size = end - begin + 1

		if // window_condition {
            result
            window_state
			begin += 1 // shrink window
		}
	}
    
    return result
}
```