package utils

import (
	"fmt"
	"log"
	"transactions/entities"

	"github.com/jung-kurt/gofpdf"
)


func GenerateReceiptPDF(receipt entities.Receipt, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(190, 10, receipt.CompanyName)
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 8, fmt.Sprintf("Transaction ID: %d", receipt.TransactionID))
	pdf.Ln(8)
	pdf.Cell(190, 8, fmt.Sprintf("Order Date: %s", receipt.OrderDate.Format("2006-01-02 15:04:05")))
	pdf.Ln(12)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Customer Information:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 8, fmt.Sprintf("Full Name: %s", receipt.CustomerName))
	pdf.Ln(8)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(80, 8, "Product/Service")
	pdf.Cell(30, 8, "Price/Unit")
	pdf.Cell(30, 8, "Quantity")
	pdf.Cell(30, 8, "Total")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	for _, item := range receipt.Items {
		totalPrice := float64(item.UnitPrice)
		pdf.Cell(80, 8, item.Name)
		pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", totalPrice), "", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", totalPrice), "", 0, "C", false, 0, "")
		pdf.Ln(8)
	}

	pdf.Ln(8)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(190, 8, fmt.Sprintf("Total Amount: %.2f", float64(receipt.TotalAmount)))
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 8, fmt.Sprintf("Payment Method: %s", receipt.PaymentMethod))
	pdf.Ln(8)

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return err
	}

	log.Println("Receipt generated successfully:", filename)
	return nil
}
