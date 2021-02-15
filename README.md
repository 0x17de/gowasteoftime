# gowasteoftime

A go library for parsing timestamps.

## Example Usage

```
package main

import (
  "fmt"
  "log"
  "time"
  wt "github.com/0x17de/gowasteoftime/pkg/wasteoftime"
)

func main() {
  layout, err := wt.ParseLayout("%Y-%m-%d %H:%M:%S")
  if err != nil {
    log.Fatalf("Failed to parse date: %v", err)
  }

  t1, err := wt.ParseDate(layout, "2006-01-02 15:04:05")
  if err != nil {
    log.Fatalf("Failed to parse date: %v", err)
  }
  fmt.Printf("Date: %s\n", t1.Time().Format(time.RFC3339))

  t2, err := wt.ParseDateWithFormat("%Y-%m-%d %H:%M:%S", "2006-01-02 15:04:05")
  if err != nil {
    log.Fatalf("Failed to parse date: %v", err)
  }
  fmt.Printf("Date: %s\n", t2.Time().Format(time.RFC3339))

  return
}
```

## Format Flags

| Formatter | Description                                                                  | Example                   |
|-----------|------------------------------------------------------------------------------|---------------------------|
| %Y        | 4-digit year                                                                 | 2006                      |
| %m        | 2-digit month                                                                | 02                        |
| %d        | 2-digit day                                                                  | 27                        |
| %H        | 2-digit hour                                                                 | 14                        |
| %M        | 2-digit minute                                                               | 53                        |
| %S        | 2-digit second                                                               | 59                        |
| %p        | AM/PM modifier. Used together with %H                                        | 14 AM                     |
| %F        | Optional fraction, 0-9 digits                                                | .123, ., .123456789                      |
| %b        | Month as text                                                                | Jan, January              |
| %a        | Weekday. Not used for parsing                                                | Mon, Monday               |
| %z        | Timezone                                                                     | EST                       |
| %N        | Unix timestamp (10 digits). Assumed to be milliseconds when passed 13 digits | 1136210645, 1136210645000 |
| %1m       | Month, either 1 or 2 digits                                                  | 4                         |
| %1d       | Day, either 1 or 2 digits                                                    | 4                         |
| %1H       | Hour, either 1 or 2 digits                                                   | 4                         |
| %1M       | Minute, either 1 or 2 digits                                                 | 4                         |
| %1S       | Second, either 1 or 2 digits                                                 | 4                         |
