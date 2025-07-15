import type { Invoice } from "../types/invoice";

export function formatInvoiceDate(invoice: Invoice): Invoice {
  return {
    ...invoice,
    invoice_date: invoice.invoice_date.slice(0, 10), // ini boleh klo mau tampil format rapi
  }
}