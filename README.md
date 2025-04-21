# 📚 Text File Word Counter

This Go project reads a `.txt` file and analyzes the word frequency within it. It displays the **total number of words** and the **top 5 most frequently used words** in the file.

## ✅ Features

- Reads text files line by line
- Cleans and normalizes words (e.g., removes punctuation, lowercases)
- Calculates word frequency using a `map[string]int`
- Displays:
  - Total number of words
  - Top 5 most frequent words with their counts

---

## 🛠️ Requirements

- Go 1.20 or later

---

## 📂 Project Structure

```text
projectGolang/ 
├── main.go
├── file.txt
├── go.mod
├── README.md
├── file/
│   └── reader.go
├── counter/
│   └── counter.go
```

---

## 🚀 Usage Instructions

1. **Clone the repository** (or place your `.go` files as shown above).

    ```bash
    git clone https://github.com/Olzzhass/projectGo.git

2. **Initialize Go module** (if not already):

   ```bash
   go mod init projectGolang

3. **Run the program**

    ```bash
    go run main.go <path-to-your-text-file>

4. **Example Output**

```text
Total words in file: 372
Top 5 most frequent words:
1. the (28 occurrences)
2. and (17 occurrences)
3. of (15 occurrences)
4. to (14 occurrences)
5. a (13 occurrences)
```