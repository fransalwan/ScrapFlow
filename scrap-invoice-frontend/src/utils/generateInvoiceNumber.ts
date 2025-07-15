export function generateInvoiceNumber(lastCount: number): string {
  const now = new Date()
  const dateStr = now.toISOString().slice(0, 10).replace(/-/g, '') // YYYYMMDD
  const sequence = String(lastCount + 1).padStart(4, '0') // 0001, 0002
  return `INV-${dateStr}-${sequence}`
}
