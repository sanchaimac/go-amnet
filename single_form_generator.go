package fundconnext

import (
	"github.com/phpdave11/gofpdf"
)

type FormMeta struct {
	IntermediaryName string
	ApplicantDate    string
}

func renderHeader(pdf *gofpdf.Fpdf, formMeta FormMeta) {
	gap := 12.3
	pdf.SetFont("THSarabunNew", "bold", 15)
	pdf.Text(228, 75, formMeta.IntermediaryName)
	pdf.SetFont("THSarabunNew", "bold", 12)
	for i, v := range []rune(formMeta.ApplicantDate[6:8]) {
		pdf.Text(455+float64(i)*gap, 73, string(v))
	}
	for i, v := range []rune(formMeta.ApplicantDate[4:6]) {
		pdf.Text(482+float64(i)*gap, 73, string(v))
	}
	for i, v := range []rune(formMeta.ApplicantDate[0:4]) {
		pdf.Text(509+float64(i)*gap, 73, string(v))
	}
}

// TODO: Passport and Alien
func renderIDType(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)
	if customer.IdentificationCardType == "CITIZEN_CARD" && customer.AccompanyingDocument != nil {
		if *customer.AccompanyingDocument == "CITIZEN_CARD" {
			pdf.Text(121, 123, "\u2714")
			pdf.SetFont("THSarabunNew", "bold", 12)

			// Card Number
			gap := 12.3
			pdf.Text(350, 120, customer.CardNumber[0:1])
			for i, v := range []rune(customer.CardNumber[1:5]) {
				pdf.Text(364+float64(i)*gap, 120, string(v))
			}
			for i, v := range []rune(customer.CardNumber[5:10]) {
				pdf.Text(416+float64(i)*gap, 120, string(v))
			}
			for i, v := range []rune(customer.CardNumber[10:12]) {
				pdf.Text(479+float64(i)*gap, 120, string(v))
			}
			pdf.Text(506, 120, customer.CardNumber[12:])

			if customer.CardExpiryDate != "N/A" {
				for i, v := range []rune(customer.CardExpiryDate[6:8]) {
					pdf.Text(289+float64(i)*gap, 139, string(v))
				}
				for i, v := range []rune(customer.CardExpiryDate[4:6]) {
					pdf.Text(316+float64(i)*gap, 139, string(v))
				}
				for i, v := range []rune(customer.CardExpiryDate[0:4]) {
					pdf.Text(342+float64(i)*gap, 139, string(v))
				}
			} else {
				pdf.SetFont("ZapfDingbats", "", 10)
				pdf.Text(399, 141, "\u2714")
			}
		} else if *customer.AccompanyingDocument == "ALIEN_CARD" {
			pdf.Text(121, 161, "\u2714")
			pdf.SetFont("THSarabunNew", "bold", 12)

			gap := 12.3
			pdf.Text(350, 158, customer.CardNumber[0:1])
			for i, v := range []rune(customer.CardNumber[1:5]) {
				pdf.Text(364+float64(i)*gap, 158, string(v))
			}
			for i, v := range []rune(customer.CardNumber[5:10]) {
				pdf.Text(416+float64(i)*gap, 158, string(v))
			}
			for i, v := range []rune(customer.CardNumber[10:12]) {
				pdf.Text(479+float64(i)*gap, 158, string(v))
			}
			pdf.Text(506, 158, customer.CardNumber[12:])

			for i, v := range []rune(customer.CardExpiryDate[6:8]) {
				pdf.Text(290+float64(i)*gap, 177, string(v))
			}
			for i, v := range []rune(customer.CardExpiryDate[4:6]) {
				pdf.Text(317+float64(i)*gap, 177, string(v))
			}
			for i, v := range []rune(customer.CardExpiryDate[0:4]) {
				pdf.Text(343+float64(i)*gap, 177, string(v))
			}
		}
	} else if customer.IdentificationCardType == "PASSPORT" {
		pdf.Text(121, 199, "\u2714")
		pdf.SetFont("THSarabunNew", "bold", 12)

		gap := 12.3
		for i, v := range []rune(customer.CardNumber) {
			pdf.Text(350+float64(i)*gap, 198, string(v))
		}
		for i, v := range []rune(customer.CardExpiryDate[6:8]) {
			pdf.Text(292+float64(i)*gap, 219, string(v))
		}
		for i, v := range []rune(customer.CardExpiryDate[4:6]) {
			pdf.Text(319+float64(i)*gap, 219, string(v))
		}
		for i, v := range []rune(customer.CardExpiryDate[0:4]) {
			pdf.Text(345+float64(i)*gap, 219, string(v))
		}
		if customer.PassportCountry != nil {
			pdf.SetFont("THSarabunNew", "bold", 15)
			pdf.Text(270, 241, *customer.PassportCountry)
		}

	}

}

