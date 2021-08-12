package errcode

func TogRPCError(err*Error) error  {
	s := status.New(TogRPCCode(err.Code()), err.Msg())
	return s.Err()
}

func TogRPCCode(code int) codes.Code  {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthhenticated
	case AccessDenied.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case LimitExceed.Code():
		statusCode = codes.ResourceExceeded
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknow
	}
	return statusCode
}

