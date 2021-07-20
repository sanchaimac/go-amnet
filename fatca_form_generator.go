package fundconnext

import (
	"github.com/phpdave11/gofpdf"
)

type FATCAForm struct {
	ProvidedToTH         string `json:"providedToTH"`
	ProvidedToEN         string `json:"providedToEN"`
	FATCADate            string `json:"fatcaDate"`
	Title                string `json:"title"`
	ApplicantName        string `json:"applicantName"`
	Nationality          string `json:"nationality"`
	CitizenIDCard        string `json:"citizenIDCard"`
	PassportNo           string `json:"passportNo"`
	Status1              bool   `json:"status1"`
	Status2              bool   `json:"status2"`
	Status3              bool   `json:"status3"`
	Status4              bool   `json:"status4"`
	Status5              bool   `json:"status5"`
	Status6              bool   `json:"status6"`
	Status7              bool   `json:"status7"`
	Status8              bool   `json:"status8"`
	SignatureOfApplicant string `json:"signatureOfApplicant"`
	ApplicantDate        string `json:"applicantDate"`
	SignatureOfOfficer   string `json:"signatureOfOfficer"`
	NameOfOfficer        string `json:"nameOfOfficer"`
	OfficerSignedDate    string `json:"officerSignedDate"`
}

func renderPage1IDCard(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("THSarabunNew", "", 15)
	r := []rune(form.CitizenIDCard)
	gap := 12.35
	for i, v := range r {
		pdf.Text(130+float64(i)*gap, 408, string(v))
	}
}

func renderPage1ApplicantDate(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("THSarabunNew", "", 15)
	year := form.FATCADate[0:4]
	month := form.FATCADate[4:6]
	day := form.FATCADate[6:8]

	smallGap := 13.0
	bigGap := 32.3
	for i, v := range []rune(day) {
		pdf.Text(393+float64(i)*smallGap, 318, string(v))
	}

	for i, v := range []rune(month) {
		pdf.Text(393+bigGap+float64(i)*smallGap, 318, string(v))
	}

	for i, v := range []rune(year) {
		pdf.Text(393+2.0*bigGap+float64(i)*smallGap, 318, string(v))
	}
}

func renderPage1ProvidedTo(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("THSarabunNew", "", 12)
	pdf.Text(120, 155, form.ProvidedToTH)
	pdf.Text(100, 218, form.ProvidedToEN)
}

func renderPage1ApplicantInfo(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("THSarabunNew", "", 15)
	pdf.Text(195, 375, form.Title)
	pdf.Text(425, 341, form.Nationality)
	pdf.Text(70, 362, form.ApplicantName)
}

func renderPage1CheckStatus(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("ZapfDingbats", "", 12)
	if form.Status1 {
		pdf.Text(510, 577, "\u2714")
	} else {
		pdf.Text(552, 577, "\u2714")
	}

	if form.Status2 {
		pdf.Text(508, 680, "\u2714")
	} else {
		pdf.Text(551, 680, "\u2714")
	}

	if form.Status3 {
		pdf.Text(493, 785, "\u2714")
	} else {
		pdf.Text(535, 785, "\u2714")
	}

}

func renderPage2CheckStatus(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("ZapfDingbats", "", 15)

	if form.Status4 {
		pdf.Text(509, 245, "\u2714")
	} else {
		pdf.Text(553, 243, "\u2714")
	}

	if form.Status5 {
		pdf.Text(513, 270, "\u2714")
	} else {
		pdf.Text(556, 270, "\u2714")
	}

	if form.Status6 {
		pdf.Text(517, 311, "\u2714")
	} else {
		pdf.Text(561, 311, "\u2714")
	}

	if form.Status7 {
		pdf.Text(518, 367, "\u2714")
	} else {
		pdf.Text(562, 367, "\u2714")
	}

	if form.Status8 {
		pdf.Text(519, 416, "\u2714")
	} else {
		pdf.Text(562, 416, "\u2714")
	}

}

func renderPage3Signature(pdf *gofpdf.Fpdf, form FATCAForm) {
	pdf.SetFont("THSarabunNew", "", 12)
	pdf.Text(175, 637, form.ApplicantName)
	pdf.Text(410, 630, form.ApplicantDate)
	pdf.Text(194, 700, form.NameOfOfficer)
	pdf.Text(410, 693, form.OfficerSignedDate)
}

func (f *FundConnext) GenerateFATCAForm(pdfPath string, form FATCAForm, dest string, opt *FormOption) error {
	pdf := setupPDF()

	pdf.AddPage()
	importPDF(pdf, pdfPath, 1)
	renderPage1ProvidedTo(pdf, form)
	renderPage1ApplicantInfo(pdf, form)
	renderPage1IDCard(pdf, form)
	renderPage1ApplicantDate(pdf, form)
	renderPage1CheckStatus(pdf, form)

	pdf.AddPage()
	importPDF(pdf, pdfPath, 2)
	renderPage2CheckStatus(pdf, form)

	pdf.AddPage()
	importPDF(pdf, pdfPath, 3)
	renderPage3Signature(pdf, form)

	err := pdf.OutputFileAndClose(dest)
	if err != nil {
		panic(err)
	}

	return nil
}
