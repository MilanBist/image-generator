package models


type Register struct{
	Username	string 	`json:"userName"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}


type Login struct{
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
}


type Response struct{
	Success		bool		`json:"success"`
	Message		string		`json:"message"`
	Data 		any			`json:"data,omitempty"`
	Error 		any 		`json:"error,omitempty"`
}