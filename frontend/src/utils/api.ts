import axios from 'axios';
import Cookies from 'js-cookie';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

// Add a request interceptor to add the auth token to requests
api.interceptors.request.use(
  (config) => {
    const token = Cookies.get('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const searchByName = async (name: string) => {
  try {
    const response = await api.get(`/searchbyname?name=${encodeURIComponent(name)}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const searchByNumber = async (phoneNumber: string) => {
  try {
    const response = await api.get(`/searchbynumber?phone_number=${encodeURIComponent(phoneNumber)}`);
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const markSpam = async (phoneNumber: string) => {
  try {
    const response = await api.post('/mark-spam', { phone_number: phoneNumber });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export default api;
