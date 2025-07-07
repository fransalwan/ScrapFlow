import axios from 'axios';

const API = 'http://localhost:8080/api';

export const getInvoices = () => axios.get(`${API}/invoices`);
export const getInvoiceById = (id) => axios.get(`${API}/invoice/${id}`);
export const createScaleDetail = (data) => axios.post(`${API}/scale_detail`, data);
export const updateScaleDetail = (id, data) => axios.put(`${API}/scale_detail/${id}`, data);
export const deleteScaleDetail = (id) => axios.delete(`${API}/scale_detail/${id}`);
export const generateSummary = (invoiceId) => axios.post(`${API}/summary/${invoiceId}`);
