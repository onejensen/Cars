package datatypes

type Data struct {
	Path          string
	Manufacturers []Manufacturer
	Categories    []Category
	Models        []Model
}

type Category struct {
	ID   int
	Name string
}

type Manufacturer struct {
	ID           int
	Name         string
	Country      string
	FoundingYear int
}

type Model struct {
	ID             int
	Name           string
	ManufacturerID int
	Manufacturer   string
	CategoryID     int
	Category       string
	Year           int
	Specifications Specifications
	Image          string
	Country        string
	FoundingYear   int
}
type Specifications struct {
	Engine       string
	Horsepower   int
	Transmission string
	Drivetrain   string
}

type ErrMsg struct {
	Code    int
	Message string
}
