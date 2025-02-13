import React, { useState } from 'react';

const FilterComponent = () => {
  const [openSections, setOpenSections] = useState({
    price: false,
    star: false,
    guestRating: false,
    amenities: false,
    propertyType: false,
  });

  const toggleSection = (section) => {
    setOpenSections((prevState) => ({
      ...prevState,
      [section]: !prevState[section],
    }));
  };

  const [priceRange, setPriceRange] = useState([0, 100000]);
  const [stars, setStars] = useState([]);
  const [guestRating, setGuestRating] = useState(null);
  const [amenities, setAmenities] = useState([]);
  const [propertyTypes, setPropertyTypes] = useState([]);

  const handlePriceChange = (e) => {
    const newValue = parseInt(e.target.value, 10);
    setPriceRange([priceRange[0], newValue]);
  };

  const toggleStar = (star) => {
    setStars((prevStars) =>
      prevStars.includes(star)
        ? prevStars.filter((s) => s !== star)
        : [...prevStars, star]
    );
  };

  const selectGuestRating = (rating) => {
    setGuestRating(rating);
  };

  const toggleAmenity = (amenity) => {
    setAmenities((prevAmenities) =>
      prevAmenities.includes(amenity)
        ? prevAmenities.filter((a) => a !== amenity)
        : [...prevAmenities, amenity]
    );
  };

  const togglePropertyType = (type) => {
    setPropertyTypes((prevTypes) =>
      prevTypes.includes(type)
        ? prevTypes.filter((t) => t !== type)
        : [...prevTypes, type]
    );
  };

  return (
    <div className="w-full max-w-sm bg-white rounded-lg shadow-lg p-4">
      {/* Google Map */}
      <div className="mb-6 relative flex justify-center items-center">
        <div className="h-56 rounded-lg overflow-hidden">
          <iframe
            src="https://www.google.com/maps/embed?pb=..."
            className="w-full h-full"
            allowFullScreen=""
            loading="lazy"
            referrerPolicy="no-referrer-when-downgrade"
          />
        </div>
        <button className="absolute w-32 h-10 py-2 px-4 bg-purple-600 text-white rounded-lg text-sm font-medium">
          Show on map
        </button>
      </div>

      {/* Accordion Sections */}
      <div className="mb-4">
        {/* Price Filter */}
        <button
          onClick={() => toggleSection('price')}
          className="w-full py-2 bg-gray-200 text-left rounded-lg font-medium"
        >
          Price
        </button>
        {openSections.price && (
          <div className="mt-2 p-2 bg-gray-100 rounded-lg">
            <input
              type="range"
              className="w-full"
              min="0"
              max="100000"
              value={priceRange[1]}
              onChange={handlePriceChange}
            />
            <div className="flex justify-between text-sm">
              <span>₦{priceRange[0]}</span>
              <span>₦{priceRange[1]}</span>
            </div>
          </div>
        )}
      </div>

      <div className="mb-4">
        {/* Star Filter */}
        <button
          onClick={() => toggleSection('star')}
          className="w-full py-2 bg-gray-200 text-left rounded-lg font-medium"
        >
          Star
        </button>
        {openSections.star && (
          <div className="mt-2 p-2 bg-gray-100 rounded-lg">
            {Array.from({ length: 5 }).map((_, index) => (
              <button
                key={index}
                onClick={() => toggleStar(5 - index)}
                className={`px-3 py-2 border rounded-lg ${
                  stars.includes(5 - index)
                    ? 'bg-yellow-500 text-white'
                    : 'text-yellow-500'
                }`}
              >
                ★ {5 - index}
              </button>
            ))}
          </div>
        )}
      </div>

      <div className="mb-4">
        {/* Guest Rating Filter */}
        <button
          onClick={() => toggleSection('guestRating')}
          className="w-full py-2 bg-gray-200 text-left rounded-lg font-medium"
        >
          Guest Rating
        </button>
        {openSections.guestRating && (
          <div className="mt-2 p-2 bg-gray-100 rounded-lg">
            {['Exceptional', 'Excellent', 'Very Good', 'Good'].map(
              (rating, index) => (
                <div
                  className="flex justify-between items-center mb-2"
                  key={index}
                >
                  <label className="flex items-center">
                    <input
                      type="radio"
                      name="guest-rating"
                      className="mr-2"
                      onChange={() => selectGuestRating(index + 6)}
                      checked={guestRating === index + 6}
                    />
                    {index + 6}+ {rating}
                  </label>
                </div>
              )
            )}
          </div>
        )}
      </div>

      <div className="mb-4">
        {/* Amenities Filter */}
        <button
          onClick={() => toggleSection('amenities')}
          className="w-full py-2 bg-gray-200 text-left rounded-lg font-medium"
        >
          Amenities
        </button>
        {openSections.amenities && (
          <div className="mt-2 p-2 bg-gray-100 rounded-lg">
            {[
              'Parking',
              'Swimming Pool',
              'Restaurant',
              'Room Service',
              'Fitness Center',
            ].map((amenity, index) => (
              <div
                className="flex justify-between items-center mb-2"
                key={index}
              >
                <label className="flex items-center">
                  <input
                    type="checkbox"
                    className="mr-2"
                    onChange={() => toggleAmenity(amenity)}
                    checked={amenities.includes(amenity)}
                  />
                  {amenity}
                </label>
              </div>
            ))}
          </div>
        )}
      </div>

      <div className="mb-4">
        {/* Property Type Filter */}
        <button
          onClick={() => toggleSection('propertyType')}
          className="w-full py-2 bg-gray-200 text-left rounded-lg font-medium"
        >
          Property Type
        </button>
        {openSections.propertyType && (
          <div className="mt-2 p-2 bg-gray-100 rounded-lg">
            {['Aparthotel', 'Apartment', 'Guest House', 'Hotel', 'Lodge'].map(
              (type, index) => (
                <div
                  className="flex justify-between items-center mb-2"
                  key={index}
                >
                  <label className="flex items-center">
                    <input
                      type="checkbox"
                      className="mr-2"
                      onChange={() => togglePropertyType(type)}
                      checked={propertyTypes.includes(type)}
                    />
                    {type}
                  </label>
                </div>
              )
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default FilterComponent;
