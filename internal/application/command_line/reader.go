package command_line

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/jpdejavite/messari-THA-solution/internal/application/usecase"
)

func ReadFromCommanLine(rd io.Reader) {
	start := time.Now()

	reader := bufio.NewReader(rd)

	processTrade := usecase.NewProcessTrade()

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "BEGIN" {
			continue
		}

		if text == "END" {
			break
		}

		processTrade.Execute(text)
	}

	for _, marketSummary := range processTrade.GetTradesSummary() {
		b, err := json.Marshal(marketSummary)
		if err != nil {
			// TODO handle error
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}

	elapsed := time.Since(start)
	log.Printf("execution time took %s", elapsed)
}
