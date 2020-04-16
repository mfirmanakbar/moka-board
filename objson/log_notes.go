package objson

type LogNotes struct {
	Products        LogNotesDetail `json:"products"`
	Customers       LogNotesDetail `json:"customers"`
	Sales           LogNotesDetail `json:"sales"`
	SalesReturns    LogNotesDetail `json:"returns"`
	ReceivePayments LogNotesDetail `json:"receive_payments"`
	Others          LogNotesDetail `json:"others"`
}

type LogNotesDetail struct {
	Failed int `json:"failed"`
	Total  int `json:"total"`
}
