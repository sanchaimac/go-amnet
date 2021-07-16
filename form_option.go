package fundconnext

import (
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

type FormOption struct {
	Start *int
}

func importPDF(pdf *gofpdf.Fpdf, pdfPath string, page int) {
	tpl := gofpdi.ImportPage(pdf, pdfPath, page, "/MediaBox")
	gofpdi.UseImportedTemplate(pdf, tpl, 0, 0, 595.28, 841.89)
}

func setupPDF() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "pt", "A4", "../templates/fonts/")
	pdf.AddUTF8Font("THSarabunNew", "", "THSarabunNew.ttf")
	pdf.AddUTF8Font("THSarabunNew", "bold", "THSarabunNew_Bold.ttf")
	pdf.AddUTF8Font("ZapfDingbats", "", "ZapfDingbats.ttf")
	return pdf
}