func renderPersonalInfo(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)
	if customer.Title == "MR" {
		pdf.Text(121, 262, "\u2714")
	} else if customer.Title == "MRS" {
		pdf.Text(193, 262, "\u2714")
	} else if customer.Title == "MISS" {
		pdf.Text(275, 262, "\u2714")
	} else if customer.Title == "OTHER" {
		pdf.Text(360, 262, "\u2714")

		if customer.TitleOther != nil {
			pdf.SetFont("THSarabunNew", "bold", 15)
			pdf.Text(421, 260, *customer.TitleOther)
		}
	}

	pdf.SetFont("THSarabunNew", "bold", 15)
	pdf.Text(170, 282, customer.THFirstName+" "+customer.THLastName)
	pdf.Text(177, 300, customer.ENFirstName+" "+customer.ENLastName)
	pdf.Text(208, 338, customer.Nationality)

	pdf.SetFont("THSarabunNew", "bold", 12)
	gap := 12.3
	for i, v := range []rune(customer.BirthDate[6:8]) {
		pdf.Text(217+float64(i)*gap, 317, string(v))
	}
	for i, v := range []rune(customer.BirthDate[4:6]) {
		pdf.Text(244+float64(i)*gap, 317, string(v))
	}
	for i, v := range []rune(customer.BirthDate[0:4]) {
		pdf.Text(271+float64(i)*gap, 317, string(v))
	}

}

func renderMartialStatus(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)

	if customer.MaritalStatus == "Single" {
		pdf.Text(167, 357, "\u2714")

	} else if customer.MaritalStatus == "Married" {
		pdf.Text(236, 357, "\u2714")
		pdf.SetFont("THSarabunNew", "bold", 15)
		if customer.Spouse != nil {
			pdf.Text(361, 396, *customer.Spouse.THFirstName+" "+*customer.Spouse.THLastName)
			pdf.Text(373, 410, *customer.Spouse.ENFirstName+" "+*customer.Spouse.ENLastName)
		}
	}
}

func renderContactInformation(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("THSarabunNew", "bold", 15)
	pdf.Text(192, 457, customer.MobileNumber)

	if customer.Email != nil {
		pdf.Text(360, 474, *customer.Email)
	}

	if customer.Fax != nil {
		pdf.Text(111, 475, *customer.Fax)
	}

	if customer.Phone != nil {
		pdf.Text(454, 457, *customer.Phone)
	}

}

func renderIdentificationDocumentAddress(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("THSarabunNew", "bold", 15)
	pdf.Text(129, 533, customer.IdentificationDocument.No)
	if customer.IdentificationDocument.Moo != nil {
		pdf.Text(249, 533, *customer.IdentificationDocument.Moo)
	}
	if customer.IdentificationDocument.Building != nil {
		pdf.Text(425, 533, *customer.IdentificationDocument.Building)
	}
	if customer.IdentificationDocument.RoomNo != nil {
		pdf.Text(136, 551, *customer.IdentificationDocument.RoomNo)
	}
	if customer.IdentificationDocument.Floor != nil {
		pdf.Text(231, 551, *customer.IdentificationDocument.Floor)
	}
	if customer.IdentificationDocument.Soi != nil {
		pdf.Text(311, 551, *customer.IdentificationDocument.Soi)
	}
	if customer.IdentificationDocument.Road != nil {
		pdf.Text(467, 551, *customer.IdentificationDocument.Road)
	}
	pdf.Text(195, 567, customer.IdentificationDocument.SubDistrict)
	pdf.Text(434, 567, customer.IdentificationDocument.District)
	pdf.Text(119, 587, customer.IdentificationDocument.Province)
	pdf.Text(336, 587, customer.IdentificationDocument.PostalCode)
	pdf.Text(459, 587, customer.IdentificationDocument.Country)
}

func renderWorkAddress(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("THSarabunNew", "bold", 15)

	if customer.Work != nil {
		if customer.CompanyName != nil {
			pdf.Text(194, 684, *customer.CompanyName)
		}
		pdf.Text(129, 700, customer.Work.No)
		if customer.Work.Moo != nil {
			pdf.Text(249, 700, *customer.Work.Moo)
		}
		if customer.Work.Building != nil {
			pdf.Text(425, 700, *customer.Work.Building)
		}
		if customer.Work.RoomNo != nil {
			pdf.Text(136, 718, *customer.Work.RoomNo)
		}
		if customer.Work.Floor != nil {
			pdf.Text(231, 718, *customer.Work.Floor)
		}
		if customer.Work.Soi != nil {
			pdf.Text(311, 718, *customer.Work.Soi)
		}
		if customer.Work.Road != nil {
			pdf.Text(467, 718, *customer.Work.Road)
		}
		pdf.Text(195, 734, customer.Work.SubDistrict)
		pdf.Text(434, 734, customer.Work.District)
		pdf.Text(119, 751, customer.Work.Province)
		pdf.Text(336, 751, customer.Work.PostalCode)
		pdf.Text(459, 751, customer.Work.Country)
		if customer.WorkPosition != nil {
			pdf.Text(140, 768, *customer.WorkPosition)
		}
	}

}

