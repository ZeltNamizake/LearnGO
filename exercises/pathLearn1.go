package main
import(
 "fmt"
 "path/filepath"
)

/*
.Join = join path
.Base = get file name
.Dir = get folder name
.Ext = get extension file
/*

/*func main() {
  pathFile := filepath.Join("/sdcard", "Music", "music.mp3")
  fmt.Println("Base: ", filepath.Base(pathFile))
  fmt.Println("Dir: ", filepath.Dir(pathFile))
  fmt.Println("Ext: ", filepath.Ext(pathFile))
}*/

func main(){
  files := []string{
   "a.txt",
   "b.log",
   "c.jpg",
   "d.log",
  }
  
  for _, f := range files {
    if filepath.Ext(f) == ".log" {
      name := f
      ext := filepath.Ext(name)
      base := name[:len(name)-len(ext)]
      fmt.Println(base)
    }
  }
}