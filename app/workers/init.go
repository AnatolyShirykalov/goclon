package workers

import (
	"../../config/admin"
	"../../parser"
	"errors"
	"fmt"
	"github.com/qor/worker"
)

type ParserJobArgument struct {
	Start int64
	End   int64
}

func (a ParserJobArgument) String() string {
	return fmt.Sprintf("Argument: start: %v, end: %v", a.Start, a.End)
}

var Worker *worker.Worker

func Init() {
	Worker = worker.New()
	Worker.RegisterJob(&worker.Job{
		Name: "Parse Classic-online.ru",
		Handler: func(arg interface{}, qorJob worker.QorJobInterface) error {
			qorJob.AddLog("Started parsing classic-online.ru")
			args := arg.(*ParserJobArgument)
			qorJob.AddLog(args.String())
			if args.Start >= args.End {
				return errors.New("Bad range (start >= end)")
			}
			client := parser.NewClient()
			for i := args.Start; i < args.End; i++ {
				uri := fmt.Sprintf("http://classic-online.ru/archive/?file_id=%d", i)
				parser.Page(client, uri)
				qorJob.SetProgress(uint((100 * (i - args.Start)) / (args.End - args.Start)))
			}
			return nil
		},
		Resource: admin.Admin.NewResource(&ParserJobArgument{}),
	})

	admin.Admin.AddResource(Worker)
}
