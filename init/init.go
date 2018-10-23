// this package must be imported first,
// to ensure it is initialized before all the other package.
// so it can print out the real earliest starting time.
package init

import (
	"log"

	"github.com/lovego/config"
)

func init() {
	log.Printf(`starting.(%s)`, config.Env())
}
