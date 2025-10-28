package server

type MathServer struct {
}

type Args struct {
	A, B int
}

func (ms *MathServer) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}
