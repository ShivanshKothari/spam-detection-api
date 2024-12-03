import Link from 'next/link';
import { ShieldCheckIcon, MagnifyingGlassIcon, ExclamationTriangleIcon } from '@heroicons/react/24/outline';

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-50 to-white dark:from-gray-900 dark:to-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 sm:py-24">
        <div className="text-center">
          <h1 className="text-4xl font-bold tracking-tight text-gray-900 dark:text-white sm:text-6xl">
            Spam Detection Platform
          </h1>
          <p className="mt-6 text-lg leading-8 text-gray-600 dark:text-gray-300">
            Protect yourself from unwanted calls and messages. Search, report, and stay informed about potential spam numbers.
          </p>
          <div className="mt-10 flex items-center justify-center gap-x-6">
            <Link
              href="/login"
              className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Get started
            </Link>
            <Link href="/search" className="text-sm font-semibold leading-6 text-gray-900 dark:text-white">
              Search numbers <span aria-hidden="true">→</span>
            </Link>
          </div>
        </div>

        <div className="mt-20">
          <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
            <div className="relative p-6 bg-white dark:bg-gray-800 rounded-lg shadow-sm">
              <div className="absolute top-6 left-6">
                <ShieldCheckIcon className="h-8 w-8 text-indigo-600" />
              </div>
              <div className="mt-8 pt-4">
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white">Protect Yourself</h3>
                <p className="mt-2 text-gray-600 dark:text-gray-300">
                  Stay safe from scams and unwanted calls by checking numbers before answering.
                </p>
              </div>
            </div>

            <div className="relative p-6 bg-white dark:bg-gray-800 rounded-lg shadow-sm">
              <div className="absolute top-6 left-6">
                <MagnifyingGlassIcon className="h-8 w-8 text-indigo-600" />
              </div>
              <div className="mt-8 pt-4">
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white">Search Numbers</h3>
                <p className="mt-2 text-gray-600 dark:text-gray-300">
                  Quickly search and verify phone numbers against our spam database.
                </p>
              </div>
            </div>

            <div className="relative p-6 bg-white dark:bg-gray-800 rounded-lg shadow-sm">
              <div className="absolute top-6 left-6">
                <ExclamationTriangleIcon className="h-8 w-8 text-indigo-600" />
              </div>
              <div className="mt-8 pt-4">
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white">Report Spam</h3>
                <p className="mt-2 text-gray-600 dark:text-gray-300">
                  Help the community by reporting spam numbers you encounter.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
