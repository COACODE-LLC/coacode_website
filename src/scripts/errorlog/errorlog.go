package errorlog

import(
  "log"
  "os"
)

func LogError(err error) {
  file, fileErr := os.OpenFile("logs", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  
  if err != nil {
    log.Println(fileErr)
  }
  
  defer file.Close()

  log.SetOutput(file)
  log.Println(err);
}
