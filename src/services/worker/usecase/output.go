package usecase

type LoginOutput struct {
	Session string `json:"session"`
}

func ZeroLoginOutput() LoginOutput {
	return LoginOutput{
		Session: "",
	}
}

type AddOutput struct {
	ID string `json:"id"`
}

func ZeroAddOutput() AddOutput {
	return AddOutput{
		ID: "",
	}
}
