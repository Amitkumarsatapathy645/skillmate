import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8000/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const setAuthToken = (token: string | null) => {
  if (token) {
    api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete api.defaults.headers.common['Authorization'];
  }
};

export interface RegisterData {
  name: string;
  email: string;
  password: string;
  role: 'client' | 'freelancer';
}

export interface LoginData {
  email: string;
  password: string;
}

export interface ServiceData {
  title: string;
  description: string;
  price: number;
  tags: string[];
  city: string;
}

export interface RequestData {
  title: string;
  description: string;
  skills: string[];
  budget: number;
  city: string;
}

export interface Service {
  id: string;
  title: string;
  description: string;
  price: number;
  tags: string[];
  city: string;
  freelancer: string;
  created_at: number;
}

export const register = async (data: RegisterData) => {
  const response = await api.post('/auth/register', data);
  return response.data;
};

export const login = async (data: LoginData) => {
  const response = await api.post('/auth/login', data);
  return response.data;
};

export const browseServices = async (params: {
  skill?: string;
  city?: string;
  min_price?: string;
  max_price?: string;
}) => {
  const response = await api.get<Service[]>('/services', { params });
  return response.data;
};

export const createService = async (data: ServiceData) => {
  const response = await api.post('/services', data);
  return response.data;
};

export const getServicesByFreelancer = async () => {
  const response = await api.get<Service[]>('/services');
  return response.data;
};

export const createRequest = async (data: RequestData) => {
  const response = await api.post('/requests', data);
  return response.data;
};

export default api;