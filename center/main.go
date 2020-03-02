package center

import (
	"rpg-demo/center/rpc"
)

func StartCenter() {

	rpc.StartGrpcService()
}