func renderCurrentAddress(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)
	if customer.CurrentAddressSameAsFlag != nil && *customer.CurrentAddressSameAsFlag == "IdDocument" {
		pdf.Text(43, 629, "\u2714")
	} else {
		pdf.Text(343, 629, "\u2714")
		pdf.SetFont("THSarabunNew", "bold", 15)
		if customer.Current != nil {
			pdf.Text(129, 655, customer.Current.No)
			if customer.Current.Moo != nil {
				pdf.Text(249, 655, *customer.Current.Moo)
			}
			if customer.Current.Building != nil {
				pdf.Text(425, 655, *customer.Current.Building)
			}
			if customer.Current.RoomNo != nil {
				pdf.Text(136, 673, *customer.Current.RoomNo)
			}
			if customer.Current.Floor != nil {
				pdf.Text(231, 673, *customer.Current.Floor)
			}
			if customer.Current.Soi != nil {
				pdf.Text(311, 673, *customer.Current.Soi)
			}
			if customer.Current.Road != nil {
				pdf.Text(467, 673, *customer.Current.Road)
			}
			pdf.Text(195, 689, customer.Current.SubDistrict)
			pdf.Text(434, 689, customer.Current.District)
			pdf.Text(119, 709, customer.Current.Province)
			pdf.Text(336, 709, customer.Current.PostalCode)
			pdf.Text(459, 709, customer.Current.Country)

		}
	}
}

func renderOccupation(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)
	occupationPositions := map[string]*[2]float64{
		"20":  {40, 111},
		"80":  {40, 129},
		"25":  {40, 147},
		"130": {40, 203},
		"60":  {40, 221},
		"160": {40, 240},
		"70":  {40, 260},
		"150": {40, 280},
		"90":  {311, 111},
		"140": {311, 129},
		"120": {311, 147},
		"40":  {311, 203},
		"50":  {311, 221},
		"110": {311, 240},
		"30":  {311, 260},
		"170": {311, 280},
	}
	pos := occupationPositions[customer.OccupationId]
	if pos != nil {
		pdf.Text(pos[0], pos[1], "\u2714")
	}
	if customer.OccupationId == "170" {
		if customer.OccupationOther != nil {
			pdf.SetFont("THSarabunNew", "bold", 15)
			pdf.Text(330, 294, *customer.OccupationOther)
		}
	}
}

func renderBusinessType(pdf *gofpdf.Fpdf, customer IndividualCustomerDocument) {
	pdf.SetFont("ZapfDingbats", "", 10)
	busniessPositions := map[string]*[2]float64{
		"20":  {40, 376},
		"170": {40, 393},
		"40":  {40, 409},
		"60":  {40, 426},
		"130": {40, 445},

		"70":  {40, 462},
		"30":  {40, 479},
		"80":  {40, 496},
		"90":  {40, 513},
		"110": {40, 530},
		"120": {40, 547},
		"140": {40, 564},
		"155": {40, 581},
		"160": {40, 598},
		"150": {40, 615},
		"180": {40, 632},
	}
	if customer.BusinessTypeId != nil {
		pos := busniessPositions[*customer.BusinessTypeId]
		if pos != nil {
			pdf.Text(pos[0], pos[1], "\u2714")
		}
		if *customer.BusinessTypeId == "180" {
			if customer.BusinessTypeOther != nil {
				pdf.SetFont("THSarabunNew", "bold", 15)
				pdf.Text(228, 631, *customer.BusinessTypeOther)
			}
		}
	}
}

func (f *FundConnext) GenerateSingleForm(pdfPath string, customer IndividualCustomerDocument, account IndividualAccountDocument, formMeta FormMeta, dest string, opt *FormOption) error {
	pdf := setupPDF()
	pdf.AddPage()
	importPDF(pdf, pdfPath, 4)
	renderHeader(pdf, formMeta)
	renderIDType(pdf, customer)
	renderPersonalInfo(pdf, customer)
	renderMartialStatus(pdf, customer)
	renderContactInformation(pdf, customer)
	renderIdentificationDocumentAddress(pdf, customer)
	renderCurrentAddress(pdf, customer)

	pdf.AddPage()
	importPDF(pdf, pdfPath, 5)
	renderOccupation(pdf, customer)
	renderBusinessType(pdf, customer)
	renderWorkAddress(pdf, customer)

	pdf.AddPage()
	importPDF(pdf, pdfPath, 6)
	pdf.AddPage()
	importPDF(pdf, pdfPath, 7)
	pdf.AddPage()
	importPDF(pdf, pdfPath, 8)
	pdf.AddPage()
	importPDF(pdf, pdfPath, 9)
	pdf.AddPage()
	importPDF(pdf, pdfPath, 10)
	err := pdf.OutputFileAndClose(dest)

	if err != nil {
		panic(err)
	}

	return nil
}
