'use client';

import { useState } from 'react';
import { useMutation } from '@tanstack/react-query';
import { motion } from 'framer-motion';
import { createRequest, setAuthToken, RequestData } from '@/lib/api';

export default function PostRequest() {
  const [form, setForm] = useState<RequestData>({
    title: '',
    description: '',
    skills: [],
    budget: 0,
    city: '',
  });
  const [token, setToken] = useState<string>(
    typeof window !== 'undefined' ? localStorage.getItem('token') || '' : ''
  );

  const mutation = useMutation({
    mutationFn: createRequest,
    onSuccess: () => {
      alert('Request posted successfully!');
      setForm({ title: '', description: '', skills: [], budget: 0, city: '' });
    },
    onError: (error: any) => {
      alert(`Error: ${error.response?.data?.error || 'Something went wrong'}`);
    },
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setForm({
      ...form,
      [name]:
        name === 'budget'
          ? parseFloat(value)
          : name === 'skills'
          ? value.split(',').map((skill) => skill.trim())
          : value,
    });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setAuthToken(token);
    mutation.mutate(form);
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center py-8">
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
        className="bg-white p-8 rounded-lg shadow-md w-full max-w-md"
      >
        <h1 className="text-2xl font-bold mb-6 text-center">Post a Service Request</h1>
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
            <label className="block text-gray-700">Skills (comma-separated)</label>
            <input
              type="text"
              name="skills"
              value={form.skills.join(',')}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Budget</label>
            <input
              type="number"
              name="budget"
              value={form.budget}
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
            {mutation.isPending ? 'Posting...' : 'Post Request'}
          </button>
        </form>
      </motion.div>
    </div>
  );
}