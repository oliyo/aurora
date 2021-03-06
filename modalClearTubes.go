package main

import (
	"fmt"
	"html"

	"github.com/kr/beanstalk"
)

// modalClearTubes render modal popup for delete job in tubes.
func modalClearTubes(server string) string {
	var err error
	var tubeList string
	var bstkConn *beanstalk.Conn
	if bstkConn, err = beanstalk.Dial("tcp", server); err != nil {
		return ``
	}
	tubes, _ := bstkConn.ListTubes()
	for _, v := range tubes {
		tubeList += fmt.Sprintf(`<div class="checkbox">
                                <label class="">
                                    <input type="checkbox" name="%s" value="1">
                                    <b>%s</b>
                                </label>
                            </div>`, v, html.EscapeString(v))
	}
	return fmt.Sprintf(`<div class="modal fade" id="clear-tubes" data-cookie="tubefilter" tabindex="-1" role="dialog" aria-labelledby="clear-tubes-label" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title" id="clear-tubes-label">Clear multiple tubes</h4>
            </div>
            <div class="modal-body">
                <form>
                    <fieldset>
                        <div class="form-group">
                            <label for="focusedInput">Tube name
                                <small class="text-muted">(supports <a href="http://james.padolsey.com/javascript/regex-selector-for-jquery/" target="_blank">jQuery
                                        regexp</a> syntax)
                                </small>
                            </label>

                            <div class="input-group">
                                <input class="form-control focused" id="tubeSelector" type="text" placeholder="prefix*"
                                       value="%s">

                                <div class="input-group-btn">
                                    <a href="#" class="btn btn-info" id="clearTubesSelect">Select</a>
                                </div>

                            </div>

                        </div>
                    </fieldset>
                    <div>
                        <strong>Tube list</strong>
                        %s
                    </div>
                </form>
            </div>
            <div class="modal-footer">

                <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                <a href="#" class="btn btn-success" id="clearTubes">Clear selected tubes</a>
                <br/><br/>

                <p class="text-muted text-right small">
                    * Tube clear works by peeking to all jobs and deleting them in a loop.
                </p>
            </div>
        </div>
    </div>
</div>`, selfConf.TubeSelector, tubeList)
}
