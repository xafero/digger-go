package diggerclassic

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type ScoreStorage struct {
}

func NewScoreStorage() ScoreStorage {
	d := ScoreStorage{}
	return d
}

func CreateInStorage(mem *Scores) {
	WriteToStorage(mem)
}

func getScoreFile() string {
	fileName := "digger.sco"
	filePath, _ := filepath.Abs(fileName)
	return filePath
}

func ReadFromStorage(mem *Scores) bool {
	scoFile := getScoreFile()
	_, err := os.Stat(scoFile)
	if os.IsNotExist(err) {
		return false
	}
	fileIn, err := os.Open(scoFile)
	if err != nil {
		return false
	}
	defer fileIn.Close()
	br := bufio.NewScanner(fileIn)
	sc := make([]ScoreTuple, 10)
	for i := 0; i < 10; i++ {
		br.Scan()
		name := br.Text()
		br.Scan()
		score, _ := strconv.ParseInt(br.Text(), 10, 0)
		sc[i] = NewScoreTuple(name, int(score))
	}
	mem.scores = sc
	return true
}

func WriteToStorage(mem *Scores) bool {
	scoFile := getScoreFile()
	fileOut, err := os.Create(scoFile)
	if err != nil {
		return false
	}
	defer fileOut.Close()
	scoreinit := mem.scoreinit
	scorehigh := mem.scorehigh
	for i := 0; i < 10; i++ {
		fmt.Fprint(fileOut, scoreinit[i+1])
		fmt.Fprintln(fileOut)
		fmt.Fprintf(fileOut, "%v", scorehigh[i+2])
		fmt.Fprintln(fileOut)
	}
	return true
}
