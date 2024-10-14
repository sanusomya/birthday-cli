package wish

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sanusomya/birthday-cli/config"
)

func sendMessage(name string) {

	telegramUtils := config.GetTelegramVariables()
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage?text=happy birthday [%v](tg://user?id=%v)&chat_id=@%v&parse_mode=MarkdownV2",telegramUtils[0],name, telegramUtils[1],telegramUtils[2])
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sucessfully sent message for user "+name)
}
