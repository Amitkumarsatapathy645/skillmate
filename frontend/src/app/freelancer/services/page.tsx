'use client';

import { useState } from 'react';
import { useMutation, useQuery } from '@tanstack/react-query';
import { motion } from 'framer-motion';
import {
  createService,
  getServicesByFreelancer,
  setAuthToken,
  ServiceData,
  Service,
} from '@/lib/api';

export default function FreelancerServices() {
  const [form, setForm] = useState<ServiceData>({
    title: '',
    description: '',
    price: 0,
    tags: [],
    city: '',
  });
  const [token, setToken] = useState<string>(
    typeof window !== 'undefined' ? localStorage.getItem('token') || '' : ''
  );

  const mutation = useMutation({
    mutationFn: createService,
    onSuccess: () => {
      alert('Service posted successfully!');
      setForm({ title: '', description: '', price: 0, tags: [], city: '' });
      refetch();
    },
    onError: (error: any) => {
      alert(`Error: ${error.response?.data?.error || 'Something went wrong'}`);
    },
  });

  const { data: services, isLoading, refetch } = useQuery<Service[]>({
    queryKey: ['freelancerServices'],
    queryFn: () => {
      setAuthToken(token);
      return getServicesByFreelancer();
    },
    enabled: !!token,
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setForm({
      ...form,
      [name]:
        name === 'price'
          ? parseFloat(value)
          : name === 'tags'
          ? value.split(',').map((tag) => tag.trim())
          : value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setAuthToken(token);
    mutation.mutate(form);
  };

  return (
    <div className="min-h-screen bg-gray-100 py-8">
      <div className="max-w-6xl mx-auto px-4">
        <motion.h1
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="text-3xl font-bold text-center mb-8"
        >
          Freelancer Services
        </motion.h1>

        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="bg-white p-8 rounded-lg shadow-md mb-8"
        >
          <h2 className="text-xl font-semibold mb-4">Post a New Service</h2>
          <form onSubmit={handleSubmit}>
            <div className="mb-4">
              <label className="block text-gray-700">Title</label>
              <input
                type="text"
                name="title"
                value={form.title}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700">Description</label>
              <textarea
                name="description"
                value={form.description}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700">Price</label>
              <input
                type="number"
                name="price"
                value={form.price}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700">Tags (comma-separated)</label>
              <input
                type="text"
                name="tags"
                value={form.tags.join(',')}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700">City</label>
              <input
                type="text"
                name="city"
                value={form.city}
                onChange={handleChange}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div className="mb-4">
              <label className="block text-gray-700">JWT Token</label>
              <input
                type="text"
                value={token}
                onChange={(e) => {
                  setToken(e.target.value);
                  localStorage.setItem('token', e.target.value);
                }}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <button
              type="submit"
              className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
              disabled={mutation.isPending}
            >
              {mutation.isPending ? 'Posting...' : 'Post Service'}
            </button>
          </form>
        </motion.div>

        <h2 className="text-xl font-semibold mb-4">Your Services</h2>
        {isLoading ? (
          <p>Loading...</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {services?.map((service) => (
              <motion.div
                key={service.id}
                initial={{ opacity: 0, scale: 0.9 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ duration: 0.5 }}
                className="bg-white p-6 rounded-lg shadow-md"
              >
                <h3 className="text-lg font-semibold">{service.title}</h3>
                <p className="text-gray-600 mt-2">{service.description}</p>
                <p className="text-gray-800 font-bold mt-2">₹{service.price}</p>
                <p className="text-gray-500 mt-1">City: {service.city}</p>
                <p className="text-gray-500 mt-1">Tags: {service.tags.join(', ')}</p>
              </motion.div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}