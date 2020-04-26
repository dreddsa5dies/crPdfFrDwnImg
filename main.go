package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/signintech/gopdf"

	s "github.com/dreddsa5dies/simpleDirWalker"
)

type timeSlice []s.PathAndFileInfo

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].FileInfo.ModTime().Before(p[j].FileInfo.ModTime())
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(os.Args[0] + " catalog")
	} else {
		lstFiles, err := s.SDW(os.Args[1])
		if err != nil {
			log.Println(err)
		}

		for _, v := range lstFiles {
			if v.FileInfo.IsDir() {
				log.Println(v.FileInfo.Name())
				var lstJPG timeSlice
				lstJPG, err = s.SDW(v.FileInfo.Name())
				if err != nil {
					log.Println(err)
				}

				pdf := gopdf.GoPdf{}
				pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 841.89, H: 595.28}}) //595.28, 841.89 = A4

				sort.Sort(lstJPG)

				for _, val := range lstJPG {
					pdf.AddPage()

					pdf.Image(val.FullPath, 10, 10, &gopdf.Rect{W: 831.89, H: 585.28}) //print image
				}

				pdf.WritePdf(v.FileInfo.Name() + ".pdf")
			}
		}
	}
}
