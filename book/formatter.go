package book

type BookFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	FileImage        string `json:"file_image"`
	Quantity         int    `json:"quantity"`
	Slug             string `json:"slug"`
}

func FormatBook(book Book) BookFormatter {
	bookFormatter := BookFormatter{}
	bookFormatter.ID = book.ID
	bookFormatter.UserID = book.UserID
	bookFormatter.Name = book.Name
	bookFormatter.ShortDescription = book.ShortDescription
	bookFormatter.FileImage = book.FileImage
	bookFormatter.Quantity = book.Quantity
	bookFormatter.Slug = book.Slug

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

type BookDetailFormatter struct {
	ID               int               `json:"id"`
	UserID           int               `json:"user_id"`
	Name             string            `json:"name"`
	ShortDescription string            `json:"short_description"`
	Description      string            `json:"description"`
	FileImage        string            `json:"file_image"`
	Quantity         int               `json:"quantity"`
	Slug             string            `json:"slug"`
	User             BookUserFormatter `json:"user"`
}

type BookUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatBookDetail(book Book) BookDetailFormatter {
	bookDetailFormatter := BookDetailFormatter{}
	bookDetailFormatter.ID = book.ID
	bookDetailFormatter.UserID = book.UserID
	bookDetailFormatter.Name = book.Name
	bookDetailFormatter.ShortDescription = book.ShortDescription
	bookDetailFormatter.Description = book.Description
	bookDetailFormatter.FileImage = book.FileImage
	bookDetailFormatter.Quantity = book.Quantity
	bookDetailFormatter.Slug = book.Slug

	user := book.User

	bookUserFormatter := BookUserFormatter{}
	bookUserFormatter.Name = user.Name
	bookUserFormatter.ImageURL = user.ProfilePic

	bookDetailFormatter.User = bookUserFormatter

	return bookDetailFormatter
}