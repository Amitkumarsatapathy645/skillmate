'use client';

import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { motion } from 'framer-motion';
import { browseServices, Service } from '@/lib/api';

export default function BrowseServices() {
  const [filters, setFilters] = useState<{
    skill: string;
    city: string;
    min_price: string;
    max_price: string;
  }>({
    skill: '',
    city: '',
    min_price: '',
    max_price: '',
  });

  const { data: services, isLoading, refetch } = useQuery<Service[]>({
    queryKey: ['services', filters],
    queryFn: () => browseServices(filters),
  });

  const handleFilterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFilters({ ...filters, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    refetch();
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
          Browse Services
        </motion.h1>

        <form onSubmit={handleSubmit} className="mb-8 flex flex-col md:flex-row gap-4">
          <input
            type="text"
            name="skill"
            value={filters.skill}
            onChange={handleFilterChange}
            placeholder="Skill (e.g., plumber)"
            className="px-4 py-2 border rounded flex-1"
          />
          <input
            type="text"
            name="city"
            value={filters.city}
            onChange={handleFilterChange}
            placeholder="City (e.g., Burla)"
            className="px-4 py-2 border rounded flex-1"
          />
          <input
            type="number"
            name="min_price"
            value={filters.min_price}
            onChange={handleFilterChange}
            placeholder="Min Price"
            className="px-4 py-2 border rounded flex-1"
          />
          <input
            type="number"
            name="max_price"
            value={filters.max_price}
            onChange={handleFilterChange}
            placeholder="Max Price"
            className="px-4 py-2 border rounded flex-1"
          />
          <button
            type="submit"
            className="bg-blue-500 text-white px-6 py-2 rounded hover:bg-blue-600"
          >
            Search
          </button>
        </form>

        {isLoading ? (
          <p className="text-center">Loading...</p>
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
                <h2 className="text-xl font-semibold">{service.title}</h2>
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