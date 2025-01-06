import React from 'react';

const SearchResult = ({ results }) => {
  return (
    <div className="ml-3 w-full">
      {/* Search Results Header */}
      <h3 className="text-xl font-semibold">Ikeja: {results.length} properties found</h3>
      
      {/* Filters / Sort Options */}
      <div className="flex mt-4 space-x-4">
        <div className="p-4 border rounded-lg cursor-pointer">Recommended</div>
        <div className="p-4 border rounded-lg cursor-pointer">Price</div>
        <div className="p-4 border rounded-lg cursor-pointer">User Rating</div>
        <div className="p-4 border rounded-lg cursor-pointer">Distance</div>
      </div>

      {/* Search Results List */}
      <div className="mt-6 space-y-4">
        {results.map((result, index) => (
          <div key={index} className="flex border rounded-lg p-4 shadow-lg hover:shadow-xl transition-shadow">
            {/* Property Image */}
            <div className="flex-shrink-0 w-40 h-32 bg-gray-200 rounded-lg overflow-hidden">
              <img src={result.image} alt={result.title} className="w-full h-full object-cover" />
            </div>

            {/* Property Details */}
            <div className="ml-4 flex-grow">
              <h3 className="text-lg font-medium text-gray-800">{result.title}</h3>
              <div className="flex items-center mt-1 text-sm text-gray-600">
                <span className="mr-2 text-yellow-500">{'⭐'.repeat(result.stars)}</span>
                <span>{result.rating}</span>
                <span className="ml-2 text-gray-500">- Very Good</span>
              </div>
              <div className="mt-2 flex items-center text-sm text-gray-600">
                <span className="mr-4">{result.price}</span>
                <span className="mr-4">{result.distance} km</span>
                <span>{result.userRating} User Rating</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default SearchResult;
