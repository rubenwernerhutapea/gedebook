package book

type BookFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	FileImage        string `json:"file_image"`
	Quantity         int    `json:"quantity"`
}

func FormatBook(book Book) BookFormatter {
	bookFormatter := BookFormatter{}
	bookFormatter.ID = book.ID
	bookFormatter.UserID = book.UserID
	bookFormatter.Name = book.Name
	bookFormatter.ShortDescription = book.ShortDescription
	bookFormatter.FileImage = book.FileImage
	bookFormatter.Quantity = book.Quantity

	return bookFormatter
}

func FormatBooks(books []Book) []BookFormatter {
	booksFormatter := []BookFormatter{}

	for _, book := range books {
		bookFormatter := FormatBook(book)
		booksFormatter = append(booksFormatter, bookFormatter)
	}

	return booksFormatter
}