package logger

import (
  "log"
  "fmt"
  "io/ioutil"
  "os"
  "time"
)
const logDir = "config/logs"
func Log(str ...string) {
  makeLogDir()

  line := ""
  for _, s := range str {
    line = line + s
  }
  line = line + "\r\n"
  log.Print(line)
  t := time.Now()
  logname := fmt.Sprintf("%s_log.log",t.Format("200601"))
  logfile := fmt.Sprintf("%s/%s",logDir, logname)
  line = fmt.Sprintf("%s %s",t.Format("2006/01/02 15:04:05"), line)
  if !FileExists(logfile) {
    err := ioutil.WriteFile(logfile, []byte(line), 0755)
    CheckError(err)
  } else {
    file, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0755)
    CheckError(err)
    defer file.Close()
    if _, err := file.WriteString(line); err != nil {
      CheckError(err)
    }
  }

}
func CheckError(err error) (ok bool) {
  // here for obsolescence
  return ErrorOK(err)
}
func ErrorOK(err error) (ok bool) {
  ok = true // All is OK, err == nil
  if err != nil {
    ok = false // There's an error, print it
    e := fmt.Sprintf("%v",err)
    Log(e)
  }
  return
}
func Fatal(err error) {
  e := fmt.Sprintf("%v",err)
  Log(e)
  log.Fatal(err)
}
func makeLogDir(){
  _, err := os.Stat(logDir)
  if os.IsNotExist(err) {
    os.Mkdir(logDir, 0755)
  }
}
func FileExists(file string) bool {
  exists := false
  _, err := os.Stat(file);
  if os.IsNotExist(err) {
    log.Print("File does not exist : ", file);
  } else if err == nil {
    exists = true
  } else {
    log.Fatal(err)
    log.Print("File %s does not exist\n", file);
  }
  return exists
}
