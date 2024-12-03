'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'react-hot-toast';
import { searchByName, searchByNumber, markSpam } from '@/utils/api';
import { MagnifyingGlassIcon, ExclamationTriangleIcon } from '@heroicons/react/24/outline';

interface SearchForm {
  searchTerm: string;
  searchType: 'name' | 'number';
}

interface SearchResult {
  name: string;
  phone_number: string;
  spam_likelihood: number;
}

export default function SearchPage() {
  const [isLoading, setIsLoading] = useState(false);
  const [results, setResults] = useState<SearchResult[]>([]);
  const { register, handleSubmit, formState: { errors } } = useForm<SearchForm>();

  const onSubmit = async (data: SearchForm) => {
    try {
      setIsLoading(true);
      const searchFunction = data.searchType === 'name' ? searchByName : searchByNumber;
      const response = await searchFunction(data.searchTerm);
      setResults(Array.isArray(response) ? response : [response]);
      if (response.length === 0) {
        toast.success('No results found');
      }
    } catch (error: any) {
      toast.error(error.response?.data?.error || 'Failed to search');
    } finally {
      setIsLoading(false);
    }
  };

  const handleMarkSpam = async (phoneNumber: string) => {
    try {
      await markSpam(phoneNumber);
      toast.success('Number marked as spam');
      // Refresh search results
      const updatedResults = results.map(result => {
        if (result.phone_number === phoneNumber) {
          return { ...result, spam_likelihood: result.spam_likelihood + 1 };
        }
        return result;
      });
      setResults(updatedResults);
    } catch (error: any) {
      toast.error(error.response?.data?.error || 'Failed to mark as spam');
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <div className="text-center">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">
            Search Phone Numbers
          </h1>
          <p className="mt-2 text-gray-600 dark:text-gray-400">
            Search by name or phone number to check spam status
          </p>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} className="mt-8 space-y-6">
          <div className="flex gap-4">
            <div className="flex-1">
              <input
                {...register('searchTerm', { 
                  required: 'Search term is required',
                  minLength: {
                    value: 2,
                    message: 'Search term must be at least 2 characters'
                  }
                })}
                type="text"
                className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                placeholder="Enter name or phone number"
              />
              {errors.searchTerm && (
                <p className="mt-1 text-sm text-red-600">{errors.searchTerm.message}</p>
              )}
            </div>
            <div className="w-40">
              <select
                {...register('searchType')}
                className="block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
              >
                <option value="name">By Name</option>
                <option value="number">By Number</option>
              </select>
            </div>
            <button
              type="submit"
              disabled={isLoading}
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50"
            >
              {isLoading ? (
                'Searching...'
              ) : (
                <>
                  <MagnifyingGlassIcon className="h-5 w-5 mr-2" />
                  Search
                </>
              )}
            </button>
          </div>
        </form>

        <div className="mt-8">
          {results.map((result, index) => (
            <div
              key={index}
              className="bg-white dark:bg-gray-800 shadow rounded-lg p-6 mb-4"
            >
              <div className="flex justify-between items-start">
                <div>
                  <h3 className="text-lg font-medium text-gray-900 dark:text-white">
                    {result.name}
                  </h3>
                  <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                    {result.phone_number}
                  </p>
                </div>
                <div className="text-right">
                  <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                    result.spam_likelihood > 5
                      ? 'bg-red-100 text-red-800'
                      : result.spam_likelihood > 2
                      ? 'bg-yellow-100 text-yellow-800'
                      : 'bg-green-100 text-green-800'
                  }`}>
                    Spam Reports: {result.spam_likelihood}
                  </span>
                  <button
                    onClick={() => handleMarkSpam(result.phone_number)}
                    className="mt-2 inline-flex items-center px-3 py-1 border border-transparent text-xs font-medium rounded shadow-sm text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                  >
                    <ExclamationTriangleIcon className="h-4 w-4 mr-1" />
                    Mark as Spam
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
