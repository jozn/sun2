package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	//"ms/sun/helper"
	"os"
	"path/filepath"
	"regexp"
    //"ms/sun/helper"
    "strconv"
)


//const DIR = `E:\WEB\files\file\dl\file\pic\music\2013\06\`
/*const DIR = `E:\WEB\files\file\dl\file\pic\photo\`
const OUT_LOG = `C:\Go\_gopath\src\ms\sun2\scripts\server_migration\logs\`
const MOVE_DIR = `E:\WEB\files\file\dl\file\pic\photo_moved\`*/
const DIR = `/dl/photo/`
const OUT_LOG = `/dl/photo/`
const MOVE_DIR = `/dl/photo_moved/`

var filesToProcced = make(chan fileName, 10000)
var dirsToProcced = make(chan string)

var allFiles = make(map[string][]fileName, 1000000)
var toIgnore = make([]fileName, 0)
var toMove = make([]fileName, 0)
var unknown = make([]fileName, 0)

var cnt = 0

func main() {
	go proccedFiles()
	proccedDirs()

	walkDir(DIR)

	selectImage()
	logResult()

	move(true)
}

func proccedDirs() {

}

func walkDir(dir string) {
	fmt.Println("dir: ", dir)
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			//walkDir(path)
			//dirsToProcced <- path
			if dir != path {
				walkDir(path)
			}
		} else {
			fn := fileName{
				path: path,
				name: f.Name(),
			}
			fn.extract()
			filesToProcced <- fn
		}
		return nil
	})
}

func selectImage() {
	for _, fms := range allFiles {
		sel := fms[0]
		for _, fn := range fms {
			if sel.size < fn.size {
				sel = fn
			}
		}
		toMove = append(toMove, sel)
		for _, fn := range fms {
			if sel != fn {
				toIgnore = append(toIgnore, fn)
			}
		}
	}
}

func proccedFiles() {
	for f := range filesToProcced {
		//helper.PertyPrint(f)
		//fmt.Printf("%# v\n", f)

		if f.valid {
			allFiles[f.baseImageName] = append(allFiles[f.baseImageName], f)
		} else {
			unknown = append(unknown, f)
		}
		cnt++
		if cnt%10000 == 0 {
			fmt.Println("procceded: ", cnt)
			fmt.Printf("%# v\n", f)
		}

	}
}

func sizeLog() {
	i := 0
	var totalSize = 0
	for _, fn := range toMove {
		info, err := os.Stat(fn.path)
		if err == nil {
			totalSize += int(info.Size())
		}
		i++
	}
	s := fmt.Sprintf("num: %d - size: %d ", i, totalSize)
	ioutil.WriteFile("size.txt", []byte(s), os.ModePerm)
}

func logResult() {
	moveBuff := bytes.NewBufferString("")
	for _, fn := range toMove {
		moveBuff.Write([]byte(fn.path + " \n"))
	}

	ignoreBuff := bytes.NewBufferString("")
	for _, fn := range toIgnore {
		ignoreBuff.Write([]byte(fn.path + " \n"))
	}

	unknownBuff := bytes.NewBufferString("")
	for _, fn := range unknown {
		unknownBuff.Write([]byte(fn.path + " \n"))
	}

	//movedResBuff := bytes.NewBufferString("")
	//for _, fn := range toMove {
	//	o := fmt.Sprintf("%s -> %s \n", fn.path, MOVE_DIR+fn.name)
	//	movedResBuff.Write([]byte(o))
	//}

	os.Chdir(OUT_LOG)

	ioutil.WriteFile("./move.txt", moveBuff.Bytes(), os.ModePerm)
	ioutil.WriteFile("./ignore.txt", ignoreBuff.Bytes(), os.ModePerm)
	ioutil.WriteFile("./unknow.txt", unknownBuff.Bytes(), os.ModePerm)
	//ioutil.WriteFile("moved_locations.txt", movedResBuff.Bytes(), os.ModePerm)

	move(false)
	sizeLog()
	//fmt.Printf("to moves: %# v\n", toMove)
	//fmt.Printf("to ignore: %# v\n", toIgnore)
	//fmt.Printf("unknown: %# v\n", unknown)
}

type fileName struct {
	name          string
	ext           string
	baseImageName string
	size          int
	path          string
	valid         bool
}

var regSize = regexp.MustCompile(`_(\d+)_?\.?`)
var regExt = regexp.MustCompile(`\.\w+`)

func (f *fileName) extract() {
	//f.ext = f.path[strings.Index(f.path, "."):]
	//r:=regSize.MatchString(f.path)
	f.ext = regExt.FindString(f.path)
	groups := regSize.FindStringSubmatch(f.path)
	if len(groups) == 2 {
		f.size = StrToInt(groups[1], 0)
	} else {
		f.size = 9999999999
	}

	if len(f.name) > 32 {
		f.baseImageName = f.name[:32]
		f.valid = true
	} else {
		f.valid = false
	}

	//fmt.Println(r)
}

func move(move bool) {
	os.MkdirAll(MOVE_DIR, os.ModeDir)
	i := 0
	dir_cnt := 0
	movedResBuff := bytes.NewBufferString("")
	for _, fn := range toMove {
		num := i % 2000
		out_dir := fmt.Sprintf("%sout_%d/", MOVE_DIR, dir_cnt)
		if num == 0 {
			os.MkdirAll(out_dir, os.ModeDir)
			dir_cnt++
		}
		oo := out_dir + fn.name
		o := fmt.Sprintf("%s -> %s \n", fn.path, oo)
		movedResBuff.Write([]byte(o))
		if move {
			os.Rename(fn.path, out_dir+fn.name)
		}
		i++
	}
	ioutil.WriteFile("./moved_locations.txt", movedResBuff.Bytes(), os.ModePerm)

}

func StrToInt(str string, defualt int) int {
    r64, err := strconv.ParseInt(str, 10, 64)
    if err != nil {
        return defualt
    }
    return int(r64)
}