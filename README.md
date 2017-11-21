# cnwordcount

Count Chinese character numbers in a file.

统计文件中中文字符的个数。

## Usage

```bash
cnwordcount -f <filename>
```

## Output

```bash
$ cnwordcount -f README.md
README.md 12
```

## Principle

Based on Chinese words range in Unicode.

Commonly used **0x4e00** to **0x9fa5**.

Reference https://jimmysong.io/cheatsheets/unicode