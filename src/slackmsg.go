package main

import (
	//"bytes"
	//"log"
	//"net/http"
	"os"
	"text/template"
	"log"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("slackmsg <slack-url-file>")
	}
	sendMsg(os.Args[1])
}

func sendMsg(urlFile string) {
	type monitValues struct {
		MONIT_SERVICE             string
		MONIT_EVENT               string
		MONIT_DESCRIPTION         string
		MONIT_DATE                string
		MONIT_HOST                string
		MONIT_PROCESS_PID         string
		MONIT_PROCESS_MEMORY      string
		MONIT_PROCESS_CHILDREN    string
		MONIT_PROCESS_CPU_PERCENT string
		MONIT_PROGRAM_STATUS      string
	}

	values := monitValues{
		MONIT_SERVICE:             os.Getenv("MONIT_SERVICE"),
		MONIT_EVENT:               os.Getenv("MONIT_EVENT"),
		MONIT_DESCRIPTION:         os.Getenv("MONIT_DESCRIPTION"),
		MONIT_DATE:                os.Getenv("MONIT_DATE"),
		MONIT_HOST:                os.Getenv("MONIT_HOST"),
		MONIT_PROCESS_PID:         os.Getenv("MONIT_PROCESS_PID"),
		MONIT_PROCESS_MEMORY:      os.Getenv("MONIT_PROCESS_MEMORY"),
		MONIT_PROCESS_CHILDREN:    os.Getenv("MONIT_PROCESS_CHILDREN"),
		MONIT_PROCESS_CPU_PERCENT: os.Getenv("MONIT_PROCESS_CPU_PERCENT"),
		MONIT_PROGRAM_STATUS:      os.Getenv("MONIT_PROGRAM_STATUS"),
	}


	payload := `{
    "text": "monit alert -- {{.MONIT_EVENT}}",
    "attachments": [
        {
            "text": "{{.MONIT_EVENT}} Service {{.MONIT_SERVICE}}\n{{.MONIT_DESCRIPTION}}",
					"fields": [
		{ "title": "Date", "value": "{{.MONIT_DATE}}", "short": true },
		{ "title": "Host", "value": "{{.MONIT_HOST}}", "short": true }
				]

        }
    ]
}
`
	dat, err := ioutil.ReadFile(urlFile)
	if err != nil {
		log.Fatalf("Can not access slack-url-file: %s - %s", urlFile, err.Error())
	}
	url:= strings.TrimSpace(string(dat))

	t := template.Must(template.New("monitMsg").Parse(payload))
	buf := bytes.Buffer{}
	if err := t.Execute(&buf, values); err != nil {
		log.Fatal("Error: ", err)
	}


	req, err := http.NewRequest("POST", url, &buf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Send to Slack Failed: %s", err)
	}
}
